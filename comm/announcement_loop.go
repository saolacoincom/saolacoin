// Copyright (c) 2020 The Meter.io developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package comm

import (
	"github.com/dfinlab/meter/block"
	"github.com/dfinlab/meter/comm/proto"
	"github.com/dfinlab/meter/meter"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/rlp"
)

type announcement struct {
	newBlockID meter.Bytes32
	peer       *Peer
}

func (c *Communicator) announcementLoop() {
	const maxFetches = 3 // per block ID

	fetchingPeers := map[discover.NodeID]bool{}
	fetchingBlockIDs := map[meter.Bytes32]int{}

	fetchDone := make(chan *announcement)

	for {
		select {
		case <-c.ctx.Done():
			return
		case ann := <-fetchDone:
			delete(fetchingPeers, ann.peer.ID())
			if n := fetchingBlockIDs[ann.newBlockID] - 1; n > 0 {
				fetchingBlockIDs[ann.newBlockID] = n
			} else {
				delete(fetchingBlockIDs, ann.newBlockID)
			}
		case ann := <-c.announcementCh:
			if f, n := fetchingPeers[ann.peer.ID()], fetchingBlockIDs[ann.newBlockID]; !f && n < maxFetches {
				fetchingPeers[ann.peer.ID()] = true
				fetchingBlockIDs[ann.newBlockID] = n + 1

				c.goes.Go(func() {
					defer func() {
						select {
						case fetchDone <- ann:
						case <-c.ctx.Done():
						}
					}()
					c.fetchBlockByID(ann.peer, ann.newBlockID)
				})
			} else {
				ann.peer.logger.Debug("skip new block ID announcement")
			}
		}
	}
}

func (c *Communicator) fetchBlockByID(peer *Peer, newBlockID meter.Bytes32) {
	if _, err := c.chain.GetBlockHeader(newBlockID); err != nil {
		if !c.chain.IsNotFound(err) {
			peer.logger.Error("failed to get block header", "err", err)
		}
	} else {
		// already in chain
		return
	}

	result, err := proto.GetBlockByID(c.ctx, peer, newBlockID)
	if err != nil {
		peer.logger.Debug("failed to get block by id", "err", err)
		return
	}
	if len(result) == 0 {
		peer.logger.Debug("get nil block by id")
		return
	}

	var blk block.Block
	if err := rlp.DecodeBytes(result, &blk); err != nil {
		peer.logger.Debug("failed to decode block got by id", "err", err)
		return
	}

	c.newBlockFeed.Send(&NewBlockEvent{
		Block: &blk,
	})
}
