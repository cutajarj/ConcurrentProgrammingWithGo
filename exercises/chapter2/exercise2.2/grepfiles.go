package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

/*
  You can run this by executing:
  go run grepfiles.go Mozilla ../../commonfiles/txtfile1 ../../commonfiles/txtfile2 ../../commonfiles/txtfile3
*/

func grepFile(filename string, searchStr string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	if strings.Contains(string(content), searchStr) {
		fmt.Println(filename, "contains a match with", searchStr)
	} else {
		fmt.Println(filename, "does NOT contain a match with", searchStr)
	}
}

func main() {
	searchStr := os.Args[1]
	filenames := os.Args[2:]
	for _, filename := range filenames {
		go grepFile(filename, searchStr)
	}
	time.Sleep(2 * time.Second)
}
