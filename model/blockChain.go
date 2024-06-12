package model

type BlockChain struct {
	Blocks []*Block
}

func NewBlockChain() *BlockChain {
	return &BlockChain{
		Blocks: []*Block{NewGenesisBlock()},
	}
}

func (m *BlockChain) AddBlock(data string) {
	newBlock := NowBlock(m.Blocks[len(m.Blocks)-1].Hash, data)
	m.Blocks = append(m.Blocks, newBlock)
}
