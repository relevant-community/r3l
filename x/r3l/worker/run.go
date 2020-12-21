package worker

import (
	"fmt"
	"math"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/relevant-community/r3l/x/r3l/types"
	"github.com/spf13/cobra"

	rep "github.com/relevant-community/reputation/non-deterministic"
)

// RunWorkerProcess is our custom worker code
func RunWorkerProcess(cmd *cobra.Command, clientCtx client.Context) {
	queryData, err := queryData(cmd, clientCtx)
	if err != nil {
		fmt.Println("Error fetching data", err)
		return
	}

	updatedScores := computeRank(queryData.votes, queryData.scores, queryData.rankSources)

	// We use BlockHeight - 1 to reflect that the scores were based on data from that block
	msg := types.NewMsgScores(clientCtx.GetFromAddress(), clientCtx.Height-1, updatedScores)

	err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)

	// TODO retry if the TX fails
	if err != nil {
		fmt.Println("TX ERROR", err)
	}
}

// computeRank runs the pagerank algorithm
func computeRank(votes []*types.MsgVote, scores []*types.MsgScore, rankSources []*types.MsgRankSource) []types.Score {
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
		// fmt.Println(id, pRank, nRank)
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

// func RunWorker(clientCtx) {
// }

// func RunEmbeddedWorker(ctx sdk.Context, k keeper.Keeper) {
// defer func() {
// 	if r := recover(); r != nil {
// 		fmt.Println("Failed to run worker", r)
// 	}
// }()

// time.Sleep(4000 * time.Millisecond)
// fmt.Println("start block", ctx.BlockHeight())

// // test blow up gas
// // for i := 0; i < 100000; i++ {
// // 	k.GetAllVote(ctx)
// // }

// votes := k.GetAllVote(ctx)
// scores := k.GetAllScore(ctx)
// rankSources := k.GetAllRankSource(ctx)

// updatedScores := ComputeRank(votes, scores, rankSources)
// SetScores(ctx, k, updatedScores)

// fmt.Println("worker gas limit", ctx.BlockGasMeter().Limit(), ctx.GasMeter().GasConsumed())
// fmt.Println("out of gas / is past limit", ctx.BlockGasMeter().IsOutOfGas(), ctx.BlockGasMeter().IsPastLimit())

// endGas := ctx.BlockGasMeter().GasConsumed()
// fmt.Println("end worker process", endGas)
// }
