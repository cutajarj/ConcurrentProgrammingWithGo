package main

import (
    "fmt"
    "os"
    "path/filepath"
    "time"
)

func handleDirectories(dirs <-chan string, files chan<- string) {
    toPush := make([]string, 0)
    appendAllFiles := func(path string) {
        fmt.Println("Reading all files from", path)
        filesInDir, _ := os.ReadDir(path)
        fmt.Printf("Pushing %d files from %s\n", len(filesInDir), path)
        for _, f := range filesInDir {
            toPush = append(toPush, filepath.Join(path, f.Name()))
        }
    }
    for {
        if len(toPush) == 0 {
            appendAllFiles(<-dirs)
        } else {
            select {
            case fullpath := <-dirs:
                appendAllFiles(fullpath)
            case files <- toPush[0]:
                toPush = toPush[1:]
            }
        }
    }
}

func handleFiles(files <-chan string, dirs chan<- string) {
    for path := range files {
        file, _ := os.Open(path)
        fileInfo, _ := file.Stat()
        if fileInfo.IsDir() {
            fmt.Printf("Pushing %s directory\n", fileInfo.Name())
            dirs <- path
        } else {
            fmt.Printf("File %s, size: %.2fKB, last modified: %s\n",
                fileInfo.Name(), float64(fileInfo.Size()) / 1024.0,
                fileInfo.ModTime().Format(time.ANSIC))
        }
    }
}

func main() {
    filesChannel := make(chan string)
    dirsChannel := make(chan string)
    go handleFiles(filesChannel, dirsChannel)
    go handleDirectories(dirsChannel, filesChannel)
    dirsChannel <- os.Args[1]
    time.Sleep(60 * time.Second)
}
