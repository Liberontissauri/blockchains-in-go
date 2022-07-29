package main

import (
	"math/big"

	"github.com/liberontissauri/blockchains-in-go/blockchain"
)

func main()  {
	mytarget := big.NewInt(1)
	mytarget.Lsh(mytarget, uint(253))
	myblockchain := blockchain.CreateBlockchain()
	myblockchain.AddBlock([] byte{0xab,0xcd,0x00}, 0, *mytarget, []byte {0xab})
	myblockchain.AddBlock([] byte{0xef,0xaa,0x00}, 0, *mytarget, []byte {0xab})
	myblockchain.Display(0)
	myblockchain.Display(1)
	myblockchain.Display(2)
}