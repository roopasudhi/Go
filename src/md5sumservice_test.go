package main

import (
    "testing"
    "net/http"
    "strings"
    "io/ioutil"
    "exercise/md5"
)

func TestPost(t *testing.T) {
	service := NewMd5Service(":9091")
	go service.ListenAndServe()
	
	//Make put request
	c := &http.Client{}
	rq, _ := http.NewRequest("POST", "http://localhost:9091/md5/", strings.NewReader("test:chicago"))
	r,_ :=c.Do(rq)
	body, _ := ioutil.ReadAll(r.Body)
	//Verify
	if string(body) != md5.Md5sum([]byte("test:chicago")) {
		t.Error("Add Not Successful",string(body))
	}
	service.Stop()
	
}

