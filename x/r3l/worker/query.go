package worker

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/relevant-community/r3l/x/r3l/types"
	"github.com/spf13/cobra"
)

type queryResult struct {
	votes       []*types.MsgVote
	scores      []*types.MsgScore
	rankSources []*types.MsgRankSource
}

func queryData(cmd *cobra.Command, clientCtx client.Context) (res queryResult, err error) {
	pageReq, err := client.ReadPageRequest(cmd.Flags())
	res = queryResult{}

	if err != nil {
		return res, err
	}

	queryClient := types.NewQueryClient(clientCtx)

	params := &types.QueryAllVoteRequest{
		Pagination: pageReq,
	}

	votes, err := queryClient.AllVote(context.Background(), params)
	if err != nil {
		return res, err
	}

	paramsScore := &types.QueryAllScoreRequest{
		Pagination: pageReq,
	}
	scores, err := queryClient.AllScore(context.Background(), paramsScore)
	if err != nil {
		return res, err
	}

	paramsRankSource := &types.QueryAllRankSourceRequest{
		Pagination: pageReq,
	}
	rankSources, err := queryClient.AllRankSource(context.Background(), paramsRankSource)
	if err != nil {
		return res, err
	}
	res.votes = votes.Vote
	res.scores = scores.Score
	res.rankSources = rankSources.RankSource

	return res, nil
}
