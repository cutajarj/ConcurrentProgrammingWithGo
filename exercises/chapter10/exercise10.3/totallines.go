package main

import (
    "fmt"
    "io"
    "net/http"
    "strings"
)

func main() {
    const pagesToDownload = 30
    totalLines := 0
    for i := 1000; i < 1000 + pagesToDownload; i++ {
        url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
        fmt.Println("Downloading", url)
        resp, _ := http.Get(url)
        if resp.StatusCode != 200 {
            panic("Server's error: " + resp.Status)
        }
        bodyBytes, _ := io.ReadAll(resp.Body)
        totalLines += strings.Count(string(bodyBytes), "\n")
        resp.Body.Close()
    }
    fmt.Println("Total lines:", totalLines)
}

