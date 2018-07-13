package main 

import (
	"fmt"
	"net"
	"log"
	"bufio"
)

const (
	PORT = ":8080" 			/* assigned port number 		*/
	CONN_TYPE = "tcp"		/* assigned connection protocol */
)

// start running the server with tcp connection 
// for multiple client connections 
// returns and logs any errors 
func StartServer() error {
	ln, err := net.Listen(CONN_TYPE, PORT)				/* set up server listening 	*/ 
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	
	fmt.Println("Server is running...")
	fmt.Println("Server listening on port...", PORT)	/* successfully listening 	*/
	
	for {												/* accept incoming requests */ 
		conn, err := ln.Accept()  
		if err != nil {
			log.Printf("%s", err)
			return err 
		}
		log.Println("Request from client made")			/* log new client request 	*/ 
		// go routine 
		go handleConnection(conn)						/* handle client connection */
	}
}

// handles client connections 
// reads input from buffer
func handleConnection(conn net.Conn) {
	fmt.Println("Handling new connection...")
	reader := bufio.NewReader(conn)						/* Set up buffer reader 	*/ 	
	request, err := reader.ReadString('\n') 			/* Read client request 		*/ 
	if err != nil {
		log.Println("%s", err)
		return 
	}
	log.Println(request)
}


