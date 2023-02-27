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
                hMd5.Write(hashOnFile)
                next <- 0
            }(file.Name(), prev, next)
            prev = next
        }
    }
    <-next
    fmt.Printf("%s - %x\n", dir, hMd5.Sum(nil))
}
