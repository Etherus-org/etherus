package app

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"math"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"

	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

// format of query data
type jsonRequest struct {
	Method string          `json:"method"`
	ID     json.RawMessage `json:"id,omitempty"`
	Params []interface{}   `json:"params,omitempty"`
}

// rlp decode an etherum transaction
func decodeTx(txBytes []byte) (*types.Transaction, error) {
	tx := new(types.Transaction)
	rlpStream := rlp.NewStream(bytes.NewBuffer(txBytes), 0)
	if err := tx.DecodeRLP(rlpStream); err != nil {
		return nil, err
	}
	return tx, nil
}

//-------------------------------------------------------
// convenience methods for validators

// Receiver returns the receiving address based on the selected strategy
// #unstable
func (app *EthermintApplication) Receiver() common.Address {
	if app.requestBeginBlock != nil {
		validator := app.requestBeginBlock.Header.Proposer
		if validator.Power > 0 {
			pubKey := validator.PubKey.GetData()

			var pubKeyBytes [32]byte
			copy(pubKeyBytes[:], pubKey)
			nAddr, err := app.validators.Contract.GetNodeReceiver(&bind.CallOpts{BlockNumber: big.NewInt(app.requestBeginBlock.Header.Height - 1)}, pubKeyBytes)

			if err != nil {
				app.logger.Error("Error getting node for validator: ", "err", err)
			}
			app.logger.Info("The receiver for currentBlock is: ", "address", nAddr, "validator", hex.EncodeToString(pubKey))
			return nAddr
		}
	}
	return common.Address{}
}

// SetValidators sets new validators on the strategy
// #unstable
func (app *EthermintApplication) SetValidators(validators []abciTypes.Validator) {
	if app.strategy != nil {
		app.strategy.SetValidators(validators)
	}
}

// GetUpdatedValidators returns an updated validator set from the strategy
// #unstable
func (app *EthermintApplication) GetUpdatedValidators() abciTypes.ResponseEndBlock {
	if app.validators != nil && app.requestBeginBlock != nil && app.requestBeginBlock.Header.Height > 1 {
		curHeight := app.requestBeginBlock.Header.Height

		prevBlockNumber := big.NewInt(curHeight - 1)
		var prevPrevBlockNumber *big.Int
		if curHeight >= 2 {
			prevPrevBlockNumber = big.NewInt(curHeight - 2)
		}
		/*
			blockInfo, err := app.backend.GetBlock(curHeight - 1)
			if err != nil {
				ethUtils.Fatalf("could not get previous block from tendermint: %v", err)
			}

			prevBlockHash := blockInfo.BlockMeta.BlockID.Hash

			ht := ethdb.NewTable(app.appDb, "nextvalidators")
			serializedValidators, err := ht.Get(prevBlockHash)
			if err != nil && err != errors.ErrNotFound {
				ethUtils.Fatalf("error from app database: %v", err)
			}

			_ = serializedValidators
		*/

		compvals, err := app.validators.Contract.GetActiveCompactedValidators(&bind.CallOpts{BlockNumber: prevBlockNumber})
		if err == nil && len(compvals.ValidatorsCompacted) >= 0 {
			newValidators, newValsHash := getNewValidators(compvals.ValidatorsCompacted, compvals.ValidatorsPubKeys)

			//We need to return new validators only in case they have changed.
			//Because otherwise Tendermint resets Proposer choosing procedure
			if prevPrevBlockNumber != nil {
				compvalsPrev, err := app.validators.Contract.GetActiveCompactedValidators(&bind.CallOpts{BlockNumber: prevPrevBlockNumber})
				if err == nil {
					oldValidators, oldValsHash := getNewValidators(compvalsPrev.ValidatorsCompacted, compvalsPrev.ValidatorsPubKeys)
					if bytes.Compare(newValsHash, oldValsHash) == 0 {
						app.logger.Info("Validators have not changed since last block, returning no change")
						return abciTypes.ResponseEndBlock{}
					}

					updatedValidators := getUpdatedValidators(newValidators, oldValidators)
					return abciTypes.ResponseEndBlock{ValidatorUpdates: updatedValidators}
				} else {
					app.logger.Error("Can not get compacted validators: ", "err", err)
					return abciTypes.ResponseEndBlock{}
				}

			}

			//There is no previous block, so return the full set of validators
			return abciTypes.ResponseEndBlock{ValidatorUpdates: newValidators}
		} else {
			app.logger.Error("Can not get current validators: ", "err", err)
		}
	}
	return abciTypes.ResponseEndBlock{}
}

// CollectTx invokes CollectTx on the strategy
// #unstable
func (app *EthermintApplication) CollectTx(tx *types.Transaction) {
	if app.strategy != nil {
		app.strategy.CollectTx(tx)
	}
}

//Normalization factor for cut deposit. In 1 ETR
const DEPOSIT_NORMALIZATION = 232830643 // = 1 ether / 2^32
const MAX_VALIDATORS = 128

func getNewValidators(compvals [][32]byte, pubkeys [][32]byte) ([]abciTypes.Validator, []byte) {
	type Vldtr struct {
		Address    []byte
		PubKey     []byte
		Deposit    uint64
		PauseCause byte
	}

	vldtrs := make([]Vldtr, len(compvals))
	for i, v := range compvals {
		vldtr := &vldtrs[i]
		vldtr.Address = v[12:]
		vldtr.Deposit = binary.BigEndian.Uint64(v[4:12])
		vldtr.PauseCause = v[3]
		vldtr.PubKey = pubkeys[i][:]
	}

	sort.Slice(vldtrs, func(i, j int) bool {
		//From large to small
		return vldtrs[i].Deposit > vldtrs[j].Deposit
	})

	//Now let us compute logarithmic average deposit
	var depAvg float64 = 0
	var depAvgInt uint64 = 0
	var valCount int64 = 0

	for _, v := range vldtrs {
		if v.PauseCause > 0 {
			continue
		}

		//Normalize the deposit (make it look approximately in ETR)
		depAvg += math.Log(float64(v.Deposit) / DEPOSIT_NORMALIZATION)
		valCount++

		if valCount >= MAX_VALIDATORS {
			break
		}
	}

	if valCount > 0 {
		//Average deposit
		depAvg /= float64(valCount)
		depAvgInt = uint64(math.Round(math.Exp(depAvg)))
	}

	hash := sha256.New()
	bufferForUint64 := make([]byte, 8)

	//Now let us form active validators list with their voting power
	var retVal [MAX_VALIDATORS]abciTypes.Validator
	slice := retVal[0:0]
	for _, v := range vldtrs {
		if v.PauseCause > 0 { //Skip paused validators
			continue
		}

		//Compute power to ensure
		// 1. Linear growth of power up to logarithmic average of deposit
		// 2. Relatively fast (but slower than linear) growth a little beyond that average
		// 3. Fast degradation of growth at large distance from that average (if more than average)
		// 4. Maximum power can not exceed twice of the power of logarithmic average deposit

		power := v.Deposit / DEPOSIT_NORMALIZATION
		if power > depAvgInt {
			power = depAvgInt + (power-depAvgInt)*depAvgInt/power
		}

		val := abciTypes.Ed25519Validator(v.PubKey, int64(power))
		val.Address = crypto.Address(tmhash.Sum(v.PubKey))
		hash.Write(v.PubKey)

		binary.BigEndian.PutUint64(bufferForUint64, power)
		hash.Write(bufferForUint64)

		slice = append(slice, val)
		if len(slice) >= MAX_VALIDATORS {
			break
		}
	}

	hashBytes := make([]byte, sha256.Size)
	return slice, hash.Sum(hashBytes[:0])

}

func getUpdatedValidators(newValidators []abciTypes.Validator, oldValidators []abciTypes.Validator) []abciTypes.Validator {
	type MapItem struct {
		v         *abciTypes.Validator
		preserved bool
	}

	oldMap := make(map[string]MapItem)
	for i, _ := range oldValidators {
		v := &oldValidators[i]
		oldMap[string(v.Address)] = MapItem{v: v}
	}

	var retVal [MAX_VALIDATORS]abciTypes.Validator
	slice := retVal[0:0]

	//Add changed or new validators
	for _, v := range newValidators {
		item, ok := oldMap[string(v.Address)]
		if !ok || v.Power != item.v.Power { //The item is new or updated
			slice = append(slice, v)
		}
		if ok {
			//Flag an item that it made it into new set
			oldMap[string(v.Address)] = MapItem{v: item.v, preserved: true}
		}
	}

	//Now add removed validators
	for _, v := range oldMap {
		if !v.preserved {
			var oldValidator = *v.v
			oldValidator.Power = 0
			slice = append(slice, oldValidator)
		}
	}

	return slice
}
