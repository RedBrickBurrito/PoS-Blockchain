package blockchain

import (
	"log"
	"net"
	"os"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/RedBrickBurrito/pos-blockchain/network"
)

type BlockChain struct {
	Block *Block
}

var Blockchain []Block

var Announcements = make(chan string)

var Validators = make(map[string]int)

func InitBlockChain() *BlockChain {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	t := time.Now()
	genesisBlock := Block{}
	genesisBlock = Block{0, t.String(), 0, CreateBlockHash(genesisBlock), "", ""}
	spew.Dump(genesisBlock)
	Blockchain = append(Blockchain, genesisBlock)

	server, err := net.Listen("tcp", ":"+os.Getenv("ADDR"))
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	go func() {
		for candidate := range CandidateBlocks {
			Mutex.Lock()
			TempBlocks = append(TempBlocks, candidate)
			Mutex.Unlock()
		}
	}()

	go func() {
		for {
			PickWinner()
		}
	}()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go HandleConn(conn)
	}

}
