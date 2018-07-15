package main 

import (
	"fmt"
	"net"
	"log"
	"bufio"
	"strings"
)

const (
	PORT = ":8080" 										/* assigned port number 	*/
	CONN_TYPE = "tcp"									/* assigned conn protocol 	*/
)

var	pi = CreatePackageIndexer() 						/* instan a package indexer */ 

// start running the server with tcp connection 
// for multiple client connections 
// returns and logs any errors 
func StartServer() error {
	ln, err := net.Listen(CONN_TYPE, PORT)				/* set up server listening 	*/ 
	if err != nil {
		log.Printf("Error listening: ", err)
		return err
	}
	fmt.Println("Server is running...")
	fmt.Println("Server listening on port...", PORT)	/* successfully listening 	*/
	for {												/* accept incoming requests */ 
		conn, err := ln.Accept()  
		if err != nil {
			log.Printf("Error accepting request: ", err)
			return err 
		}
		// go routine 
		log.Println("[+] Handling new connection...")
		go handleConnection(conn)						/* handle client connection */
	}
	ln.Close()
	return nil
}

// handles client connections 
// reads input from buffer and logs request 
func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)						/* Set up buffer reader 	*/ 	
		request, err := reader.ReadString('\n') 			/* Read client request 		*/ 
		if err != nil {
			log.Println("Error reading request: ", err)
			return 
		}
		req := ParseRequest(request)						/* Send string  to parser 	*/
		response := pi.HandleRequest(req)					/* Send req to indexer 		*/
		// log request and response here  
		log.Printf("[Request: %s] [Response: %s]", 
			strings.TrimSuffix(request, "\n"), strings.TrimSuffix(response, "\n"))
		writer := bufio.NewWriter(conn)
		_, err1 := writer.WriteString(response)
		if err1 != nil {
			fmt.Errorf("Issue writing back to client %s", err1)
		}
		writer.Flush()
	}
}

