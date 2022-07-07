package blockchain

import (
	"bytes"
	"crypto/sha256"
)

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

type Blockchain struct {
	Blocks []*Block
}

func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
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
