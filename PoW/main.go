package main

import (
	"mutexTest/Chain/Block"
	"mutexTest/Chain/BlockChain"
)

func main()  {
	first := Block.GenerateFirstBlock("创世区块")
	second := Block.GenerateNextBlock("第二个区块", first)

	header := BlockChain.CreateHeadNode(&first)
	BlockChain.AddNode(&second, header)

	BlockChain.ShowNode(header)
}