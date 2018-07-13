// file to parse client requests 
package main 

import (
	"fmt"
	"bytes"
)

// parseRequest takes the client request in the form of a
// byte array and splits the request into its three parts
// <command> <package> <dependencies> 
// It returns these three in the Request struct 
func ParseRequest(req []byte) Request {
	splitReq := bytes.Split(req, []byte("|"))		/* split the byte[] by | */
	if len(splitReq) != 3 {
		log.Println("Error: incorrect request format")
		return 
	}

	// no dependencies included, then dep is ""


	return Request("", "", "")
}

// parse dependencies by , 
func ParseDep(dep []byte) []string {

}