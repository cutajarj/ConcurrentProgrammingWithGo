package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

func fileSearch(dir string, filename string,
	wg *sync.WaitGroup, mutex *sync.Mutex, matches *[]string) {
	files, _ := os.ReadDir(dir)
	for _, file := range files {
		fpath := filepath.Join(dir, file.Name())
		if strings.Contains(file.Name(), filename) {
			mutex.Lock()
			*matches = append(*matches, fpath)
			mutex.Unlock()
		}
		if file.IsDir() {
			wg.Add(1)
			go fileSearch(fpath, filename, wg, mutex, matches)
		}
	}
	wg.Done()
}

func main() {
	results := make([]string, 0)
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go fileSearch(os.Args[1], os.Args[2], &wg, &mutex, &results)
	wg.Wait()
	mutex.Lock()
	sort.Strings(results)
	fmt.Println(strings.Join(results, "\n"))
	mutex.Unlock()
}
