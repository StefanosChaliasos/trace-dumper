package utils

import (
	"context"
	"math/big"
	"time"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"
)

var (
	Ether = big.NewInt(1e18)
)

func WaitPendingTx(ctx context.Context, client *ethclient.Client, hash common.Hash) {
	sleep := 1000
	for {
		if _, ispending, _ := client.TransactionByHash(ctx, hash); !ispending {
			return
		}
		log.Debug("wait tx is pending", "sleep ms", sleep)
		time.Sleep(time.Millisecond * time.Duration(sleep))
	}
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
