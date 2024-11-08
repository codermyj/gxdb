package main

import (
	fm "filemanager"
)

const (
	UINT64_LEN = 8
)

type LogManager struct {
	fileManager *fm.FileManager
	logFile     string
	logPage     *fm.Page
}

func main() {

}
