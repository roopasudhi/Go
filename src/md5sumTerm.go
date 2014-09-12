package main

import (
	"exercise/md5"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
)

func print(content []byte, err error) {
	if err != nil {
		fmt.Errorf("ERROR:", err)
	}
	fmt.Println("Hash:", md5.Md5sum(content))
}

func printMD5() {
	fi, _ := os.Stdin.Stat()
	s := fi.Size()
	switch args := len(os.Args); {
	case args <= 1 && s != 0:
		{
			fmt.Println("Reading from Stdin")
			print(ioutil.ReadAll(os.Stdin))

		}
	case args > 1:
		{
			fmt.Println("Reading from file:", os.Args[1])
			print(ioutil.ReadFile(os.Args[1]))
		}
	default:
		{
			_, filename, _, _ := runtime.Caller(1)
			fmt.Println("Reading from file:", filename)
			print(ioutil.ReadFile(filename))
		}

	}
}
/*
func main() {
	printMD5()
}
*/