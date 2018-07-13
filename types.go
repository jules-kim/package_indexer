package main 

// This type is used to encapsulate the 
// three aspects of the client request,
// the command, package, and dependencies 
type Request struct {
	comm string
	pack string 
	dep []string
}