package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestKvService(t *testing.T) {
	service := newKVService(":9090")
	go service.ListenAndServe()
	time.Sleep(1 * time.Second)
	service.Stop()
	time.Sleep(1 * time.Second)
}

func TestKvServicePutOperation(t *testing.T) {
	//Start Service
	service := newKVService(":9090")
	go service.ListenAndServe()

	//Make put request
	c := &http.Client{}
	rq, _ := http.NewRequest("PUT", "http://localhost:9090/kv/", strings.NewReader("roopa:chicago"))
	c.Do(rq)
	//Make GET request
	r, err := http.Get("http://localhost:9090/kv/roopa")
	if err != nil {
		t.Error(err)
	}
	body, _ := ioutil.ReadAll(r.Body)
	//Verify
	if string(body) != "chicago" {
		t.Error("Add Not Successful")
	}
	//Stop Service
	service.Stop()
}

func TestKvServiceUpdateOperation(t *testing.T) {
	//Start Service
	service := newKVService(":9090")
	go service.ListenAndServe()

	//Make put request
	c := &http.Client{}
	rqu, _ := http.NewRequest("PUT", "http://localhost:9090/kv/", strings.NewReader("roopa:Illinois"))
	c.Do(rqu)
	//Make GET request
	ru, _ := http.Get("http://localhost:9090/kv/roopa")
	bu, _ := ioutil.ReadAll(ru.Body)
	//Verify
	if string(bu) != "Illinois" {
		t.Error("Update Not Successful")
	}
	//Stop Service
	service.Stop()
}

func TestKvServiceCountOperation(t *testing.T) {
	//Start Service
	service := newKVService(":9090")
	go service.ListenAndServe()

	//Make put request
	c := &http.Client{}
	rqu, _ := http.NewRequest("PUT", "http://localhost:9090/kv/", strings.NewReader("roopa:Illinois"))
	c.Do(rqu)
	rqc, _ := http.NewRequest("PUT", "http://localhost:9090/kv/", strings.NewReader("roopa_new:Chicago"))
	c.Do(rqc)
	rqc, _ = http.NewRequest("PUT", "http://localhost:9090/kv/", strings.NewReader("test:Chicago"))
	c.Do(rqc)
	rqc, _ = http.NewRequest("PUT", "http://localhost:9090/kv/", strings.NewReader("test1:Chicago"))
	c.Do(rqc)

	//Make GET request
	rc, _ := http.Get("http://localhost:9090/kv/count/")
	buc, _ := ioutil.ReadAll(rc.Body)
	//Verify

	if string(buc) != "4" {
		t.Error("Count Not Successful", string(buc))
	}

	rc, _ = http.Get("http://localhost:9090/kv/count/roopa")
	buc, _ = ioutil.ReadAll(rc.Body)
	if string(buc) != "2" {
		t.Error("Count Not Successful", string(buc))
	}

	//Stop Service
	service.Stop()
}

func TestKvServiceDeleteOperation(t *testing.T) {
	//Start Service
	c := &http.Client{}
	service := newKVService(":9090")
	go service.ListenAndServe()

	//Make PUT Request
	rqu, _ := http.NewRequest("PUT", "http://localhost:9090/kv/", strings.NewReader("roopa:Illinois"))
	c.Do(rqu)
	rqc, _ := http.NewRequest("PUT", "http://localhost:9090/kv/", strings.NewReader("roopa_new:Chicago"))
	c.Do(rqc)
	rqc, _ = http.NewRequest("PUT", "http://localhost:9090/kv/", strings.NewReader("test:Chicago"))
	c.Do(rqc)
	rqc, _ = http.NewRequest("PUT", "http://localhost:9090/kv/", strings.NewReader("test1:Chicago"))
	c.Do(rqc)
	//Make Delete Request
	rd, _ := http.NewRequest("DELETE", "http://localhost:9090/kv/roopa", nil)
	c.Do(rd)
	//Make Get Request
	rc, _ := http.Get("http://localhost:9090/kv/count/roopa")
	buc, _ := ioutil.ReadAll(rc.Body)
	//Verify
	if string(buc) != "1" {
		t.Error("Count Not Successful After Delete", string(buc))
	}

	service.Stop()

}
