package main

import (
	"exercise/md5"
	"io/ioutil"
	"net/http"
)

const (
	POST = "POST"
)

func md5Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == POST {
		buf, _ := ioutil.ReadAll(r.Body)
		hash := md5.Md5sum(buf)
		w.Write([]byte(hash))
	}
}

func main() {
	http.HandleFunc("/", md5Handler)
	http.ListenAndServe(":9090", nil)
}
