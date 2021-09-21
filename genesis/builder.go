// Copyright (c) 2020 The Meter.io developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package genesis

import (
	"math"

	"github.com/saolacoincom/saolacoin/block"
	"github.com/saolacoincom/saolacoin/lvldb"
	"github.com/saolacoincom/saolacoin/meter"
	"github.com/saolacoincom/saolacoin/runtime"
	"github.com/saolacoincom/saolacoin/state"
	"github.com/saolacoincom/saolacoin/tx"
	"github.com/saolacoincom/saolacoin/xenv"
	"github.com/pkg/errors"
)

// Builder helper to build genesis block.
type Builder struct {
	timestamp uint64
	gasLimit  uint64

	stateProcs []func(state *state.State) error
	calls      []call
	extraData  [28]byte
}

type call struct {
	clause *tx.Clause
	caller meter.Address
}

// Timestamp set timestamp.
func (b *Builder) Timestamp(t uint64) *Builder {
	b.timestamp = t
	return b
}

// GasLimit set gas limit.
func (b *Builder) GasLimit(limit uint64) *Builder {
	b.gasLimit = limit
	return b
}

// State add a state process
func (b *Builder) State(proc func(state *state.State) error) *Builder {
	b.stateProcs = append(b.stateProcs, proc)
	return b
}

// Call add a contrct call.
func (b *Builder) Call(clause *tx.Clause, caller meter.Address) *Builder {
	b.calls = append(b.calls, call{clause, caller})
	return b
}

// ExtraData set extra data, which will be put into last 28 bytes of genesis parent id.
func (b *Builder) ExtraData(data [28]byte) *Builder {
	b.extraData = data
	return b
}

// ComputeID compute genesis ID.
func (b *Builder) ComputeID() (meter.Bytes32, error) {
	kv, err := lvldb.NewMem()
	if err != nil {
		return meter.Bytes32{}, err
	}
	blk, _, err := b.Build(state.NewCreator(kv))
	if err != nil {
		return meter.Bytes32{}, err
	}
	return blk.Header().ID(), nil
}

// Build build genesis block according to presets.
func (b *Builder) Build(stateCreator *state.Creator) (blk *block.Block, events tx.Events, err error) {
	state, err := stateCreator.NewState(meter.Bytes32{})
	if err != nil {
		return nil, nil, err
	}

	for _, proc := range b.stateProcs {
		if err := proc(state); err != nil {
			return nil, nil, errors.Wrap(err, "state process")
		}
	}

	rt := runtime.New(nil, state, &xenv.BlockContext{
		Time:     b.timestamp,
		GasLimit: b.gasLimit,
	})

	for _, call := range b.calls {
		out := rt.ExecuteClause(call.clause, 0, math.MaxUint64, &xenv.TransactionContext{
			Origin: call.caller,
		})
		if out.VMErr != nil {
			return nil, nil, errors.Wrap(out.VMErr, "vm")
		}
		for _, event := range out.Events {
			events = append(events, (*tx.Event)(event))
		}
	}

	stage := state.Stage()
	stateRoot, err := stage.Commit()
	if err != nil {
		return nil, nil, errors.Wrap(err, "commit state")
	}

	parentID := meter.Bytes32{0xff, 0xff, 0xff, 0xff} //so, genesis number is 0
	copy(parentID[4:], b.extraData[:])

	return new(block.Builder).
		ParentID(parentID).
		Timestamp(b.timestamp).
		GasLimit(b.gasLimit).
		StateRoot(stateRoot).
		ReceiptsRoot(tx.Transactions(nil).RootHash()).
		Build(), events, nil
}
