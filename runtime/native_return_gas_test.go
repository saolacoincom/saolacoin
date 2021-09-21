package runtime

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/saolacoincom/saolacoin/builtin"
	"github.com/saolacoincom/saolacoin/lvldb"
	"github.com/saolacoincom/saolacoin/state"
	"github.com/saolacoincom/saolacoin/meter"
	"github.com/saolacoincom/saolacoin/tx"
	"github.com/saolacoincom/saolacoin/xenv"
)

func TestNativeCallReturnGas(t *testing.T) {
	kv, _ := lvldb.NewMem()
	state, _ := state.New(meter.Bytes32{}, kv)
	state.SetCode(builtin.Measure.Address, builtin.Measure.RuntimeBytecodes())

	inner, _ := builtin.Measure.ABI.MethodByName("inner")
	innerData, _ := inner.EncodeInput()
	outer, _ := builtin.Measure.ABI.MethodByName("outer")
	outerData, _ := outer.EncodeInput()

	innerOutput := New(nil, state, &xenv.BlockContext{}).ExecuteClause(
		tx.NewClause(&builtin.Measure.Address).WithData(innerData),
		0,
		math.MaxUint64,
		&xenv.TransactionContext{})
	assert.Nil(t, innerOutput.VMErr)

	outerOutput := New(nil, state, &xenv.BlockContext{}).ExecuteClause(
		tx.NewClause(&builtin.Measure.Address).WithData(outerData),
		0,
		math.MaxUint64,
		&xenv.TransactionContext{})
	assert.Nil(t, outerOutput.VMErr)

	innerGasUsed := math.MaxUint64 - innerOutput.LeftOverGas
	outerGasUsed := math.MaxUint64 - outerOutput.LeftOverGas

	// gas = enter1 + prepare2 + enter2 + leave2 + leave1
	// here returns prepare2
	assert.Equal(t, uint64(1562), outerGasUsed-innerGasUsed*2)
}
