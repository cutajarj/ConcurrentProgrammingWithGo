package listing10_1

import (
	"crypto/md5"
	"io"
	"os"
)

func FHash(filepath string) []byte {
	file, _ := os.Open(filepath)
	defer file.Close()

	hMd5 := md5.New()
	io.Copy(hMd5, file)

	return hMd5.Sum(nil)
}
