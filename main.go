package main

import (
	"github.com/liberontissauri/blockchains-in-go/blockchain"
)

func main()  {
	myblockchain := blockchain.CreateBlockchain()
	myblockchain.AddBlock([] byte{0xab,0xcd,0x00})
	myblockchain.AddBlock([] byte{0xef,0xaa,0x00})
	myblockchain.Display(0)
	myblockchain.Display(1)
	myblockchain.Display(2)
}