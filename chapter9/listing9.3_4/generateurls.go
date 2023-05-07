package main

import "fmt"

func generateUrls(quit <-chan int) <-chan string {
    urls := make(chan string)
    go func() {
        defer close(urls)
        for i := 100; i <= 130; i++ {
            url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
            select {
            case urls <- url:
            case <-quit:
                return
            }
        }
    }()
    return urls
}

func main() {
    quit := make(chan int)
    defer close(quit)
    results := generateUrls(quit)
    for result := range results {
        fmt.Println(result)
    }
}
