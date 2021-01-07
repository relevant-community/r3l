package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/relevant-community/r3l/x/oracle/types"
)

const (
	MethodGet          = "GET"
	RestParamClaimHash = "evidence-hash"
)

// RegisterRoutes registers oracle-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	r.HandleFunc("custom/oracle/"+types.QueryListClaim, listClaimHandler(clientCtx)).Methods("GET")
	r.HandleFunc("custom/oracle/"+types.QueryClaim, queryClaimHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	// r.HandleFunc("/oracle/claim", createClaimHandler(clientCtx)).Methods("POST")
}
