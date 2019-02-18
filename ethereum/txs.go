package ethereum

import (
	"bytes"

	"github.com/ethereum/go-ethereum/core"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
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

	b.WaitForServer()

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
	buf := new(bytes.Buffer)
	if err := tx.EncodeRLP(buf); err != nil {
		return err
	}

	_, err := b.client.BroadcastTxSync(buf.Bytes())
	return err
}
