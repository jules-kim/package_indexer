package main 

import (
	"fmt"
	"net"
	"log"
)

const (
	PORT = ":8080" 			/* assigned port number 		*/
	CONN_TYPE = "tcp"		/* assigned connection protocol */
)

// start running the server with tcp connection 
func StartServer() error {
	ln, err := net.Listen(CONN_TYPE, PORT)					/* set up server listening 	*/ 
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	
	fmt.Println("Server is running...")
	fmt.Println("Server listening on port...", PORT)		/* successfully listening 	*/
	
	for {													/* accept incoming requests */ 
		_, err := ln.Accept() // conn 
		if err != nil {
			log.Printf("%s", err)
			return err 
		}
		fmt.Println("Request from client made.")
		//go handleConnection(conn)							/* handle client requests 	*/
	}
}