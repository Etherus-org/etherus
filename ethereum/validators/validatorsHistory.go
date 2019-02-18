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

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/syndtr/goleveldb/leveldb/errors"
	abciTypes "github.com/tendermint/tendermint/abci/types"
)

type ValidatorsHistory struct {
	db            *ethdb.Database
	tblValidators ethdb.Database
	tblReceiver   ethdb.Database
}

type ValidatorsHistoryItem struct {
	//Validators hash returned for this block's EndBlock
	ValidatorsHash []byte
	//Validators returned for this block's EndBlock
	Validators []abciTypes.Validator
}

type ReceiverHistoryItem struct {
	//Address of the reward receiver
	Address common.Address
}

// NewValidators creates a struct exposing convenient high-level operations for interacting with
// the Validators contract
func NewValidatorsHistory(db *ethdb.Database) *ValidatorsHistory {
	return &ValidatorsHistory{
		db:            db,
		tblValidators: ethdb.NewTable(*db, "vld"),
		tblReceiver:   ethdb.NewTable(*db, "rcv"),
	}
}

func (vh *ValidatorsHistory) GetHistoryValidators(appHash []byte) (*ValidatorsHistoryItem, error) {
	val, err := vh.tblValidators.Get(appHash)
	if err == errors.ErrNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, nil
	}
	var vhi ValidatorsHistoryItem
	err = rlp.DecodeBytes(val, &vhi)
	if err != nil {
		return nil, err
	}

	return &vhi, nil
}

func (vh *ValidatorsHistory) GetHistoryReceiver(blockHash []byte) (*ReceiverHistoryItem, error) {
	val, err := vh.tblReceiver.Get(blockHash)
	if err == errors.ErrNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, nil
	}
	var vri ReceiverHistoryItem
	err = rlp.DecodeBytes(val, &vri)
	if err != nil {
		return nil, err
	}

	return &vri, nil
}

//PutHistoryValidators - Fixes validators by etherus block hash
func (vh *ValidatorsHistory) PutHistoryValidators(appHash []byte, item *ValidatorsHistoryItem) error {
	buf := new(bytes.Buffer)
	rlp.Encode(buf, item)
	return vh.tblValidators.Put(appHash, buf.Bytes())
}

//PutHistoryReceiver - Fixes receiver address by Tendermint block hash
func (vh *ValidatorsHistory) PutHistoryReceiver(blockHash []byte, item *ReceiverHistoryItem) error {
	buf := new(bytes.Buffer)
	rlp.Encode(buf, item)
	return vh.tblReceiver.Put(blockHash, buf.Bytes())
}
