package blockchain

import (
	"bytes"
	"fmt"
	"math"
	"math/big"
)

type Blockchain struct {
	blocks []*Block
	mining_rate int
	update_period float64
	current_target *big.Int
}

func CreateBlockchain(mining_rate int, update_period float64) *Blockchain{
	blockchain := &Blockchain{mining_rate: mining_rate, update_period: update_period}
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
	if(blockchain.blocks[block_index].IsValid()) {
		fmt.Printf("Yes")
	} else {
		fmt.Printf("No")
	}
	fmt.Printf("\n\n")

}

func (blockchain *Blockchain) AddBlock(Block *Block){
	if(len(blockchain.GetBlocks()) % blockchain.mining_rate == 0) {
		blockchain.current_target = blockchain.CalculateTarget()
	}
	blockchain.blocks = append(blockchain.blocks, Block)
}

func (blockchain *Blockchain) AddGenesisBlock() {
	blockchain.current_target = blockchain.CalculateTarget()
	blockchain.blocks = append(blockchain.blocks, CreateGenesisBlock(500))
}

func (blockchain *Blockchain) GetBlocks() []*Block {
	return blockchain.blocks
}

func (blockchain *Blockchain) GetBlock(block_index int) *Block {
	return blockchain.blocks[block_index]
}

func (blockchain *Blockchain) GetTopBlock() *Block {
	return blockchain.GetBlocks()[len(blockchain.GetBlocks())-1]
}

func (blockchain *Blockchain) GetCurrentTarget() *big.Int {
	return blockchain.current_target
}

func (blockchain *Blockchain) CalculateTarget() *big.Int {
	current_block_sequence_number := math.Floor(float64(len(blockchain.blocks) / blockchain.mining_rate))
	if(current_block_sequence_number == 0) {
		target := big.NewInt(1)
		target.Lsh(target, uint(255))
		return target
	}
	start_sequence_block := blockchain.blocks[(int(current_block_sequence_number - 1) * blockchain.mining_rate)]
	end_sequence_block := blockchain.blocks[(int(current_block_sequence_number) * blockchain.mining_rate - 1)]

	completion_time := end_sequence_block.Header.Timestamp - start_sequence_block.Header.Timestamp

	var ratio float64 = float64(completion_time) / (blockchain.update_period * float64(blockchain.mining_rate))
	if (ratio > 4) {ratio = 4}
	if (ratio < 0.25) {ratio = 0.25}

	target := blockchain.blocks[len(blockchain.blocks) - 1].Header.Target
	target_float := big.NewFloat(0).SetInt(&target)
	final_target := big.NewFloat(0).Mul(target_float, big.NewFloat(ratio))

	final_target_int, _ := final_target.Int(nil)
	return final_target_int
}

func (blockchain *Blockchain) ValidateBlockchain() bool {
	for i := 0; i < len(blockchain.blocks); i++ {
		if(!blockchain.blocks[i].IsValid()) {
			return false
		}
		if(i != 0) {
			if(bytes.Compare(blockchain.blocks[i].Header.PrevHash, blockchain.blocks[i-1].Hash) != 0) {
				return false
			}
			if(blockchain.blocks[i].Header.Timestamp < blockchain.blocks[i-1].Header.Timestamp) {
				return false
			}
		}
		if(i != len(blockchain.blocks) - 1) {
			if(bytes.Compare(blockchain.blocks[i].Hash, blockchain.blocks[i+1].Header.PrevHash) != 0) {
				return false
			}
			if(blockchain.blocks[i].Header.Timestamp > blockchain.blocks[i+1].Header.Timestamp) {
				return false
			}
		}
	}
	return true
}