package blockchain

import (
	"bytes"
	"crypto"
	_ "crypto/sha256"
	"encoding/binary"
	"time"
	"math/big"
)

type Header struct {
	Timestamp int64
	Data []byte
	PrevHash []byte
	Nonce uint64
	Target big.Int
}

func createHeader(data []byte, prevHash []byte, nonce uint64, target big.Int) *Header {
	return &Header{time.Now().Unix(), data, prevHash, nonce, target}
}

func (header *Header) getByteTimestamp() []byte {
	timestamp := make([]byte, 8) // Make an empy bytes array to store the converted timestamp
	binary.LittleEndian.PutUint64(timestamp, uint64(header.Timestamp)) // Convert the int64 block.timestamp to an 8 byte and storing it in timestamp
	return timestamp
}
func (header *Header) getByteNonce() []byte {
	nonce := make([]byte, 8)
	binary.LittleEndian.PutUint64(nonce, uint64(header.Nonce))
	return nonce
}

func (header *Header) getByteArray() []byte {
	return bytes.Join([][]byte {header.getByteTimestamp(), header.Data, header.PrevHash, header.getByteNonce(), header.Target.Bytes()}, []byte{})
}

func (header *Header) computeHash() []byte {
	header_hash := crypto.SHA256.New()
	header_hash.Write(header.getByteArray())
	return header_hash.Sum(nil)
}

type Block struct {
	Header Header
	Hash []byte
}

func (block *Block) isValid() bool{
	computed_hash := block.Header.computeHash()
	big_int_hash := big.NewInt(0)
	big_int_hash.SetBytes(computed_hash)
	if(!bytes.Equal(computed_hash, block.Hash) && len(block.Header.PrevHash) != 0) {return false}
	if(big_int_hash.Cmp(&block.Header.Target) == -1) {return true}
	return false
}

func CreateNewBlock(data []byte, prev_block_hash []byte, nonce uint64, target big.Int, hash []byte) *Block {
	block := &Block{*createHeader(data, prev_block_hash, nonce, target), hash}
	return block
}

func CreateGenesisBlock(target_bits uint) *Block {
	target := big.NewInt(1)
	target.Lsh(target, target_bits)

	genesis_header := createHeader([]byte{0b00000000,0b00000000,}, []byte{}, 0, *target)
	header_hash := crypto.SHA256.New()
	header_hash.Write(genesis_header.getByteArray())
	
	return CreateNewBlock([]byte{0b00000000,0b00000000,}, []byte{}, 0, *target, header_hash.Sum(nil))
}