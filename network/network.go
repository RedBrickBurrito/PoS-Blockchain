package network

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	Block "github.com/RedBrickBurrito/pos-blockchain/blockchain"
	Blockchain "github.com/RedBrickBurrito/pos-blockchain/blockchain"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	validators := Blockchain.Validators
	mutex := Blockchain.Mutex

	go func() {
		announcements := Blockchain.Announcements

		for {
			msg := <-announcements
			io.WriteString(conn, msg)
		}
	}()

	var address string

	io.WriteString(conn, "Enter token balance:")
	scanBalance := bufio.NewScanner(conn)
	for scanBalance.Scan() {
		balance, err := strconv.Atoi(scanBalance.Text())
		if err != nil {
			log.Printf("%v not a number: %v", scanBalance.Text(), err)
			return
		}
		t := time.Now()
		address = Blockchain.CreateHash(t.String())
		validators[address] = balance

		fmt.Println(validators)
		break
	}

	io.WriteString(conn, "\n Enter a new BPM:")

	scanBPM := bufio.NewScanner(conn)

	go func() {
		for {
			for scanBPM.Scan() {
				bpm, err := strconv.Atoi(scanBPM.Text())

				if err != nil {
					log.Printf("%v not a number: %v", scanBPM.Text(), err)
					delete(validators, address)
				}

				mutex.Lock()
				Blockchain := Blockchain.Blockchain
				oldLastIndex := Blockchain[len(Blockchain)-1]
				mutex.Unlock()

				newBlock, err := Block.CreateBlock(oldLastIndex, bpm, address)
				if err != nil {
					log.Println(err)
					continue
				}
				if Block.IsBlockValid(newBlock, oldLastIndex) {
					Block.CandidateBlocks <- newBlock

				}
				io.WriteString(conn, "\n Enter a new BPM:")
			}
		}
	}()

	for {
		time.Sleep(time.Minute)
		mutex.Lock()
		output, err := json.Marshal(Blockchain.Blockchain)
		mutex.Unlock()
		if err != nil {
			log.Fatal(err)
		}
		io.WriteString(conn, string(output)+"\n")

	}
}
