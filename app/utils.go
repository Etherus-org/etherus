package app

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
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
	if app.currentBlockValidator != nil {
		pubKey := app.currentBlockValidator.PubKey.GetData()

		var pubKeyBytes [32]byte
		copy(pubKeyBytes[:], pubKey)
		nAddr, err := app.validators.Contract.GetNodeAddr(&bind.CallOpts{BlockNumber: app.previousBlockNumber}, pubKeyBytes)

		if err != nil {
			app.logger.Error("Error getting node for validator: ", "err", err)
		}
		app.logger.Info("The receiver for currentBlock is: ", "address", nAddr, "validator", hex.EncodeToString(app.currentBlockValidator.PubKey.GetData()))
		return nAddr
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
	if app.validators != nil && (app.previousBlockNumber == nil || app.previousBlockNumber.Int64() > 0) {

		prevBlockNumber := app.previousBlockNumber
		if prevBlockNumber == nil {
			prevBlockNumber = app.backend.Ethereum().BlockChain().CurrentBlock().Number()
		}
		prevPrevBlockNumber := big.NewInt(app.previousBlockNumber.Int64() - 1)
		if prevPrevBlockNumber.Int64() < 0 {
			prevPrevBlockNumber = nil
		}

		compvals, err := app.validators.Contract.GetCompactedValidators(&bind.CallOpts{BlockNumber: prevBlockNumber})
		if err == nil && len(compvals.ValidatorsCompacted) > 0 {
			newValidators, newValsHash := getUpdatedValidators(compvals.ValidatorsCompacted, compvals.ValidatorsIndex)

			//We need to return new validators only in case they have changed.
			//Because otherwise Tendermint resets Proposer choosing procedure
			if prevPrevBlockNumber != nil {
				compvalsPrev, err := app.validators.Contract.GetCompactedValidators(&bind.CallOpts{BlockNumber: prevPrevBlockNumber})
				if err == nil {
					_, oldValsHash := getUpdatedValidators(compvalsPrev.ValidatorsCompacted, compvalsPrev.ValidatorsIndex)
					if bytes.Compare(newValsHash, oldValsHash) == 0 {
						app.logger.Info("Validators have not changed since last block, returning no change")
						return abciTypes.ResponseEndBlock{}
					} else {
						app.logger.Error("Can not get previous validators: ", "err", err)
					}
				}
			}

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

//Normalization factor for cut deposit. In 0.1 ETR
var norm = big.NewInt(23283064)

func getUpdatedValidators(compvals [][32]byte, pubkeys [][32]byte) ([]abciTypes.Validator, []byte) {
	type Vldtr struct {
		Address    []byte
		PubKey     []byte
		Deposit    big.Int
		PauseCause byte
	}

	vldtrs := make([]Vldtr, len(compvals))
	for i, v := range compvals {
		vldtr := &vldtrs[i]
		vldtr.Address = v[12:]
		vldtr.Deposit.SetBytes(v[4:12])
		vldtr.PauseCause = v[3]
		vldtr.PubKey = pubkeys[i][:]
	}

	sort.Slice(vldtrs, func(i, j int) bool {
		return vldtrs[i].Deposit.Cmp(&vldtrs[j].Deposit) < 0
	})

	hash := sha256.New()

	var retVal [128]abciTypes.Validator
	slice := retVal[0:0]
	for _, v := range vldtrs {
		val := abciTypes.Ed25519Validator(v.PubKey, v.Deposit.Div(&v.Deposit, norm).Int64())
		val.Address = crypto.Address(tmhash.Sum(v.PubKey))
		hash.Write(v.PubKey)
		hash.Write(v.Deposit.Bytes())

		slice = append(slice, val)
		if len(slice) >= 128 {
			break
		}
	}

	hashBytes := make([]byte, sha256.Size)
	return slice, hash.Sum(hashBytes[:0])

}
