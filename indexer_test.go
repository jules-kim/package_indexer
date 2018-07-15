package main 

import (
	"testing"
)

/* Test Indexing packages */

// index a package with no dependencies in an empty indexer 
func TestIndex_No_Dep_Empty_OK(t *testing.T) {
	pi := CreatePackageIndexer()
	p := Package{name: "cloog"}
	if pi.Index(&p) != "OK" {
		t.Errorf("Expected indexing to succeed") 
	}
}
// index a package with no dependencies in a nonempty indexer
func TestIndex_NO_Dep_NONEMPTY_OK(t *testing.T) {
	pi := CreatePackageIndexer()
	p := &Package{name: "cloog"}
	pi.packs[p.name] = p
	if pi.Index(p) != "OK" {
		t.Errorf("Expected indexing to succeed") 
	}
}

// index a package that has its one dependency installed 
func TestIndex_One_Dep_OK(t *testing.T) {
	pi := CreatePackageIndexer()
	d := &Package{name: "gmp"}
	p := Package{name:"cloog", deps: []*Package{d}}
	pi.packs[d.name] = d
	if pi.Index(&p) != "OK" {
		t.Errorf("Expected indexing to succeed") 
	}
}

// index a package that does not have dependencies installed 
func TestIndex_Dep_FAIL(t *testing.T) {
	pi := CreatePackageIndexer()
	d := &Package{name: "gmp"}
	p := Package{name:"cloog", deps: []*Package{d}}
	if pi.Index(&p) != "FAIL" {
		t.Errorf("Expected indexing to fail") 
	}
}	



/* Test Removing packages */


/* Test Querying packages */ 

// test indexer querying for a package in an index
// that contains only one package
func TestQuery_OK(t *testing.T) {
	pi := CreatePackageIndexer()
	p := &Package{name: "cloog"}
	pi.packs[p.name] = p
	if pi.Query(p.name) != "OK" {
		t.Errorf("Expected querying to succeed") 
	}
}

// query for package that exists in an index
// with multiple packages 
func TestQuery_Multi_Packs_OK(t *testing.T) {
	pi := CreatePackageIndexer()
	cloog := &Package{name: "cloog"}
	ceylon := &Package{name: "cloog"}
	p := []*Package{cloog, ceylon}
	for _, pack := range p {
		pi.packs[pack.name] = pack
	}
	if pi.Query("cloog") != "OK" {
		t.Errorf("Expected querying to succeed") 
	}
}

// query for a package in an empty index 
func TestQuery_Empty_FAIL(t *testing.T) {
	pi := CreatePackageIndexer()
	p := &Package{name: "cloog"}
	if pi.Query(p.name) != "FAIL" {
		t.Errorf("Expected querying to fail") 
	}
}

// query for a package that is not indexed
func TestQuery_Nonempty_FAIL(t *testing.T) {
	pi := CreatePackageIndexer()
	p := &Package{name: "cloog"}
	pi.packs[p.name] = p
	if pi.Query("pkg-config") != "FAIL" {
		t.Errorf("Expected querying to fail") 
	}
}

/* Test for errors */ 
