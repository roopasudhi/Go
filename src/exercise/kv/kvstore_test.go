package kv

import (
	"testing"
)

func TestPut(t *testing.T) {
	s := NewStore()
	s.Put("test", "chicago")
	if s.kv["test"] != "chicago" {
		t.Error("Add Not Successful")
	}
}

func TestUpdate(t *testing.T) {
	s := NewStore()
	s.Put("test", "chicago")
	s.Put("test", "Illinois")
	if s.kv["test"] != "Illinois" {
		t.Error("Add Not Successful")
	}
}

func TestDelete(t *testing.T) {
	s := NewStore()
	s.Put("test", "chicago")
	s.Put("test1", "Illinois")
	s.Delete("test1")
	if len(s.kv) != 1 {
		t.Error("Delete Not Successful")
	}
	s.Delete("test3")
}

func TestCountWithPrefix(t *testing.T) {
	s := NewStore()
	s.Put("test", "chicago")
	s.Put("test1", "Illinois")

	if len(s.kv) != s.Count("test") {
		t.Error("Count Not Successful")
	}

}
func TestCount(t *testing.T) {
	s := NewStore()
	s.Put("test", "chicago")
	s.Put("test1", "Illinois")
	s.Put("count", "Illinois")
	if len(s.kv) != s.Count("") {
		t.Error("Count Not Successful")
	}
}
