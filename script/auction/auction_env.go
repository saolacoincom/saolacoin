// Copyright (c) 2020 The Meter.io developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package auction

import (
	"github.com/saolacoincom/saolacoin/meter"
	"github.com/saolacoincom/saolacoin/state"
	setypes "github.com/saolacoincom/saolacoin/script/types"
	"github.com/saolacoincom/saolacoin/xenv"
)

//
type AuctionEnv struct {
	*setypes.ScriptEnv
	auction *Auction
}

func NewAuctionEnv(auction *Auction, state *state.State, txCtx *xenv.TransactionContext, to *meter.Address) *AuctionEnv {
	return &AuctionEnv{
		auction:   auction,
		ScriptEnv: setypes.NewScriptEnv(state, txCtx, to),
	}
}

func (env *AuctionEnv) GetAuction() *Auction { return env.auction }
