package filemanager

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// FileManager 文件管理
type FileManager struct {
	dbDirectory string              // 文件目录路径
	blockSize   uint64              // 单个块大小
	isNew       bool                // 是否新建
	openFiles   map[string]*os.File // 已打开的文件
	mu          sync.Mutex          // 互斥锁
}

// NewFileManager 创建文件管理对象
func NewFileManager(dbDirectory string, blockSize uint64) (*FileManager, error) {
	fileManager := FileManager{
		dbDirectory: dbDirectory,
		blockSize:   blockSize,
		isNew:       false,
		openFiles:   make(map[string]*os.File),
	}

	if _, err := os.Stat(dbDirectory); os.IsNotExist(err) {
		// 目录不存在则生成
		fileManager.isNew = true
		err = os.Mkdir(dbDirectory, os.ModeDir)
		if err != nil {
			return nil, err
		}
	} else {
		// 如果目录已存在，则将目录下的文件删除
		err = filepath.Walk(dbDirectory, func(path string, info fs.FileInfo, err error) error {
			mode := info.Mode()
			if mode.IsRegular() {
				name := info.Name()
				if strings.HasPrefix(name, "tmp") {
					os.Remove(filepath.Join(path, name))
				}
			}
			return nil
		})

		if err != nil {
			return nil, err
		}
	}
	return &fileManager, nil
}

// getFile 获取文件对象
func (f *FileManager) getFile(fileName string) (*os.File, error) {
	path := filepath.Join(f.dbDirectory, fileName)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	f.openFiles[fileName] = file
	return file, nil
}

// 把打开文件中对应区块的内容读取到内存Page中
func (f *FileManager) Read(blk *BlockId, p *Page) (int, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	file, err := f.getFile(blk.fileName)
	if err != nil {
		return 0, nil
	}
	defer file.Close()

	count, err := file.ReadAt(p.contents(), int64(blk.blkNum*f.blockSize))
	if err != nil {
		return 0, err
	}

	return count, nil
}

// Write 把页中的内容写入到指定的块中
func (f *FileManager) Write(blk *BlockId, p *Page) (int, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	file, err := f.getFile(blk.fileName)
	if err != nil {
		return 0, nil
	}
	defer file.Close()

	count, err := file.WriteAt(p.contents(), int64(blk.blkNum*f.blockSize))
	if err != nil {
		return 0, err
	}

	return count, nil
}

// size 获取块的个数
func (f *FileManager) size(fileName string) (uint64, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	file, err := f.getFile(fileName)
	if err != nil {
		return 0, nil
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return 0, err
	}

	return uint64(info.Size()) / f.blockSize, nil
}

// Append 在文件fileName后面增加一个区块
func (f *FileManager) Append(fileName string) (BlockId, error) {
	newBlockNum, err := f.size(fileName)
	if err != nil {
		return BlockId{}, err
	}

	blk := NewBlockID(fileName, newBlockNum)
	file, err := f.getFile(fileName)
	if err != nil {
		return BlockId{}, err
	}
	defer file.Close()
	b := make([]byte, f.blockSize)
	_, err = file.WriteAt(b, int64(blk.blkNum*f.blockSize))
	if err != nil {
		return BlockId{}, err
	}

	return *blk, nil
}

// IsNew 判断FileManager对象是否为新建
func (f *FileManager) IsNew() bool {
	return f.isNew
}

// BlockSize 获取单个块的大小
func (f *FileManager) BlockSize() uint64 {
	return f.blockSize
}
