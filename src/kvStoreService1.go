package main

import (
	"exercise/kv"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"os"
)

const (
	GET    = "GET"
	PUT    = "PUT"
	POST   = "POST"
	DELETE = "DELETE"
)

func getId(p string) string {
	pos := strings.LastIndex(p, "/")
	pos++
	id := p[pos:]
	return id
}


func CrudHandler(store kv.Store) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		content, _ := ioutil.ReadAll(r.Body)
		id := getId(r.URL.Path)

		switch r.Method {
		case PUT:
			{
				kvStr := strings.Split(string(content), ":")
				store.Put(kvStr[0], kvStr[1])
			}
		case GET:
			{
				v, _ := store.Get(id)
				w.Write([]byte(v))
			}
		case DELETE:
			{
				store.Delete(id)
			}
		}
	})
}


func CountHandler(store kv.Store) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		id := getId(r.URL.Path)

		switch r.Method {
		case GET:
			{
				v := store.Count(id)
				w.Write([]byte(strconv.Itoa(v)))
			}

		}
	})
}
func startKVService(port string) {
	store := kv.NewStore()
	http.Handle("/kv/", CrudHandler(store))
	http.Handle("/kv/count/", CountHandler(store))
	http.ListenAndServe(port, nil)
}
func stop() {
	os.Exit(0)
}
func main() {
	startKVService(":9090")
	fmt.Println("Started on 9090")
}
