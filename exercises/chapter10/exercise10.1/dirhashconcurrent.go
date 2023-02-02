package main

import (
    "crypto/md5"
    "fmt"
    "github.com/cutajarj/ConcurrentProgrammingWithGo/chapter10/listing10.1"
    "os"
    "path/filepath"
    "sync"
)

func main() {
    dir := os.Args[1]
    files, _ := os.ReadDir(dir)
    hMd5 := md5.New()
    var prev, next sync.WaitGroup
    for _, file := range files {
        if !file.IsDir() {
            next = sync.WaitGroup{}
            go func(filename string, prev, next *sync.WaitGroup) {
                fpath := filepath.Join(dir, filename)
                hashOnFile := listing10_1.FHash(fpath)
                if prev != nil {
                    prev.Wait()
                }
                hMd5.Write(hashOnFile)
                next.Done()
            }(file.Name(), &prev, &next)
            prev = next
        }
    }
    next.Wait()
    fmt.Printf("%s - %x\n", dir, hMd5.Sum(nil))
}
