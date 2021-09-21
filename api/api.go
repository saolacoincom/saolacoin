// Copyright (c) 2020 The Meter.io developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package api

import (
	"net/http"
	"strings"

	"github.com/saolacoincom/saolacoin/api/accountlock"
	"github.com/saolacoincom/saolacoin/api/accounts"
	"github.com/saolacoincom/saolacoin/api/auction"
	"github.com/saolacoincom/saolacoin/api/blocks"
	"github.com/saolacoincom/saolacoin/api/debug"
	"github.com/saolacoincom/saolacoin/api/doc"
	"github.com/saolacoincom/saolacoin/api/events"
	"github.com/saolacoincom/saolacoin/api/eventslegacy"
	"github.com/saolacoincom/saolacoin/api/node"
	"github.com/saolacoincom/saolacoin/api/peers"
	"github.com/saolacoincom/saolacoin/api/slashing"
	"github.com/saolacoincom/saolacoin/api/staking"
	"github.com/saolacoincom/saolacoin/api/subscriptions"
	"github.com/saolacoincom/saolacoin/api/transactions"
	"github.com/saolacoincom/saolacoin/api/transfers"
	"github.com/saolacoincom/saolacoin/api/transferslegacy"
	"github.com/saolacoincom/saolacoin/chain"
	"github.com/saolacoincom/saolacoin/logdb"
	"github.com/saolacoincom/saolacoin/p2psrv"
	"github.com/saolacoincom/saolacoin/state"
	"github.com/saolacoincom/saolacoin/txpool"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//New return api router
func New(chain *chain.Chain, stateCreator *state.Creator, txPool *txpool.TxPool, logDB *logdb.LogDB, nw node.Network, allowedOrigins string, backtraceLimit uint32, callGasLimit uint64, p2pServer *p2psrv.Server, pubKey string) (http.HandlerFunc, func()) {
	origins := strings.Split(strings.TrimSpace(allowedOrigins), ",")
	for i, o := range origins {
		origins[i] = strings.ToLower(strings.TrimSpace(o))
	}

	router := mux.NewRouter()

	// to serve api doc and swagger-ui
	router.PathPrefix("/doc").Handler(
		http.StripPrefix("/doc/", http.FileServer(
			&assetfs.AssetFS{
				Asset:     doc.Asset,
				AssetDir:  doc.AssetDir,
				AssetInfo: doc.AssetInfo})))

	// redirect swagger-ui
	router.Path("/").HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			http.Redirect(w, req, "doc/swagger-ui/", http.StatusTemporaryRedirect)
		})

	accounts.New(chain, stateCreator, callGasLimit).
		Mount(router, "/accounts")
	eventslegacy.New(logDB).
		Mount(router, "/events")
	transferslegacy.New(logDB).
		Mount(router, "/transfers")
	eventslegacy.New(logDB).
		Mount(router, "/logs/events")
	events.New(logDB).
		Mount(router, "/logs/event")
	transferslegacy.New(logDB).
		Mount(router, "/logs/transfers")
	transfers.New(logDB).
		Mount(router, "/logs/transfer")
	blocks.New(chain).
		Mount(router, "/blocks")
	transactions.New(chain, txPool).
		Mount(router, "/transactions")
	debug.New(chain, stateCreator).
		Mount(router, "/debug")
	node.New(nw, pubKey).
		Mount(router, "/node")
	peers.New(p2pServer).Mount(router, "/peers")
	subs := subscriptions.New(chain, origins, backtraceLimit)
	subs.Mount(router, "/subscriptions")
	staking.New(chain, stateCreator).
		Mount(router, "/staking")
	slashing.New().
		Mount(router, "/slashing")
	auction.New(chain, stateCreator).
		Mount(router, "/auction")
	accountlock.New().
		Mount(router, "/accountlock")

	return handlers.CORS(
			handlers.AllowedOrigins(origins),
			handlers.AllowedHeaders([]string{"content-type"}))(router).ServeHTTP,
		subs.Close // subscriptions handles hijacked conns, which need to be closed
}
