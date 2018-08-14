// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package validators

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ya-enot/etherus/ethereum/validators/contract"
)

var (
	key, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	addr   = crypto.PubkeyToAddress(key.PublicKey)
)

func TestValidators(t *testing.T) {
	contractBackend := backends.NewSimulatedBackend(core.GenesisAlloc{addr: {Balance: big.NewInt(9000000000000000000)}})
	transactOpts := bind.NewKeyedTransactor(key)
	// Workaround for bug estimating gas in the call to Register
	transactOpts.GasLimit = big.NewInt(1000000)

	vlds, err := DeployValidators(transactOpts, contractBackend)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	contractBackend.Commit()

	transactOpts.Value = big.NewInt(100000000)
	_, err = (&contract.ValidatorsRaw{Contract: vlds.Contract}).Transfer(transactOpts)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	contractBackend.Commit()

	has, err := vlds.HasDeposit(addr)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if has {
		t.Fatal("Deposit should not be here")
	}

	contractBackend.Commit()

	transactOpts.Value = big.NewInt(2000000000000000000)
	_, err = (&contract.ValidatorsRaw{Contract: vlds.Contract}).Transfer(transactOpts)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	contractBackend.Commit()

	has, err = vlds.HasDeposit(addr)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !has {
		t.Fatal("Deposit should be here")
	}

	contractBackend.Commit()

	t.Log(has)

}
