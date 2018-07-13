package main 

import (
	"testing"
	"reflect"
)

func TestParseRequest(t *testing.T) {
	req := ParseRequest([]byte("INDEX|cloog|gmp"))
    if req.comm != "INDEX" && req.pack != "cloog" && 
    reflect.DeepEqual([1]string{"gmp"}, req.dep) {
       t.Errorf("Parsing request was incorrect") 
    }

    req2 := ParseRequest([]byte("INDEX|cloog|"))
    if req2.comm != "INDEX" && req2.pack != "cloog" && 
    reflect.DeepEqual([1]string{""}, req2.dep) {
       t.Errorf("Parsing request was incorrect") 
    }
}

func TestParseDep(t *testing.T) {

}