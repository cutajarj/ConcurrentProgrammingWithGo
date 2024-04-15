package main

import (
    "crypto/md5"
    "fmt"
    "github.com/cutajarj/ConcurrentProgrammingWithGo/chapter10/listing10.1"
    "os"
    "path/filepath"
    "sync"
)

// WaitGroup wrapper to store a WaitGroup pointer.
// This is to bypass WaitGroup's nocopy restriction.
// https://pkg.go.dev/sync#WaitGroup
type WgWrapper struct {
	Wg *sync.WaitGroup
}

func main() {
    dir := os.Args[1]
    files, _ := os.ReadDir(dir)
    hMd5 := md5.New()
    var prev, next WgWrapper
    for _, file := range files {
        if !file.IsDir() {
            next = WgWrapper{
                Wg: &sync.WaitGroup{},
            }
            next.Wg.Add(1)
            go func(filename string, prev, next WgWrapper) {
                fpath := filepath.Join(dir, filename)
                fmt.Println("Processing", fpath)
                hashOnFile := listing10_1.FHash(fpath)
                // If not the first iteration
                if prev.Wg != nil {
                    prev.Wg.Wait()
                }
                hMd5.Write(hashOnFile)
                next.Wg.Done()
            }(file.Name(), prev, next)
            prev = next
        }
    }
    next.Wg.Wait()
    fmt.Printf("%s - %x\n", dir, hMd5.Sum(nil))
}
