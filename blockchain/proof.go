package blockchain

import (
	"time"
)


func pickWinner() {
	
	mutex := Blockchain.Mutex
	time.Sleep(30 * time.Second)
	mutex.Lock()
	temp := Block.TempBlocks
	mutex.Unlock()

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

		mutex.Lock()
		setValidators := Block
		mutex.Unlock()

		k, ok := setValidators[block.Block]
		if ok {
			for i := 0; i < k; i++ {
				lotteryPool = append(lotteryPool, block.Block)
			}
		}

		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)
		lotteryWinner := lotteryPool[r.Intn(len(lotteryPool))]


		for _, block := range temp {
			if block.Validator == lotteryWinner {
				mutex.Lock()
				Blockchain = append(Blockchain, block)
				mutex.Unlock()
				for _ = range validators {
					announcements <- "\nwinning validator: " + lotteryWinner + "\n"
				}
				break
			}
		}
	}

	mutex.Lock()
	tempBlocks = []Block{}
mutex.Unlock()

	}
}
