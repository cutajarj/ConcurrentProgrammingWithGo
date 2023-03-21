package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

/*
  You can run this by executing:
  go run grepdirrec.go Mozilla ../../commonfiles
*/

func grepPath(path string, dirEntry os.DirEntry, searchStr string) {
	fullPath := filepath.Join(path, dirEntry.Name())
	if dirEntry.IsDir() {
		files, err := os.ReadDir(fullPath)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			go grepPath(fullPath, file, searchStr)
		}
	} else {
		content, err := os.ReadFile(fullPath)
		if err != nil {
			log.Fatal(err)
		}
		if strings.Contains(string(content), searchStr) {
			fmt.Println(fullPath, "contains a match with ", searchStr)
		} else {
			fmt.Println(fullPath, "does NOT contain a match with", searchStr)
		}
	}
}

func main() {
	searchStr := os.Args[1]
	path := os.Args[2]
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, dirEntry := range files {
		go grepPath(path, dirEntry, searchStr)
	}
	time.Sleep(2 * time.Second)
}
