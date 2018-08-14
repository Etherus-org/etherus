// Copyright 2015 The Etherus Authors
// This file is part of the Etherus library.
//
// The Etherus library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public LicValidatorse as published by
// the Free Software Foundation, either version 3 of the LicValidatorse, or
// (at your option) any later version.
//
// The Etherus library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public LicValidatorse for more details.
//
// You should have received a copy of the GNU Lesser General Public LicValidatorse
// along with the Etherus library. If not, see <http://www.gnu.org/licValidatorses/>.

// Package validators contains the access to validators contract.
package validators

//go:generate abigen --sol ./validators.sol --pkg contract --out ./validators.go

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ya-enot/etherus/ethereum"
	"github.com/ya-enot/etherus/ethereum/validators/contract"
)

var (
	MainNetAddress = common.HexToAddress("0x314159265dD8dbb310642f98f50C066173C1259b")
	TestNetAddress = common.HexToAddress("0x112234455c3a32fd11230c42e7bccd4a84e02010")
)

type ValidatorsManager struct {
	*contract.ValidatorsSession
	contractBackend bind.ContractBackend
}

// NewValidators creates a struct exposing convenient high-level operations for interacting with
// the Validators contract
func NewValidators(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*ValidatorsManager, error) {
	validators, err := contract.NewValidators(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &ValidatorsManager{
		&contract.ValidatorsSession{
			Contract:     validators,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

// DeployValidators deploys an instance of the Validators nameservice, with a 'first-in, first-served' root registrar.
func DeployValidators(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend) (*ValidatorsManager, error) {
	// Deploy the Validators registry
	validatorsAddr, _, _, err := contract.DeployValidators(transactOpts, contractBackend)
	if err != nil {
		return nil, err
	}

	validators, err := NewValidators(transactOpts, validatorsAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return validators, nil
}

// CreateValidators creates an access to Validators contract
func CreateValidators(backend *ethereum.Backend) (*ValidatorsManager, error) {
	// Retrieve the Ethereum service dependency to access the blockchain

	var addr common.Address
	switch backend.Config().NetworkId {
	case 15:
		addr = MainNetAddress
	default:
		addr = MainNetAddress
		//		return nil, errors.New("Unknown Network Id")
	}

	ethereum := backend.Ethereum()
	cb := eth.NewContractBackend(ethereum.ApiBackend)

	contract, err := NewValidators(nil, addr, cb)
	if err != nil {
		return nil, err
	}

	return contract, err
}
