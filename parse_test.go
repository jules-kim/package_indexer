package main 

import (
	"testing"
	"reflect"
)

func TestParseRequest(t *testing.T) {
	// One dependency included 
	req := ParseRequest([]byte("INDEX|cloog|gmp"))
    if req.comm != "INDEX" || req.pack != "cloog" || 
    !reflect.DeepEqual([]string{"gmp"}, req.dep) {
       t.Errorf("Parsing request with one dependency was incorrect") 
    }
    // No dependencies included 
    req2 := ParseRequest([]byte("INDEX|cloog|"))
    if req2.comm != "INDEX" || req2.pack != "cloog" ||
    !reflect.DeepEqual([]string{""}, req2.dep) {
       t.Errorf("Parsing request with no dependencies was incorrect") 
    }
    // Multiple dependencies included 
    req3 := ParseRequest([]byte("INDEX|cloog|gmp,isl,pkg-config"))
    arr := []string{"gmp","isl","pkg-config"}
    if req3.comm != "INDEX" || req3.pack != "cloog" || 
    !reflect.DeepEqual(arr, req3.dep) {
       t.Errorf("Parsing request with multiple dependencies was incorrect") 
    }
    // Testing badly formatted requests 
    // Testing too many | characters 
    req4 := ParseRequest([]byte("INDEX|cloog|gmp|"))
    if req4.err == "" {
    	t.Errorf("Incorrect catching of formatting error")
    }
    // Testing incorrect message pattern 
    req5 := ParseRequest([]byte("INDEX|"))
    if req5.err == "" {
    	t.Errorf("Incorrect catching of formatting error")
    }
}

func TestParseDep(t *testing.T) {
	deps := parseDep([]byte("gmp,isl,pkg-config"))
	arr := []string{"gmp","isl","pkg-config"}
	if !reflect.DeepEqual(arr, deps) {
		t.Errorf("Parsing of dependencies failed") 
	}
}