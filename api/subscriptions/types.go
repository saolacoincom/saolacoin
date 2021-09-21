// Copyright (c) 2020 The Meter.io developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package subscriptions

import (
	"github.com/saolacoincom/saolacoin/block"
	"github.com/saolacoincom/saolacoin/chain"
	"github.com/saolacoincom/saolacoin/meter"
	"github.com/saolacoincom/saolacoin/tx"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
)

//BlockMessage block piped by websocket
type BlockMessage struct {
	Number       uint32          `json:"number"`
	ID           meter.Bytes32   `json:"id"`
	Size         uint32          `json:"size"`
	ParentID     meter.Bytes32   `json:"parentID"`
	Timestamp    uint64          `json:"timestamp"`
	GasLimit     uint64          `json:"gasLimit"`
	Beneficiary  meter.Address   `json:"beneficiary"`
	GasUsed      uint64          `json:"gasUsed"`
	TotalScore   uint64          `json:"totalScore"`
	TxsRoot      meter.Bytes32   `json:"txsRoot"`
	StateRoot    meter.Bytes32   `json:"stateRoot"`
	ReceiptsRoot meter.Bytes32   `json:"receiptsRoot"`
	Signer       meter.Address   `json:"signer"`
	Transactions []meter.Bytes32 `json:"transactions"`
	Obsolete     bool            `json:"obsolete"`
}

func convertBlock(b *chain.Block) (*BlockMessage, error) {
	header := b.Header()
	signer, err := header.Signer()
	if err != nil {
		return nil, err
	}

	txs := b.Transactions()
	txIds := make([]meter.Bytes32, len(txs))
	for i, tx := range txs {
		txIds[i] = tx.ID()
	}
	return &BlockMessage{
		Number:       header.Number(),
		ID:           header.ID(),
		ParentID:     header.ParentID(),
		Timestamp:    header.Timestamp(),
		TotalScore:   header.TotalScore(),
		GasLimit:     header.GasLimit(),
		GasUsed:      header.GasUsed(),
		Beneficiary:  header.Beneficiary(),
		Signer:       signer,
		Size:         uint32(b.Size()),
		StateRoot:    header.StateRoot(),
		ReceiptsRoot: header.ReceiptsRoot(),
		TxsRoot:      header.TxsRoot(),
		Transactions: txIds,
		Obsolete:     b.Obsolete,
	}, nil
}

type LogMeta struct {
	BlockID        meter.Bytes32 `json:"blockID"`
	BlockNumber    uint32        `json:"blockNumber"`
	BlockTimestamp uint64        `json:"blockTimestamp"`
	TxID           meter.Bytes32 `json:"txID"`
	TxOrigin       meter.Address `json:"txOrigin"`
}

//TransferMessage transfer piped by websocket
type TransferMessage struct {
	Sender    meter.Address         `json:"sender"`
	Recipient meter.Address         `json:"recipient"`
	Amount    *math.HexOrDecimal256 `json:"amount"`
	Token     byte                  `json:"token"`
	Meta      LogMeta               `json:"meta"`
	Obsolete  bool                  `json:"obsolete"`
}

func convertTransfer(header *block.Header, tx *tx.Transaction, transfer *tx.Transfer, obsolete bool) (*TransferMessage, error) {
	signer, err := tx.Signer()
	if err != nil {
		return nil, err
	}

	return &TransferMessage{
		Sender:    transfer.Sender,
		Recipient: transfer.Recipient,
		Amount:    (*math.HexOrDecimal256)(transfer.Amount),
		Token:     transfer.Token,
		Meta: LogMeta{
			BlockID:        header.ID(),
			BlockNumber:    header.Number(),
			BlockTimestamp: header.Timestamp(),
			TxID:           tx.ID(),
			TxOrigin:       signer,
		},
		Obsolete: obsolete,
	}, nil
}

//EventMessage event piped by websocket
type EventMessage struct {
	Address  meter.Address   `json:"address"`
	Topics   []meter.Bytes32 `json:"topics"`
	Data     string          `json:"data"`
	Meta     LogMeta         `json:"meta"`
	Obsolete bool            `json:"obsolete"`
}

func convertEvent(header *block.Header, tx *tx.Transaction, event *tx.Event, obsolete bool) (*EventMessage, error) {
	signer, err := tx.Signer()
	if err != nil {
		return nil, err
	}
	return &EventMessage{
		Address: event.Address,
		Data:    hexutil.Encode(event.Data),
		Meta: LogMeta{
			BlockID:        header.ID(),
			BlockNumber:    header.Number(),
			BlockTimestamp: header.Timestamp(),
			TxID:           tx.ID(),
			TxOrigin:       signer,
		},
		Topics:   event.Topics,
		Obsolete: obsolete,
	}, nil
}

// EventFilter contains options for contract event filtering.
type EventFilter struct {
	Address *meter.Address // restricts matches to events created by specific contracts
	Topic0  *meter.Bytes32
	Topic1  *meter.Bytes32
	Topic2  *meter.Bytes32
	Topic3  *meter.Bytes32
	Topic4  *meter.Bytes32
}

// Match returs whether event matches filter
func (ef *EventFilter) Match(event *tx.Event) bool {
	if (ef.Address != nil) && (*ef.Address != event.Address) {
		return false
	}

	matchTopic := func(topic *meter.Bytes32, index int) bool {
		if topic != nil {
			if len(event.Topics) <= index {
				return false
			}

			if *topic != event.Topics[index] {
				return false
			}
		}
		return true
	}

	return matchTopic(ef.Topic0, 0) &&
		matchTopic(ef.Topic1, 1) &&
		matchTopic(ef.Topic2, 2) &&
		matchTopic(ef.Topic3, 3) &&
		matchTopic(ef.Topic4, 4)
}

// TransferFilter contains options for contract transfer filtering.
type TransferFilter struct {
	TxOrigin  *meter.Address // who send transaction
	Sender    *meter.Address // who transferred tokens
	Recipient *meter.Address // who received tokens
}

// Match returs whether transfer matches filter
func (tf *TransferFilter) Match(transfer *tx.Transfer, origin meter.Address) bool {
	if (tf.TxOrigin != nil) && (*tf.TxOrigin != origin) {
		return false
	}

	if (tf.Sender != nil) && (*tf.Sender != transfer.Sender) {
		return false
	}

	if (tf.Recipient != nil) && (*tf.Recipient != transfer.Recipient) {
		return false
	}
	return true
}

type BeatMessage struct {
	Number    uint32        `json:"number"`
	ID        meter.Bytes32 `json:"id"`
	ParentID  meter.Bytes32 `json:"parentID"`
	Timestamp uint64        `json:"timestamp"`
	Bloom     string        `json:"bloom"`
	K         uint32        `json:"k"`
	Obsolete  bool          `json:"obsolete"`
}
