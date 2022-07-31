package blockchain

import (
	"math/big"
)

func GenerateNewValidBlock(blockchain *Blockchain, data []byte, target big.Int) *Block {
	var nonce = uint64(0)
	prev_hash := blockchain.blocks[len(blockchain.blocks) - 1].Hash

	for {
		generated_header := CreateHeader(data, prev_hash, nonce, target)
		computed_hash := generated_header.ComputeHash()
		generated_block := CreateNewBlock(data, prev_hash, nonce, target, computed_hash)
		if(generated_block.IsValid()) {
			return generated_block
		}
	}
}