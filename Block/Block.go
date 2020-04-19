package Block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	// 上一个区块哈希
	PreHash string
	// 当前区块哈希
	HashCode string
	// 时间戳
	TimeStamp string
	// 当前网络的难度系数
	Diff int
	// 存交易信息
	Data string
	// 区块高度
	Index int
	// 随机值
	Nonce int
}

// 创建创世区块
func GenerateFirstBlock(data string) Block{
	//创建第一个区块
	var firstBlock Block
	firstBlock.PreHash = "0"
	firstBlock.TimeStamp = time.Now().String()
	// 前导0的个数
	firstBlock.Diff = 4
	firstBlock.Data = data
	firstBlock.Index = 1
	firstBlock.Nonce = 0

	// 当前块的哈希
	firstBlock.HashCode = GenerateHashValue(firstBlock)

	return firstBlock
}

// 生成区块的哈希
func GenerateHashValue(block Block) string {
	var hashdata = strconv.Itoa(block.Index) + strconv.Itoa(block.Nonce) +
		strconv.Itoa(block.Diff) + block.TimeStamp

	var sha = sha256.New()
	sha.Write([]byte(hashdata))
	hashed := sha.Sum(nil)

	return hex.EncodeToString(hashed)
}

// 创建新区块
func GenerateNextBlock(data string, prevBlock Block) Block {
	var newBlock Block
	newBlock.PreHash = prevBlock.HashCode
	newBlock.TimeStamp = time.Now().String()
	// 前导0的个数
	newBlock.Diff = 4
	newBlock.Data = data
	newBlock.Index = prevBlock.Index + 1
	newBlock.Nonce = 0

	// pow挖矿
	newBlock.HashCode = pow(newBlock.Diff, &newBlock)

	return newBlock
}

// pow工作量证明算法进行哈希碰撞
func pow(diff int, block *Block) string {
	// 不停地挖矿
	for {
		hash := GenerateHashValue(*block)
		// 每挖矿依次，打印哈希值
		fmt.Println(hash)
		// 判断哈希值前缀是否是4个0
		if strings.HasPrefix(hash, strings.Repeat("0", diff)) {
			fmt.Println("挖矿成功")
			return hash
		} else {
			block.Nonce++
		}
	}
}