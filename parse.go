package main 

import (
	"bytes"
)

// TODO: Need to trim for spaces 

// file to parse client requests into command, 
// package, and dependencies

// parseRequest takes the client request in the form of a
// byte array and splits the request into its three parts
// <command> <package> <dependencies> 
// It returns these three in a Request struct 
func ParseRequest(req []byte) Request {
	splitReq := bytes.Split(req, []byte("|"))					/* split the byte[] by "|"		*/
	if len(splitReq) != 3 {
		return Request{"", "", nil, "Incorrect request format"}
	}
	deps := parseDep(splitReq[2])								/* parse deps by "," 			*/
	com := string(splitReq[0][:])								/* convert to a string 			*/ 
	pack := string(splitReq[1][:])
	
	return Request{com, pack, deps, ""}  
}

// parse dependencies by "," and returns 
// the dependencies converted to an array of strings
func parseDep(d []byte) []string {
	deps_bytes := bytes.Split(d, []byte(","))					/* split deps[] by ","			*/ 
	deps := make([]string, len(deps_bytes))
	for i, dep := range deps_bytes {							/* convert each ele to a str 	*/ 
		deps[i] = string(dep[:])
	}

	return deps
}