package main 

import (
	"exercise/md5"
	"fmt"
)

func main() {
	content,err := md5.ReadData()
	if err!=nil{
		fmt.Errorf("ERROR:",err)
	} 
	fmt.Println("Hash:",md5.Md5sum(content))
}

