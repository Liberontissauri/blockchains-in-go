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
	for i := 0; i < len(blockchain.blocks[block_index].Data); i++ {
		data_byte := blockchain.blocks[block_index].Data[i]
		fmt.Printf("%x", data_byte)
	}

	fmt.Printf("\n")
	fmt.Printf("Hash: ")
	for i := 0; i < len(blockchain.blocks[block_index].Hash); i++ {
		hash_byte := blockchain.blocks[block_index].Hash[i]
		fmt.Printf("%x", hash_byte)
	}
	fmt.Printf("\n")
	fmt.Printf("PrevHash: ")
	for i := 0; i < len(blockchain.blocks[block_index].PrevHash); i++ {
		hash_byte := blockchain.blocks[block_index].PrevHash[i]
		fmt.Printf("%x", hash_byte)
	}
	fmt.Printf("\n\n")

}

func (blockchain *Blockchain) AddBlock(data []byte){
	prev_block_hash := blockchain.blocks[len(blockchain.blocks) - 1].Hash
	blockchain.blocks = append(blockchain.blocks, CreateNewBlock(data, prev_block_hash))
}

func (blockchain *Blockchain) AddGenesisBlock() {
	blockchain.blocks = append(blockchain.blocks, CreateGenesisBlock())
}