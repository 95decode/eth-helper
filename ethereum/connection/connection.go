package connection

import (
	"fmt"
	"math/big"

	"github.com/95decode/eth-helper/ethereum/chain"
)

type Connection interface {
	Connect() error
	LatestBlock() (*big.Int, error)
	Close()
}

type Chain struct {
	cfg  *chain.Config
	conn Connection
	stop chan<- int
}

func InitializeChain(cfg *chain.Config, syserr chan<- error) (*Chain, error) {
	config, err := chain.ParseChainConfig(cfg)
	if err != nil {
		return nil, err
	}

	fmt.Println("chain init start")

	fmt.Println(config)

	return nil, nil
}
