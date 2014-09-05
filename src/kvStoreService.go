package main

import (
	"exercise/kv"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var store kv.Store

func init() {
	store = kv.NewStore()
}

// Default Request Handler
func crudHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	content, _ := ioutil.ReadAll(r.Body)

	path := r.URL.Path
	pos := strings.LastIndex(path, "/")
	pos++
	id := path[pos:]

	switch method {
	case "PUT":
		{
			kvStr := strings.Split(string(content), ":")
			store.Put(kvStr[0], kvStr[1])
		}
	case "GET":
		{
			v, _ := store.Get(id)
			w.Write([]byte(v))
		}
	case "DELETE":
		{
			store.Delete(id)
		}
	}
}

// Default Request Handler
func countHandler(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	pos := strings.LastIndex(path, "/")
	pos++
	id := path[pos:]

	switch r.Method {
	case "GET":
		{
			v := store.Count(id)
			w.Write([]byte(strconv.Itoa(v)))
		}

	}
}
func startKVService(port string) {
	http.HandleFunc("/kv/", crudHandler)
	http.HandleFunc("/kv/count/", countHandler)
	http.ListenAndServe(port, nil)
}

func main() {
	startKVService(":9090")
	fmt.Println("Started on 9090")
}
