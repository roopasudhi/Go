package main

import (
	"exercise"
	"exercise/kv"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func getId(p string) string {
	pos := strings.LastIndex(p, "/")
	pos++
	id := p[pos:]
	return id
}

func CrudHandler(store kv.Store) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		content, _ := ioutil.ReadAll(r.Body)
		id := getId(r.URL.Path)

		switch r.Method {
		case exercise.PUT:
			{
				kvStr := strings.Split(string(content), ":")
				store.Put(kvStr[0], kvStr[1])
			}
		case exercise.GET:
			{
				v, _ := store.Get(id)
				w.Write([]byte(v))
			}
		case exercise.DELETE:
			{
				store.Delete(id)
			}
		}
	})
}

func CountHandler(store kv.Store) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		id := getId(r.URL.Path)

		switch r.Method {
		case exercise.GET:
			{
				v := store.Count(id)
				w.Write([]byte(strconv.Itoa(v)))
			}

		}
	})
}
func newKVService(port string) *exercise.HttpServer {
	store := kv.NewStore()
	server := exercise.NewHttpServer(port)
	server.AddHandler("/kv/", CrudHandler(store))
	server.AddHandler("/kv/count/", CountHandler(store))
	return server
}

/*
func main() {
	service := newKVService(":9090")
	service.ListenAndServe()
	fmt.Println("Started on 9090")
}
*/
