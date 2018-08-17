package minerRewardStrategies

import (
	"github.com/ethereum/go-ethereum/common"
)

// RewardConstant is a possible strategy
type RewardProposer struct {
	proposer common.Address
}

// Receiver returns which address should receive the mining reward
func (r *RewardProposer) Receiver() common.Address {
	return common.HexToAddress("7ef5a6135f1fd6a02593eedc869c6d41d934aef8")
}
