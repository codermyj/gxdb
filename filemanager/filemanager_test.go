package filemanager

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFileManager(t *testing.T) {
	fm, _ := NewFileManager("file_test", 400)

	blk := NewBlockID("testfile", 2)
	p1 := NewPageBySize(fm.BlockSize())
	post1 := uint64(88)
	valStr := "abcdefghijklmn"
	p1.SetString(post1, valStr)
	size := p1.MaxLengthForString(valStr)
	valInt := uint64(233)
	post2 := post1 + size
	p1.SetInt(post2, valInt)

	fm.Write(blk, p1)

	p2 := NewPageBySize(fm.BlockSize())
	fm.Read(blk, p2)

	require.Equal(t, valStr, p2.GetString(post1))
	require.Equal(t, valInt, p2.GetInt(post2))

}
