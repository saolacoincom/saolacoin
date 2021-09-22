// Copyright (c) 2020 The Meter.io developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package preset

import (
	"fmt"
)

// The initial version of main network is Edison.
type PresetConfig struct {
	CommitteeMinSize int
	CommitteeMaxSize int
	DelegateMaxSize  int
	DiscoServer      string
	DiscoTopic       string
}

var (
	MainPresetConfig = &PresetConfig{
		CommitteeMinSize: 1,
		CommitteeMaxSize: 300,
		DelegateMaxSize:  300,
		DiscoServer:      "enode://a300ec25b320faf2517f3c323207ad2d2ec7f1ef03dc5361df7240e2ee7603e8d91270c0f248789264e15f720d0cc215f26b93f06aece313ccfba05a34e56735@47.243.88.46:55555",
		DiscoTopic:       "mainnet",
	}

	ShoalPresetConfig = &PresetConfig{
		CommitteeMinSize: 11,
		CommitteeMaxSize: 300,
		DelegateMaxSize:  500,
		DiscoServer:      "enode://edef07cf4108b1b8be09b34290cda09e490a23b1380cbb20cefc250ebaf2470dd470449eb080c18f48c7f1618d7b17f815729dd89ac51193a2f305cf2dee6182@13.228.91.172:55555",
		DiscoTopic:       "shoal",
	}
)

func (p *PresetConfig) ToString() string {
	return fmt.Sprintf("CommitteeMaxSize: %v DelegateMaxSize: %v DiscoServer: %v : DiscoTopic%v",
		p.CommitteeMinSize, p.CommitteeMaxSize, p.DelegateMaxSize, p.DiscoServer, p.DiscoTopic)
}
