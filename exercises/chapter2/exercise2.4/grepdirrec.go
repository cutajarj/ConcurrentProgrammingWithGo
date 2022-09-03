package main

import (
    "io/ioutil"
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

func grepPath(path string, fileInfo os.FileInfo, searchStr string) {
    fullPath := filepath.Join(path, fileInfo.Name())
    if fileInfo.IsDir() {
        files, err := ioutil.ReadDir(fullPath)
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
            println(fullPath, "contains a match with ", searchStr)
        } else {
            println(fullPath, "does NOT contain a match with", searchStr)
        }
    }
}

func main() {
    searchStr := os.Args[1]
    path := os.Args[2]
    files, err := ioutil.ReadDir(path)
    if err != nil {
        log.Fatal(err)
    }
    for _, fileInfo := range files {
        go grepPath(path, fileInfo, searchStr)
    }
    time.Sleep(2 * time.Second)
}
