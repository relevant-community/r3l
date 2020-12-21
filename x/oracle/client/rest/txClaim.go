package rest

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/types/rest"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createClaimRequest struct {
	BaseReq     rest.BaseReq `json:"base_req"`
	Creator     string       `json:"creator"`
	BlockNumber string       `json:"blockNumber"`
	Hash        string       `json:"hash"`
}

// func createClaimHandler(clientCtx client.Context) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var req createClaimRequest
// 		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
// 			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
// 			return
// 		}

// 		baseReq := req.BaseReq.Sanitize()
// 		if !baseReq.ValidateBasic(w) {
// 			return
// 		}

// 		creator, err := sdk.AccAddressFromBech32(req.Creator)
// 		if err != nil {
// 			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
// 			return
// 		}

// 		parsedBlockNumber64, err := strconv.ParseInt(req.BlockNumber, 10, 32)
// 		if err != nil {
// 			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
// 			return
// 		}
// 		parsedBlockNumber := int32(parsedBlockNumber64)

// 		parsedHash := req.Hash

// 		msg := types.NewMsgClaim(
// 			creator,
// 			parsedBlockNumber,
// 			parsedHash,
// 		)

// 		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
// 	}
// }
