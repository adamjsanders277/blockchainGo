package main

import {
	"fmt"
}

type Block struct {
	Pos				int
	Data			BookCheckout
	Timestamp		string
	Hash 			string
	PrevHash		string
}

type BookCheckout struct {
	BookID			string 	'json:"book_id"'
	User			string 	'json:"user"'
	CheckoutDate	string 	'json:"checkout_date"'
	IsGenesis		bool	'json:"is_genesis"'
}

type Book struct {
	ID 				string	'json:"id"'
	Title			string	'json:"title"'
	Author			string	'json:"author"'
	PublishDate		string	'json:"publish_date"'
	ISBN			string	'json:"isbn"'
}

// Blockchain is an ordered list of blocks
type Blockchain struct {
	blocks []*Block
}

var BlockChain *Blockchain

func (bc *Blockchain) AddBlock (data BookCheckout) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	block := CreateBlock(prevBlock, data)
	
	if validBlock(block, prevBlock) {
		bc.blocks = append(bc.blocks, block)
	}
}

func GenesisBlock() *Block {
	return CreateBlock(&Block, BookCheckout(IsGenesis: true))
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}

func validBlock(block, prevBlock *Block) bool {
	if prevBlock.Hash != block.PrevHash {
		return false
	}
	if !block.validateHash(block.Hash) {
		return false
	}
	if prevBlock.Pos + 1 != block.Pos {
		return false
	}
	return true
}

func (b *Block) validateHash(hash string) bool {
	b.generateHash()
	if b.Hash != hash {
		return false
	}
	return true
}

func (b *Block) generateHash() {
	bytes, _ := json.Marshal(b.Data) //get json encoding, drop error
	data := string(b.Pos) + b.Timestamp + string(bytes) + b.PrevHash
	hash := sha256.New()
	hash.Write([]byte(data))
	b.Hash = hex.EncodeToString(hash.Sum(nil))
}

func CreateBlock(prevBlock *Block, checkoutItem BookCheckout) *Block {
	block := &Block{}
	block.Pos = prevBlock.Pos + 1
	block.Timestamp = time.Now().String()
	block.Data = checkoutItem
	block.PrevHash = prevBlock.Hash
	block.generateHash()

	return block
}

func main() {
	BlockChain = NewBlockchain()

	fmt.Println("HELLO")
}