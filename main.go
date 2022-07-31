package main

import (
	"github.com/Liberontissauri/blockchains-in-go/blockchain"
)

func main()  {
	myblockchain := blockchain.CreateBlockchain(2, 60 * 10)
	i := 0
	for {
		myblockchain.Display(i)
		myblockchain.AddBlock(blockchain.GenerateNewValidBlock(myblockchain, [] byte{0xab,0xcd,0x00}, *myblockchain.GetCurrentTarget()))
		i++
	}
}