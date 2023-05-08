package main

import (
    "crypto/sha256"
    "fmt"
    "github.com/cutajarj/ConcurrentProgrammingWithGo/chapter10/listing10.1"
    "os"
    "path/filepath"
)

func main() {
    dir := os.Args[1]
    files, _ := os.ReadDir(dir)
    sha := sha256.New()
    var prev, next chan int
    for _, file := range files {
        if !file.IsDir() {
            next = make(chan int)
            go func(filename string, prev, next chan int) {
                fpath := filepath.Join(dir, filename)
                hashOnFile := listing10_1.FHash(fpath)
                if prev != nil {
                    <-prev
                }
                sha.Write(hashOnFile)
                next <- 0
            }(file.Name(), prev, next)
            prev = next
        }
    }
    <-next
    fmt.Printf("%x\n", sha.Sum(nil))
}
