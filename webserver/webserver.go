package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// Listen for incoming connections on port 80
	ln, err := net.Listen("tcp", ":80")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Accept incoming connections and handle them
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Handle the connection in a new goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// Close the connection when we're done
	defer conn.Close()

	// Read incoming data
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Convert the buffer into a string and split it by lines
	request := string(buf[:n])
	fmt.Printf("Received Request: \n%s", request)

	// Split the request by lines
	lines := strings.Split(request, "\r\n")

	// The first line is the request line: "GET /path HTTP/1.1"
	if len(lines) > 0 {
		// Split the request line by spaces
		requestLineParts := strings.Fields(lines[0])

		// Check that the request line has at least 3 parts (e.g., GET, /path, HTTP/1.1)
		if len(requestLineParts) >= 3 {
			method := requestLineParts[0]
			path := requestLineParts[1]

			fmt.Printf("Method: %s, Path: %s\n", method, path)

			// Now you can construct the response body using the path
			body := fmt.Sprintf("You requested the path: %s", path)
			contentLength := len(body)

			// Prepare a valid HTTP/1.1 response
			response := fmt.Sprintf("HTTP/1.1 200 OK\r\n"+
				"Content-Type: text/plain\r\n"+
				"Content-Length: %d\r\n"+
				"Connection: close\r\n"+
				"\r\n"+
				"%s", contentLength, body)

			// Write the response back to the client
			conn.Write([]byte(response))
		}
	}
}
