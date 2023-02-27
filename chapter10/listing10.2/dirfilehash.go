package main

import (
    "fmt"
    "github.com/cutajarj/ConcurrentProgrammingWithGo/chapter10/listing10.1"
    "os"
    "path/filepath"
    "sync"
)

func main() {
    dir := os.Args[1]
    files, _ := os.ReadDir(dir)
    wg := sync.WaitGroup{}
    for _, file := range files {
        if !file.IsDir() {
            wg.Add(1)
            go func(filename string) {
                fPath := filepath.Join(dir, filename)
                hash := listing10_1.FHash(fPath)
                fmt.Printf("%s - %x\n", filename, hash)
                wg.Done()
            }(file.Name())
        }
    }
    wg.Wait()
}
