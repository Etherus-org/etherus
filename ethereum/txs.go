package ethereum

import (
	"bytes"
	"time"

	"github.com/ethereum/go-ethereum/core"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	rpcClient "github.com/tendermint/tendermint/rpc/lib/client"
)

const (
	// txChanSize is the size of channel listening to TxPreEvent.
	// The number is referenced from the size of tx pool.
	txChanSize = 4096
)

//----------------------------------------------------------------------
// Transactions sent via the go-ethereum rpc need to be routed to tendermint

// listen for txs and forward to tendermint
func (b *Backend) txBroadcastLoop() {
	b.txCh = make(chan core.TxPreEvent, txChanSize)
	b.txSub = b.ethereum.TxPool().SubscribeTxPreEvent(b.txCh)

	waitForServer(b.client)

	for {
		select {
		case event := <-b.txCh:
			if err := b.BroadcastTx(event.Tx); err != nil {
				log.Error("Broadcast error", "err", err)
			}

		// Err() channel will be closed when unsubscribing.
		case <-b.txSub.Err():
			return
		}
	}
}

// BroadcastTx broadcasts a transaction to tendermint core
// #unstable
func (b *Backend) BroadcastTx(tx *ethTypes.Transaction) error {
	result := new(ctypes.ResultBroadcastTx)

	buf := new(bytes.Buffer)
	if err := tx.EncodeRLP(buf); err != nil {
		return err
	}
	params := map[string]interface{}{
		"tx": buf.Bytes(),
	}

	_, err := b.client.Call("broadcast_tx_sync", params, result)
	return err
}

//----------------------------------------------------------------------
// wait for Tendermint to open the socket and run http endpoint

func waitForServer(c rpcClient.HTTPClient) {
	status := new(ctypes.ResultStatus)
	for {
		_, err := c.Call("status", map[string]interface{}{}, status)
		if err == nil {
			break
		}

		log.Info("Waiting for tendermint endpoint to start", "err", err)
		time.Sleep(time.Second * 3)
	}
}
