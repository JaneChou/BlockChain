package BlockChain

import (
	"fmt"
	"mutexTest/Chain/Block"
)

type Node struct {
	NextNode *Node
	Data *Block.Block
}

func CreateHeadNode(data *Block.Block) *Node {
	//创建头结点
	headNode := new(Node)
	headNode.NextNode = nil
	headNode.Data = data

	return headNode
}

// 添加结点
// 挖矿成功，添加区块
func AddNode(data *Block.Block, node *Node) *Node {
	//创建结点
	newNode := new(Node)

	newNode.NextNode = nil
	newNode.Data = data

	node.NextNode = newNode

	return newNode
}

// 查看数据
func ShowNode(node *Node)  {
	n := node
	for {
		// 没有下一个结点
		if n.NextNode == nil {
			fmt.Println(n.Data)
			break
		} else {
			fmt.Println(n.Data)
			n = n.NextNode
		}
	}
}

