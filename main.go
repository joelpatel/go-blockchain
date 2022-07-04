package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

type Blockchain struct {
	blocks []*Block
}

func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}

type Block struct {
	Hash     []byte // represents hash of this block. it is derived from data and prevhash + few other things too
	Data     []byte // represents data inside this block, it can be anything from ledgers to documents to images and so on
	PrevHash []byte // represents the last block's hash. having this previous hash allows us to link the blocks together, like a back linked list
}

// method to create the hash based on prevhash and data of this new block
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info) // currently just a placeholder as its too simple
	b.Hash = hash[:]
}

// function to create the actual block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func main() {
	chain := InitBlockchain()

	chain.AddBlock("First Block after Genesis.")
	chain.AddBlock("Second Block after Genesis.")
	chain.AddBlock("Third Block after Genesis.")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n\n", block.Hash)
	}
}
