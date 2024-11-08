package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	stat, _ := filepath.Abs(filepath.Dir("./main.go"))

	filepath.Join()

	fmt.Println(stat)

}
