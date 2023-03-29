package main

import (
    "fmt"
    "os"
    "path/filepath"
    "time"
)

func handleDirectories(dirs <-chan string, files chan<- string) {
    filesToPush := make([]string, 0)
    appendAllFiles := func(fullpath string) {
        fmt.Println("Reading all files from", fullpath)
        filesInDir, _ := os.ReadDir(fullpath)
        fmt.Printf("Pushing %d files from %s\n", len(filesInDir), fullpath)
        for _, file := range filesInDir {
            filesToPush = append(filesToPush, filepath.Join(fullpath, file.Name()))
        }
    }
    for {
        if len(filesToPush) == 0 {
            appendAllFiles(<-dirs)
        } else {
            select {
            case fullpath := <-dirs:
                appendAllFiles(fullpath)
            case files <- filesToPush[0]:
                filesToPush = filesToPush[1:]
            }
        }
    }
}

func handleFiles(files <-chan string, dirs chan<- string) {
    dirsToPush := make([]string, 0)
    appendAnyDirs := func(path string) {
        file, _ := os.Open(path)
        fileInfo, _ := file.Stat()
        if fileInfo.IsDir() {
            fmt.Printf("Pushing %s directory\n", fileInfo.Name())
            dirsToPush = append(dirsToPush, path)
        } else {
            fmt.Printf("File %s, size: %.2fKB, last modified: %s\n",
                fileInfo.Name(), float64(fileInfo.Size())/1024.0,
                fileInfo.ModTime().Format(time.ANSIC))
        }
    }
    for {
        if len(dirsToPush) == 0 {
            appendAnyDirs(<-files)
        } else {
            select {
            case f := <- files:
                appendAnyDirs(f)
            case dirs <- dirsToPush[0]:
                dirsToPush = dirsToPush[1:]
            }
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
