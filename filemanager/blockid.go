package filemanager

import (
	"crypto/sha256"
	"fmt"
)

// BlockId 表示二进制文件中的一块内容
type BlockId struct {
	fileName string // 对应的磁盘中的二进制文件
	blkNum   uint64 // 二进制文件中区块的编号
}

// NewBlockID BlockId构造函数
func NewBlockID(fileName string, blkNum uint64) *BlockId {
	return &BlockId{
		fileName: fileName,
		blkNum:   blkNum,
	}
}

// FileName 返回区块所在的文件名
func (b *BlockId) FileName() string {
	return b.fileName
}

// Number 返回区块编号
func (b *BlockId) Number() uint64 {
	return b.blkNum
}

// Equal 判断两个区块是否相等
func (b *BlockId) Equal(other *BlockId) bool {
	return b.fileName == other.fileName && b.blkNum == other.blkNum
}

// asSha256 计算区块的sha256值
func asSha256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// HashCode 返回区块的哈希值
func (b *BlockId) HashCode() string {
	return asSha256(*b)
}
