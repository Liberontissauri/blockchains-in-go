package main

import (
	"github.com/liberontissauri/blockchains-in-go/blockchain"
)

func main()  {
	myblockchain := blockchain.CreateBlockchain()
	current_target := *myblockchain.CalculateTarget()
	BLOCK_TARGET_UPDATE_RATE := 2
	i := 0
	for {
		myblockchain.Display(i)
		if(i % BLOCK_TARGET_UPDATE_RATE == 0) {current_target = *myblockchain.CalculateTarget()}
		myblockchain.AddBlock(blockchain.GenerateNewValidBlock(myblockchain, [] byte{0xab,0xcd,0x00}, current_target))
		i++
	}
}