package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
)

// getContentType returns the MIME type based on the file extension
func getContentType(filePath string) string {
	ext := strings.ToLower(filepath.Ext(filePath))
	switch ext {
	case ".html":
		return "text/html"
	case ".txt":
		return "text/plain"
	case ".css":
		return "text/css"
	case ".js":
		return "application/javascript"
	case ".json":
		return "application/json"
	case ".png":
		return "image/png"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".gif":
		return "image/gif"
	default:
		return "application/octet-stream" // Default for unknown types
	}
}

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
			path := requestLineParts[1]

				// Construct the file path
				if path == "/"{
					path = "/index.html"
				}
				filePath := "www" + path
				// Read the file from the www directory
				body, err := os.ReadFile(filePath)
				if err != nil {
					// Handle file not found or read error
					response := "HTTP/1.1 404 Not Found\r\n\r\nFile not found."
					conn.Write([]byte(response))
					return
				}

				// Determine the content type
				contentType := getContentType(filePath)

				// Now you can construct the response body using the file contents
				contentLength := len(body)
				response := fmt.Sprintf("HTTP/1.1 200 OK\r\n"+
					"Content-Type: %s\r\n"+
					"Content-Length: %d\r\n"+
					"Connection: close\r\n"+
					"\r\n"+
					"%s", contentType, contentLength, body)

				conn.Write([]byte(response))
			}
		}
	}
