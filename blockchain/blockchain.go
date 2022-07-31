package blockchain

import (
	"fmt"
	"math"
	"math/big"
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
	fmt.Printf("Target: %x", blockchain.blocks[block_index].Header.Target.Bytes())
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

func (blockchain *Blockchain) GetBlocks() []*Block {
	return blockchain.blocks
}

func (blockchain *Blockchain) GetBlock(block_index int) *Block {
	return blockchain.blocks[block_index]
}

func (blockchain *Blockchain) CalculateTarget() *big.Int {
	SEQUENCE_SIZE := 2

	current_block_sequence_number := math.Floor(float64(len(blockchain.blocks) / SEQUENCE_SIZE))
	if(current_block_sequence_number == 0) {
		target := big.NewInt(1)
		target.Lsh(target, uint(255))
		return target
	}
	start_sequence_block := blockchain.blocks[(int(current_block_sequence_number - 1) * SEQUENCE_SIZE)]
	end_sequence_block := blockchain.blocks[(int(current_block_sequence_number) * SEQUENCE_SIZE - 1)]

	completion_time := end_sequence_block.Header.Timestamp - start_sequence_block.Header.Timestamp

	var ratio float64 = float64(completion_time) / (60 * 10 * float64(SEQUENCE_SIZE))
	if (ratio > 4) {ratio = 4}
	if (ratio < 0.25) {ratio = 0.25}

	target := blockchain.blocks[len(blockchain.blocks) - 1].Header.Target
	target_float := big.NewFloat(0).SetInt(&target)
	final_target := big.NewFloat(0).Mul(target_float, big.NewFloat(ratio))

	final_target_int, _ := final_target.Int(nil)
	return final_target_int
}