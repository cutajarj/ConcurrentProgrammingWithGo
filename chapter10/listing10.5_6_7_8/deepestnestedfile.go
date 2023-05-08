package main

import (
    "fmt"
    "math"
    "os"
    "path/filepath"
    "strings"
    "sync"
)

type CodeDepth struct {
    file  string
    level int
}

func deepestNestedBlock(filename string) CodeDepth {
    code, _ := os.ReadFile(filename)
    max := 0
    level := 0
    for _, c := range code {
        if c == '{' {
            level += 1
            max = int(math.Max(float64(max), float64(level)))
        } else if c == '}' {
            level -= 1
        }
    }
    return CodeDepth{filename, max}
}

func forkIfNeeded(path string, info os.FileInfo,
    wg *sync.WaitGroup, results chan CodeDepth) {
    if !info.IsDir() && strings.HasSuffix(path, ".go") {
        wg.Add(1)
        go func() {
            results <- deepestNestedBlock(path)
            wg.Done()
        }()
    }
}

func joinResults(partialResults chan CodeDepth) chan CodeDepth {
    finalResult := make(chan CodeDepth)
    max := CodeDepth{"", 0}
    go func() {
        for pr := range partialResults {
            if pr.level > max.level {
                max = pr
            }
        }
        finalResult <- max
    }()
    return finalResult
}

func main() {
    dir := os.Args[1]
    partialResults := make(chan CodeDepth)
    wg := sync.WaitGroup{}
    filepath.Walk(dir,
        func(path string, info os.FileInfo, err error) error {
            forkIfNeeded(path, info, &wg, partialResults)
            return nil
        })
    finalResult := joinResults(partialResults)
    wg.Wait()
    close(partialResults)
    result := <-finalResult
    fmt.Printf("%s has the deepest nested code block of %d\n",
        result.file, result.level)
}
