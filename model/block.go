package model

import (
	"bytes"
	"catBlockChain/utils"
	"crypto/sha256"
	"time"
)

type Block struct {
	Version       int64
	TimeStamp     int64
	Hash          []byte // 当前区块的hash值，简化
	PrevBlockHash []byte // 前区块的hash值
	MerKelRoot    []byte // 梅克尔根
	Bits          int64  // 难度值
	Nonce         int64  // 随机值

	Data []byte // 交易信息
}

func NowBlock(prevBlockHash []byte, data string) (block *Block) {
	block = &Block{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		//Hash: TODO,
		MerKelRoot: []byte{},
		TimeStamp:  time.Now().Unix(),
		Bits:       1,
		Nonce:      1,
		Data:       []byte(data),
	}
	block.SetHash()
	return
}

func (m *Block) SetHash() {
	tmp := [][]byte{
		utils.Int2Byte(m.Version),
		utils.Int2Byte(m.TimeStamp),
		m.PrevBlockHash,
		m.MerKelRoot,
		utils.Int2Byte(m.Bits),
		utils.Int2Byte(m.Nonce),
		m.Data,
	}
	data := bytes.Join(tmp, []byte{})
	hash := sha256.Sum256(data)
	m.Hash = hash[:]
}

func NewGenesisBlock() *Block {
	return NowBlock([]byte{}, "Genesis Block!")
}
