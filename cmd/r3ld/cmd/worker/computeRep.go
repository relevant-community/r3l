package worker

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/relevant-community/r3l/x/r3l/types"

	"github.com/spf13/cobra"

	rep "github.com/relevant-community/reputation/non-deterministic"
)

// ComputeReputation is the reputation worker
func ComputeReputation(cmd *cobra.Command, clientCtx client.Context) (*types.Scores, error) {
	queryData, err := queryData(cmd, clientCtx)
	if err != nil {
		fmt.Println("Error fetching data", err)
		return nil, err
	}

	updatedScores := computeRank(queryData.votes, queryData.scores, queryData.rankSources)
	scoresMsg := types.NewScores(clientCtx.GetFromAddress(), clientCtx.Height, updatedScores)

	return scoresMsg, nil
}

// computeRank runs the pagerank algorithm
func computeRank(votes []*types.MsgVote, scores []*types.Score, rankSources []*types.MsgRankSource) []types.Score {
	graph := rep.NewGraph(0.85, 1e-8, 0)

	// set personalization vector
	for _, rankSource := range rankSources {
		pNode := rep.NewNodeInput(rankSource.Account, 0, 0)
		graph.AddPersonalizationNode(pNode)
	}

	// create links
	for _, vote := range votes {
		from := rep.NewNodeInput(vote.Creator.String(), 0, 0)
		to := rep.NewNodeInput(vote.To, 0, 0)
		graph.Link(from, to, float64(vote.Amount)/1000)
	}

	// result calback
	var updatedScores = make([]types.Score, 0)
	callback := func(id string, pRank float64, nRank float64) {
		score := types.Score{
			Id:    id,
			PRank: toFixed(pRank),
			NRank: toFixed(nRank),
		}
		updatedScores = append(updatedScores, score)
	}

	graph.Rank(callback)
	return updatedScores
}

const precision uint8 = 9

func toFixed(n float64) int64 {
	return int64(math.Floor(n * math.Pow(10, float64(precision))))
}

func toFloat(n int64) float64 {
	return float64(n) / math.Pow(10, float64(precision))
}

func genrandstr(s int) string {
	b := make([]byte, s)
	_, _ = rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}
