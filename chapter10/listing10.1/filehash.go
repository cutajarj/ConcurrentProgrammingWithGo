package listing10_1

import (
    "crypto/sha256"
    "io"
    "os"
)

func FHash(filepath string) []byte {
    file, _ := os.Open(filepath)
    defer file.Close()

    sha := sha256.New()
    io.Copy(sha, file)

    return sha.Sum(nil)
}
