package blockchain

import (
	"fmt"
)

type Blockchain struct {
	blocks []*Block
}

func CreateBlockchain() *Blockchain{
	blockchain := &Blockchain{}
	blockchain.AddGenesisBlock()
	return blockchain
}

func (blockchain *Blockchain) Display(block_index int) {
	fmt.Printf("Index: %d\n", block_index)
	fmt.Printf("Data: ")
	for i := 0; i < len(blockchain.blocks[block_index].Header.Data); i++ {
		data_byte := blockchain.blocks[block_index].Header.Data[i]
		fmt.Printf("%x", data_byte)
	}

	fmt.Printf("\n")
	fmt.Printf("Hash:   ")
	for i := 0; i < len(blockchain.blocks[block_index].Hash); i++ {
		hash_byte := blockchain.blocks[block_index].Hash[i]
		fmt.Printf("%x", hash_byte)
	}
	fmt.Printf("\n")
	fmt.Printf("PrevHash: ")
	for i := 0; i < len(blockchain.blocks[block_index].Header.PrevHash); i++ {
		hash_byte := blockchain.blocks[block_index].Header.PrevHash[i]
		fmt.Printf("%x", hash_byte)
	}
	fmt.Printf("\n")
	fmt.Printf("Target: %x", blockchain.blocks[block_index].Target.Bytes())
	fmt.Printf("\n")
	fmt.Printf("Is Valid: ")
	if(blockchain.blocks[block_index].isValid()) {
		fmt.Printf("Yes")
	} else {
		fmt.Printf("No")
	}
	fmt.Printf("\n\n")

}

func (blockchain *Blockchain) AddBlock(Block *Block){
	blockchain.blocks = append(blockchain.blocks, Block)
}

func (blockchain *Blockchain) AddGenesisBlock() {
	blockchain.blocks = append(blockchain.blocks, CreateGenesisBlock(500))
}