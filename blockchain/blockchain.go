package main

type BlockChain struct {
	blocks []*Block //storing only a pointer to a block for memory optimisations
}

func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1] //get the previous block by the last index
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock) //appending the newly created block to the blocks slice
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}
