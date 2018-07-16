package main 

import (
	"fmt"
	"net"
	"log"
	"bufio"
	"strings"
	"os"
	"os/signal"
	"syscall"
	"io"
)

const (
	// assigned port number 
	PORT = ":8080" 					
	// assigned conn protocol 				
	CONN_TYPE = "tcp"									
)

// instantiate a package indexer
var	pi = CreatePackageIndexer() 						

// start running the server with tcp connection 
// for multiple client connections 
// returns and logs any errors 
func StartServer() {
	// server listening for connections 
	ln, err := net.Listen(CONN_TYPE, PORT)				 
	if err != nil {
		log.Fatal("Error listening: ", err)
	}
	fmt.Println("Server is running...")
	fmt.Printf("Server listening on port %s\n", PORT)
	// server successfully listening for incoming connections
	for {				
		// catch any terminating signals and terminate gracefully
		c := make(chan os.Signal)
	    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	    go func() {
	        <-c
	        fmt.Println(" Shutting down server...")
	        os.Exit(1)
	    }()
		// accept any incoming requests 								
		conn, err := ln.Accept()  
		if err != nil {
			log.Printf("Error accepting request: %s", err)
		}
		// go routine 
		log.Println("[+] Handling new connection...")
		// handle client connection 
		go handleConnection(conn)						
	}
	ln.Close()
}

// handles client connections 
// reads input from buffer and logs request 
func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)		
		// Read client request until newline character			 	
		request, err := reader.ReadString('\n') 			 
		if err != nil {
			if err == io.EOF {
				log.Println("[-] Client closed connection")
			} else {
				log.Printf("Error reading request: %s", err)
			}
			return 
		}
		// send request string to parse 
		req := ParseRequest(request)	
		// Send parsed request to indexer 				
		response := pi.HandleRequest(req)					
		// log request and response  
		log.Printf("[Request: %s] [Response: %s]", 
			strings.TrimSuffix(request, "\n"), strings.TrimSuffix(response, "\n"))
		// write response back to the client 
		writer := bufio.NewWriter(conn)
		_, err1 := writer.WriteString(response)
		if err1 != nil {
			fmt.Errorf("Issue writing back to client: %s", err1)
		}
		writer.Flush()
	}
}

