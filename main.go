package main

import (
	"catBlockChain/model"
	"fmt"
)

func main() {
	bc := model.NewBlockChain()
	bc.AddBlock("A send B 1BTC")
	bc.AddBlock("B send C 1BTC")

	for _, b := range bc.Blocks {
		fmt.Printf("TimeStamp: %d \tVerson: %d \tPrevBlockHash: %x \tBlockHash: %x\tBit: %d \tNonce: %d \tData: %s\n", b.TimeStamp, b.Version, b.PrevBlockHash, b.Hash, b.Bits, b.Nonce, b.Data)
	}
}
