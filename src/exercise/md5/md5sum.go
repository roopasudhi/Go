package md5

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
)

func ReadData() ([]byte, error) {
	args := len(os.Args)
	if args <= 1 {
		fmt.Println("Reading from Stdin")
		fi, _ := os.Stdin.Stat()
		if fi.Size() != 0 {
			return ioutil.ReadAll(os.Stdin)
		}

	}
	if args > 1  {
		fmt.Println("Reading from file:", os.Args[1])
		return ioutil.ReadFile(os.Args[1])
	}
	_, filename, _, _ := runtime.Caller(1)
	fmt.Println("Reading from file:", filename)
	return ioutil.ReadFile(filename)
}

func Md5sum(content []byte) string {
	h := md5.New()
	h.Write(content)
	return hex.EncodeToString(h.Sum(nil))
}




