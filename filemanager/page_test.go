package filemanager

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSetAndGetInt(t *testing.T) {
	page := NewPageBySize(256)
	val := uint64(3456)
	page.SetInt(30, val)
	valGot := page.GetInt(30)

	require.Equal(t, valGot, val)
}

func TestSetAndGetBytes(t *testing.T) {
	b := []byte{'a', 'b', 'c', 'd'}
	page := NewPageBySize(256)
	page.SetBytes(52, b)
	bGot := page.GetBytes(52)

	require.Equal(t, b, bGot)
}

func TestSetAndGetString(t *testing.T) {
	s := "hi，数据库"
	page := NewPageBySize(256)
	page.SetString(71, s)
	sGot := page.GetString(71)

	require.Equal(t, s, sGot)
}

func TestMaxLengthForString(t *testing.T) {
	s := "hi，数据库"
	sLength := uint64(len(s))
	page := NewPageBySize(256)

	maxLength := page.MaxLengthForString(s)

	require.Equal(t, maxLength, sLength+8)
}
