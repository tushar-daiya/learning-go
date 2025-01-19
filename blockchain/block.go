package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	TimeStamp     int64
	Data          []byte //valuable information contained in the block
	PrevBlockHash []byte //hash of the previous block
	Hash          []byte //hash of the current block
	Nonce         int
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer            //buffer to store serialized data
	encoder := gob.NewEncoder(&result) //intialize a new encoder

	err := encoder.Encode(b) //encoding the block
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes() //returning the serialized data as bytes
}

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d)) //intialize a new decoder and pass the serialized data so that it can be decoded
	err := decoder.Decode(&block)                 //decodes the data and stores it in the block
	if err != nil {
		log.Panic(err)
	}

	return &block
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
