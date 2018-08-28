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

//go:generate abigen --sol ./contract/validators.sol --pkg contract --out ./contract/validators.go

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ya-enot/etherus/ethereum"
	"github.com/ya-enot/etherus/ethereum/validators/contract"
)

var (
	contractAddress = common.HexToAddress("0x0000000000000000000000000000000000000fff")
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
func CreateValidators(backend *ethereum.Backend, client *rpc.Client) (*ValidatorsManager, error) {
	// Retrieve the Ethereum service dependency to access the blockchain

	var addr common.Address = contractAddress
	/*	chainConfig := backend.Ethereum().ApiBackend.ChainConfig()

		switch chainConfig.ChainId.Int64() {
		case 15:
			addr = MainNetAddress
		default:
			return nil, errors.New("Unknown Network Id: " + strconv.FormatInt(chainConfig.ChainId.Int64(), 10))
		}
	*/
	//	ethereum := backend.Ethereum()
	cb := ethclient.NewClient(client)

	contract, err := NewValidators(&bind.TransactOpts{From: backend.Config().Etherbase}, addr, cb)
	if err != nil {
		return nil, err
	}

	return contract, err
}
