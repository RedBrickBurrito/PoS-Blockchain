package cli

import (
	"flag"
	"fmt"

	"github.com/RedBrickBurrito/pos-blockchain/blockchain"
)

type CommandLine struct{}

func (cli *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" initblockchain")
}

func (cli *CommandLine) initBlockchain() {
	blockchain.InitBlockChain()

	fmt.Println("Finished!")
}

func (cli *CommandLine) Run() {

	initBlockchainCmd := flag.NewFlagSet("initblockchain", flag.ExitOnError)

	if initBlockchainCmd.Parsed() {
		cli.initBlockchain()
	}
}
