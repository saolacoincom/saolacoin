// Copyright (c) 2020 The Meter.io developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package staking

import (
	"github.com/dfinlab/meter/meter"
	"github.com/dfinlab/meter/state"
	setypes "github.com/dfinlab/meter/script/types"
	"github.com/dfinlab/meter/xenv"
)

//
type StakingEnv struct {
	*setypes.ScriptEnv
	staking   *Staking
}

func NewStakingEnv(staking *Staking, state *state.State, txCtx *xenv.TransactionContext, to *meter.Address) *StakingEnv {
	return &StakingEnv{
		staking:   staking,
		ScriptEnv: setypes.NewScriptEnv(state, txCtx, to),
	}
}

func (env *StakingEnv) GetStaking() *Staking               { return env.staking }
