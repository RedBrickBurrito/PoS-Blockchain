package blockchain

import (
	"math/rand"
	"time"
)

var Mutex = &sync.Mutex{}

func PickWinner() {

	time.Sleep(30 * time.Second)
	Mutex.Lock()
	temp := TempBlocks
	Mutex.Unlock()

	lotteryPool := []string{}
	if len(temp) > 0 {

	OUTER:

		for _, block := range temp {
			for _, node := range lotteryPool {
				if block.Validator == node {
					continue OUTER
				}
			}
		}

		Mutex.Lock()
		setValidators := Validators
		Mutex.Unlock()

		k, ok := setValidators[block.Validator]
		if ok {
			for i := 0; i < k; i++ {
				lotteryPool = append(lotteryPool, block.Validator)
			}
		}
	}

		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)
		lotteryWinner := lotteryPool[r.Intn(len(lotteryPool))]

		for _, block := range temp {
			if block.Validator == lotteryWinner {
				Mutex.Lock()
				Blockchain = append(Blockchain, block)
				Mutex.Unlock()
				for _ = range Validators {
					Announcements <- "\nwinning validator: " + lotteryWinner + "\n"
				}
				break
			}
		}
	}

	Mutex.Lock()
	TempBlocks = []Block{}
	Mutex.Unlock()

}
