package filemanager

import "encoding/binary"

// Page 内存页
type Page struct {
	buffer []byte // 页中的内容
}

// NewPageBySize 创建已知大小的页
func NewPageBySize(blockSize uint64) *Page {
	bytes := make([]byte, blockSize)
	return &Page{
		buffer: bytes,
	}
}

// NewPageByBytes 创建已知内容的页
func NewPageByBytes(bytes []byte) *Page {
	return &Page{
		buffer: bytes,
	}
}

// GetInt 从已知位置开始获取1个64位无符号整数
func (p *Page) GetInt(offset uint64) uint64 {
	num := binary.LittleEndian.Uint64(p.buffer[offset : offset+8])
	return num
}

// uint64ToByteArray 工具函数,将64位无符号整数转换为字节数组（小端法）
func uint64ToByteArray(num uint64) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, num)
	return bytes
}

// SetInt 在某个位置设置1个整数
func (p *Page) SetInt(offset uint64, num uint64) {
	bytes := uint64ToByteArray(num)
	copy(p.buffer[offset:], bytes)
}

// GetBytes 获取某个位置的byte数组
func (p *Page) GetBytes(offset uint64) []byte {
	length := binary.LittleEndian.Uint64(p.buffer[offset : offset+8]) // 获取byte数据的长度
	buf := make([]byte, length)
	copy(buf, p.buffer[offset+8:offset+8+length]) // 根据长度length获取byte数据
	return buf
}

// SetBytes 在某个位置设置byte数组
func (p *Page) SetBytes(offset uint64, b []byte) {
	length := uint64(len(b))     // 获取长度
	p.SetInt(offset, length)     // 将长度数据设置到page中
	copy(p.buffer[offset+8:], b) // 将数据内容设置到数据中
}

// GetString 在某个位置读取1个字符串
func (p *Page) GetString(offset uint64) string {
	b := p.GetBytes(offset)
	return string(b)
}

// SetString 将某个位置内容设置为字符串
func (p *Page) SetString(offset uint64, val string) {
	b := []byte(val)
	p.SetBytes(offset, b)
}

// MaxLengthForString 返回字符串数据的整个长度（length所占的长度+实际数据占用的字节长度）
func (p *Page) MaxLengthForString(s string) uint64 {
	lengthSize := uint64(8)          // 长度数所占的字节长度，固定为8
	strLen := uint64(len([]byte(s))) // 实际字符串数据所占的长度
	return lengthSize + strLen
}

// contents 返回page缓冲区内容
func (p *Page) contents() []byte {
	return p.buffer
}
