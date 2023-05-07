package main

import (
    "fmt"
    "io"
    "net/http"
    "regexp"
    "strings"
)

func extractWords(quit <-chan int, pages <-chan string) <-chan string {
    words := make(chan string)
    go func() {
        defer close(words)
        wordRegex := regexp.MustCompile(`[a-zA-Z]+`)
        moreData, pg := true, ""
        for moreData {
            select {
            case pg, moreData = <-pages:
                if moreData {
                    for _, word := range wordRegex.FindAllString(pg, -1) {
                        words <- strings.ToLower(word)
                    }
                }
            case <-quit:
                return
            }
        }
    }()
    return words
}

func downloadPages(quit <-chan int, urls <-chan string) <-chan string {
    pages := make(chan string)
    go func() {
        defer close(pages)
        moreData, url := true, ""
        for moreData {
            select {
            case url, moreData = <-urls:
                if moreData {
                    resp, _ := http.Get(url)
                    if resp.StatusCode != 200 {
                        panic("Server's error: " + resp.Status)
                    }
                    body, _ := io.ReadAll(resp.Body)
                    pages <- string(body)
                    resp.Body.Close()
                }
            case <-quit:
                return
            }
        }
    }()
    return pages
}

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
    results := extractWords(quit, downloadPages(quit, generateUrls(quit)))
    for result := range results {
        fmt.Println(result)
    }
}
