package blockchain

import (
	"math/big"
)

func GenerateNewValidBlock(blockchain *Blockchain, data []byte, target big.Int) *Block {
	var nonce = uint64(0)
	prev_hash := blockchain.blocks[len(blockchain.blocks) - 1].Hash

	for {
		generated_header := createHeader(data, prev_hash, nonce)
		computed_hash := generated_header.computeHash()
		generated_block := CreateNewBlock(data, prev_hash, nonce, target, computed_hash)
		if(generated_block.isValid()) {
			return generated_block
		}
	}
}