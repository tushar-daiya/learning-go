package main

import (
	"time"
)

type Block struct {
	TimeStamp     int64
	Data          []byte //valuable information contained in the block
	PrevBlockHash []byte //hash of the previous block
	Hash          []byte //hash of the current block
	Nonce         int
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0} //create a block with data, block and prevBlockHash
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{}) //creating a genesis block for a blockchain
}
