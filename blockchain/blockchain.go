package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

type BlockChain struct {
	tip []byte
	db  *bolt.DB
}

type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

const dbFile = "blockchaintushar.db"
const blocksBucket = "blocks"

func (bc *BlockChain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db} //create a new iterator with the tip and the database

	return bci
}

func (i *BlockchainIterator) Next() *Block { //returns the next block in the blockchain
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket)) //get the bucket
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock) //deserialize the block

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	i.currentHash = block.PrevBlockHash //update the current hash to the previous block hash

	return block
}

func (bc *BlockChain) AddBlock(data string) {
	var lastHash []byte

	err := bc.db.View(func(tx *bolt.Tx) error { //read-only transaction
		b := tx.Bucket([]byte(blocksBucket)) //get the bucket
		lastHash = b.Get([]byte("l"))        //get the last block hash

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock(data, lastHash) //create a new block with the data and the last block hash

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize()) //put the new block in the bucket
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("l"), newBlock.Hash) //update the last block hash
		if err != nil {
			log.Panic(err)
		}

		bc.tip = newBlock.Hash //update the tip to the hash of the new block

		return nil
	})
}

func NewBlockChain() *BlockChain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil) //open the database file in read-write mode
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error { //update takes a function as an argument and returns an error
		b := tx.Bucket([]byte(blocksBucket)) //getting bucket from the database

		if b == nil { //if the bucket does not exist
			fmt.Println("No blockchain exists. Creating a new one")
			genesis := NewGenesisBlock() //creating a genesis block

			b, err := tx.CreateBucket([]byte(blocksBucket)) //creating a bucket
			if err != nil {
				log.Panic(err)
			}

			err = b.Put(genesis.Hash, genesis.Serialize()) //putting genesis block in the bucket
			if err != nil {
				log.Panic(err)
			}

			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				log.Panic(err)
			}

			tip = genesis.Hash //setting the tip to the hash of the genesis block
		} else {
			tip = b.Get([]byte("l")) //if the bucket exists, get the last block hash
		}
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	blockchain := BlockChain{tip, db} //create a new blockchain with the tip and the database

	return &blockchain

}
