package main 

import "strings"

// file to parse client requests into command, 
// package, and dependencies

// parseRequest takes the client request in the form of a
// string and splits the request into its three parts
// <command> <package> <dependencies> 
// It returns these three in a Request struct 
func ParseRequest(req string) Request {
	trimReq := strings.TrimSpace(strings.TrimSuffix(req, "\n"))
	// split the string by "|"
	splitReq := strings.Split(trimReq, "|")						
	if len(splitReq) != 3 {
		return Request{"", "", nil, "Incorrect request format"}
	}
	deps := parseDep(splitReq[2])								
	com := splitReq[0]
	pack := splitReq[1]
	
	return Request{com, pack, deps, ""}  
}

// parse dependencies by "," and returns 
// the dependencies as an array of strings 
func parseDep(d string) []string {
	deps := strings.Split(d, ",")							

	return deps
}