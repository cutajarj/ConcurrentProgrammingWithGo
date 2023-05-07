package main

import (
    "fmt"
    "io"
    "net/http"
    "strings"
)

func main() {
    const pagesToDownload = 30
    linesOnPage := make(chan int)
    finalResult := make(chan int)
    for i := 1000; i < 1000 + pagesToDownload; i++ {
        go func(id int) {
            url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", id)
            fmt.Println("Downloading", url)
            resp, _ := http.Get(url)
            if resp.StatusCode != 200 {
                panic("Server's error: " + resp.Status)
            }
            bodyBytes, _ := io.ReadAll(resp.Body)
            linesOnPage <- strings.Count(string(bodyBytes), "\n")
            resp.Body.Close()
        }(i)
    }
    go func() {
        totalLines := 0
        for i := 0; i < pagesToDownload; i++ {
            totalLines += <-linesOnPage
        }
        finalResult <- totalLines
    }()
    fmt.Println("Total lines:", <-finalResult)
}

