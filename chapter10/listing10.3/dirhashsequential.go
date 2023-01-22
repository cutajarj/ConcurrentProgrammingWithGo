package main

import (
	"crypto/md5"
	"fmt"
	"github.com/cutajarj/ConcurrentProgrammingWithGo/chapter10/listing10.1"
	"os"
	"path/filepath"
)

func main() {
	dir := os.Args[1]
	files, _ := os.ReadDir(dir)
	hMd5 := md5.New()
	for _, file := range files {
		if !file.IsDir() {
			fpath := filepath.Join(dir, file.Name())
			hashOnFile := listing10_1.FHash(fpath)
			hMd5.Write(hashOnFile)
		}
	}
	fmt.Printf("%s - %x\n", dir, hMd5.Sum(nil))
}
