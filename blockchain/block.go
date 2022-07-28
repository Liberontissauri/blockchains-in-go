package blockchain

import (
	"bytes"
	"crypto"
	_ "crypto/sha256"
	"encoding/binary"
	"time"
)

type Block struct {
	Timestamp int64
	Data []byte
	PrevHash []byte
	Hash []byte
}

func (block *Block) SetHash() {
	timestamp := make([]byte, 8) // Make an empy bytes array to store the converted timestamp
	binary.LittleEndian.PutUint64(timestamp, uint64(block.Timestamp)) // Convert the int64 block.timestamp to an 8 byte and storing it in timestamp

	// Joining the timestamp, data and previous hash to create a header
	header := bytes.Join([][]byte {timestamp, block.Data, block.PrevHash}, []byte{})
	
	header_hash := crypto.SHA256.New()
	header_hash.Write(header)
	block.Hash = header_hash.Sum(nil)
}

func CreateNewBlock(data []byte, prev_block_hash []byte) *Block {
	block := &Block{time.Now().Unix(), data, prev_block_hash, []byte{}}
	block.SetHash()
	return block
}

func CreateGenesisBlock() *Block {
	return CreateNewBlock([]byte{0b00000000,0b00000000,}, []byte{})
}