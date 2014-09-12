package main

import (
	"exercise"
	"exercise/md5"
	"io/ioutil"
	"net/http"
)


func Md5Handler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		content, _ := ioutil.ReadAll(r.Body)

		switch r.Method {
		case exercise.POST:
			{
				hash := md5.Md5sum(content)
				w.Write([]byte(hash))
			}
		}
	})
}
func NewMd5Service(port string) *exercise.HttpServer {
	server := exercise.NewHttpServer(port)
	server.AddHandler("/md5/", Md5Handler())
	return server
}

/*
func main() {
	http.HandleFunc("/", md5Handler)
	http.ListenAndServe(":9090", nil)
}
*/
