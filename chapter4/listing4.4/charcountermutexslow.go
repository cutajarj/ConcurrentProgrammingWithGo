package listing4_4

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

// CountLetters
// Note: this program us locking the entire goroutine with mutex on purpose to demonstrate
// bad placement of the lock and unlock. We fix this in the next listing
func CountLetters(url string, frequency []int, mutex *sync.Mutex) {
	mutex.Lock()
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	for _, b := range body {
		c := strings.ToLower(string(b))
		cIndex := strings.Index(allLetters, c)
		if cIndex >= 0 {
			frequency[cIndex] += 1
		}
	}
	fmt.Println("Completed:", url)
	mutex.Unlock()
}
