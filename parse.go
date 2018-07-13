// file to parse client requests 
package main 

import (
	"fmt"
	"strings"
)

// parseRequest takes the client request in the form of a
// byte array and splits the request into its three parts
// <command> <package> <dependencies> 
// It returns these three in the Request struct 
func ParseRequest(req []byte) Request {
	

}