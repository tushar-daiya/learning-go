package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const TARGET_BITS = 24

var maxNonce = math.MaxInt64

type ProofOfWork struct { // structure that holds a pointer to a block and pointer to a target
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork { //creates a new proof of work
	target := big.NewInt(1)
	target.Lsh(target, uint(256-TARGET_BITS)) //256 is the number of bits in a SHA-256 hash
	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.TimeStamp),
			IntToHex(int64(TARGET_BITS)),
			IntToHex(int64(nonce)),
		},
		[]byte{}, //empty byte slice means the pieces will be joined without any delimiter
	)

	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining th block with data: %s\n", pow.block.Data)

	for nonce < math.MaxInt64 { ///running while nonce is less than maxNonce or hash is less than the target
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:]) //converts bytes to big int

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++ //increasing nonce to get a different hash
		}

	}

	fmt.Print("\n\n")

	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pow.target) == -1

}
