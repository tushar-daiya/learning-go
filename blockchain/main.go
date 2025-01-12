package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	TimeStamp     int64
	Data          []byte //valuable information contained in the block
	PrevBlockHash []byte //hash of the previous block
	Hash          []byte //hash of the current block
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.TimeStamp, 10))                       //formatting the int64 timestamp to a decimal string representation and then converting the string to a byte slice
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{}) //join all the block components
	hash := sha256.Sum256(headers)                                                //hashing the headers

	b.Hash = hash[:] //converting 32bytes slice from sha256.Sum256 to a slice
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}} //create a block with data, block and prevBlockHash
	block.SetHash()                                                           //set the hash by calling the setHash method on the block struct
	return block
}

type BlockChain struct {
	blocks []*Block //storing only a pointer to a block for memory optimisations
}

func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1] //get the previous block by the last index
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock) //appending the newly created block to the blocks slice
}

func NewGenesisBlock() *Block {
	return NewBlock("First Block of my BlockChain", []byte{}) //creating a genesis block for a blockchain
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

func main() {
	sharchain := NewBlockChain()

	sharchain.AddBlock("Send 100 shar coins to Sameer")
	sharchain.AddBlock("Send 50 shar coins to Piyush")

	for _, block := range sharchain.blocks {
		fmt.Printf("Prev. hash %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
