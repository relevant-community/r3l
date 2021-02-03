package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
)

// NewScores creates a new Scores object
func NewScores(creator sdk.AccAddress, blockHeight int64, scores []Score) *Scores {
	return &Scores{
		Scores:      scores,
		BlockHeight: blockHeight,
	}
}

// Claim types needed for oracle

// Type return type of oracle Claim
func (msg *Scores) Type() string {
	return ScoreClaim
}

// Hash returns the hash of an Claim Content.
func (msg *Scores) Hash() tmbytes.HexBytes {
	bz, err := msg.Marshal()
	if err != nil {
		panic(err)
	}
	return tmhash.Sum(bz)
}

// GetRoundID returns the block height for when the data was used.
func (msg *Scores) GetRoundID() uint64 {
	return uint64(msg.BlockHeight)
}

// GetConcensusKey returns a key the oracle will use of vote consensus
// for deterministic results it should be the same as the hash of the content
// for nondeterministic content it should be a constant
func (msg *Scores) GetConcensusKey() string {
	return msg.Hash().String()
}
