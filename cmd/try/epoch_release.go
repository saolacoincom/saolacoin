package main

import (
	"fmt"
	"math/big"

	"github.com/saolacoincom/saolacoin/reward"
	"github.com/saolacoincom/saolacoin/script/auction"
)

func calcEpochRelease() {
	lastSummary := &auction.AuctionSummary{
		RlsdMTRG: new(big.Int).Mul(big.NewInt(54801906), big.NewInt(1e14)),
		RsvdMTRG: big.NewInt(0),
	}
	release, err := reward.ComputeEpochReleaseWithInflation(1111, 1113, lastSummary)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("RELEASE:", release)
}
