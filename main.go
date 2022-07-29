package main

import (
	"math/big"

	"github.com/liberontissauri/blockchains-in-go/blockchain"
)

func main()  {
	mytarget := big.NewInt(1)
	mytarget.Lsh(mytarget, uint(255 - 4))
	myblockchain := blockchain.CreateBlockchain()
	myblockchain.Display(0)
	myblockchain.AddBlock(blockchain.GenerateNewValidBlock(myblockchain, [] byte{0xab,0xcd,0x00}, *mytarget))
	myblockchain.Display(1)
	myblockchain.AddBlock(blockchain.GenerateNewValidBlock(myblockchain, [] byte{0xef,0xaa,0x00}, *mytarget))
	myblockchain.Display(2)
}