package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type CLI struct {
	bc *BlockChain
}

func (cli *CLI) addBlock(data string) { //add a block to the blockchain
	cli.bc.AddBlock(data)
	fmt.Println("Block added!")
}

func (cli *CLI) printChain() { //print all the blocks of the blockchain
	bci := cli.bc.Iterator() //get the iterator

	for {
		block := bci.Next() //get the next block

		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)                                //create a new proof of work
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate())) //validate the proof of work
		fmt.Println()

		if len(block.PrevBlockHash) == 0 { //if the previous block hash is empty, we have reached the genesis block
			break
		}
	}
}

func (cli *CLI) printUsage() { //utility function to print the usage of the CLI
	fmt.Println("Usage:")
	fmt.Println("  addblock -data BLOCK_DATA - add a block to the blockchain")
	fmt.Println("  printchain - print all the blocks of the blockchain")
}

func (cli *CLI) validateArgs() { //utility function to validate the arguments
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) Run() { //main function to run the CLI
	cli.validateArgs() //validating the arguments

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)     //creating a new flag set for addblock
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError) //creating a new flag set for printchain

	addBlockData := addBlockCmd.String("data", "", "Block data") //flag to get the data for the block

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}
