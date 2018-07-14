package main 

import "sync"

// This type is used to encapsulate the 
// three aspects of the client request,
// the command, package, and dependencies 
type Request struct {
	comm 	string
	pack 	string 
	dep 	[]string
	err 	string 
}

// This type is used for package storage
// within a map on the server 
type PackageIndexer struct {
	packs 	map[string]*Package 
	mutex 	*sync.Mutex
}

// This type is used to represent packages
// and their dependencies 
type Package struct {
	Name	string 
	deps 	[]*Package
}