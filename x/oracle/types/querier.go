package types

import (
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
)

const (
	QueryListClaim = "list-claim"
	QueryClaim     = "claim"
)

type QueryAllClaimParams struct {
	Page  int `json:"page" yaml:"page"`
	Limit int `json:"limit" yaml:"limit"`
}

func NewQueryAllClaimParams(page, limit int) QueryAllClaimParams {
	return QueryAllClaimParams{Page: page, Limit: limit}
}

// NewQueryClaimRequest creates a new instance of QueryClaimRequest.
func NewQueryClaimRequest(hash tmbytes.HexBytes) *QueryClaimRequest {
	return &QueryClaimRequest{ClaimHash: hash}
}

// // NewQueryAllClaimRequest creates a new instance of QueryAllClaimRequest.
// func NewQueryAllClaimRequest(pageReq *query.PageRequest) *QueryAllClaimRequest {
// 	return &QueryAllClaimRequest{Pagination: pageReq}
// }
