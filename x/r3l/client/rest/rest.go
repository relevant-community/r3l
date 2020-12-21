package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/relevant-community/r3l/x/r3l/types"
)

const (
	MethodGet = "GET"
)

// RegisterRoutes registers r3l-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	r.HandleFunc("custom/r3l/"+types.QueryListRankSource, listRankSourceHandler(clientCtx)).Methods("GET")

	r.HandleFunc("custom/r3l/"+types.QueryListNamespace, listNamespaceHandler(clientCtx)).Methods("GET")

	r.HandleFunc("custom/r3l/"+types.QueryListScore, listScoreHandler(clientCtx)).Methods("GET")

	r.HandleFunc("custom/r3l/"+types.QueryListVote, listVoteHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	r.HandleFunc("/r3l/rankSource", createRankSourceHandler(clientCtx)).Methods("POST")

	r.HandleFunc("/r3l/namespace", createNamespaceHandler(clientCtx)).Methods("POST")

	r.HandleFunc("/r3l/score", createScoreHandler(clientCtx)).Methods("POST")

	r.HandleFunc("/r3l/vote", createVoteHandler(clientCtx)).Methods("POST")

}
