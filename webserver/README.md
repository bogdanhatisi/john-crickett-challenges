# Build Your Own Basic Web Server

This challenge is to build your own basic web server.

At its core, a web server is actually quite simple. It’s a server that listens for connections from clients and responds to them. The clients make those requests using a protocol known as HTTP (and expect responses in the same protocol, obviously).

HTTP, like many early Unix and Internet protocols, is text-based and human-readable. The original HTTP specification from 1991 is short and sweet. It originally didn’t even have a version number! But it was later renamed HTTP/0.9 to differentiate it from HTTP/1.0. The HTTP/1.0 specification seems to be lost in the mists of time.

The first full formal HTTP specification is HTTP/1.1—also known as RFC2616 from 1999. You can dig through all the HTTP standards on the [W3C website](https://www.w3.org/Protocols/).

---
    
## The Challenge - Building a Basic Web Server

Early web servers were very basic. As per the HTTP/0.9 specification, they supported just a `GET` request and returned the specified document. The error messages were human-readable HTML with no way to distinguish success from failure.

We’re going to go a little beyond that and support a small subset of HTTP/1.1.

## Steps

### Step 0: Setting Up

In this step, you decide which programming language and IDE you’re going to use and set yourself up with a nice new `webserver` project. I built mine in Rust.

#### Checklist:

- [x] Choose a programming language (e.g., Rust, Go, Python, Node.js).
- [x] Set up your development environment and IDE/editor.
- [x] Initialize a new project for your web server.
- [x] Familiarize yourself with the basics of handling HTTP requests in your chosen language.

### Step 1: Creating a Basic HTTP Server

In this step your goal is to create a basic HTTP server that listens on port 80 and can handle a single TCP connection at a time. For all requests we’ll return some text that describes the requested path.

For example if our server is running locally, a client might make the curl request below and get back the simple text message.

curl http://localhost/
Requested path: /

To support this your server will need to create a socket, and bind it to the address of your server. For the purposes of this challenge that can be the IP address: 127.0.0.1 (the loopback address, also known as localhost) and port: 80 (the default HTTP port).

Once that is done, your server will need to listen for requests, accept incoming requests and then receive the incoming data.

You can learn more about sockets in the Wikipedia article on Berkeley Sockets. Your programming language probably provides a wrapper around this API in its standard library for example Python has socket, Rust has std::net and node has node:net.

For an in-depth look at network programming check out Beej’s Guide to Network Programming.

Once you can receive data from the client you’ll need to parse that data to extract the key elements. For this step that is simply taking the first line of the client request, which will look something like this:

GET / HTTP/1.1

From which you’ll need to recognise that this is the Request value, the type of request is GET, the requested resource is / and the HTTP version is 1.1.

For this step, you’ll need to return the bare minimum HTTP response:

HTTP/1.1 200 OK\r\n\r\nRequested path: <the path>\r\n

Which you will do by sending that data back over the socket. Once you have that working we have a very basic server, but it’s not much use yet. Don’t forget to close the socket when you’re done sending data.

#### Checklist:

- [x] Create a socket and bind it to 127.0.0.1:80
- [x] Listen for incoming requests
- [x] Accept a single TCP connection
- [x] Receive data from the client
- [x] Parse the request line to extract the method, path, and HTTP version
- [x] Send back the minimum HTTP response
- [x] Close the socket after sending the response

### Step 2: Serving Your First HTML Document

In this step your goal is to serve your first HTML document. When the request is for a valid path, your web server should return the document at that path and the HTTP status code 200 as in step 1.

### Security Considerations
Pause for a moment and think about the security risks you might have introduced in your simple server.

(SPOILER)
#### Security Risks in Web Server Implementation

1. **Directory Traversal Attacks**: 
   - Attackers can access files outside the intended directory.

2. **Improper Input Validation**: 
   - Vulnerable to injection attacks if user input is not validated.

3. **Denial of Service (DoS)**: 
   - Server can be overwhelmed by too many requests.

4. **Lack of HTTPS**: 
   - Data transmitted is unencrypted, susceptible to eavesdropping.

5. **Error Handling**: 
   - Detailed error messages can reveal server structure and vulnerabilities.

6. **Resource Exhaustion**: 
   - Large payloads can consume server resources.

7. **Insecure File Permissions**: 
   - Overly permissive permissions can allow unauthorized access.

8. **Lack of Authentication and Authorization**: 
   - Anyone can access any resource without restrictions.

9. **Cross-Site Scripting (XSS)**: 
   - Dynamic content without proper sanitization can lead to XSS attacks.


#### Checklist:
- [x] Create an HTML file named `index.html` in the `www` directory.
- [x] Modify the server to serve the `index.html` file for requests to `/` and `/index.html`.
- [x] Ensure the server returns a 404 status code for invalid paths.
- [x] Test the server with valid and invalid paths to confirm correct behavior.

### Step 3: Handling Multiple Clients Concurrently

In this step, your goal is to handle multiple clients at a time. After all, in the real world, we don’t want a web server that can only handle one client at a time—that’s just not going to scale.

We want to be able to handle multiple concurrent clients, binding each incoming connection to a new socket and then receiving and sending data over that socket connection. We don’t want one connection to block another, so we’ll need some support for concurrent operations in our server.

Traditionally, this was handled by creating a new thread for each connection. More recently, there has been a focus on using asynchronous frameworks. This boils down to there being an overhead to threads. If you’ve never used threads, now would be a great time to try it out. If you have, then perhaps try the async framework for your stack of choice. There’s an overview of parallelism, concurrency, and asynchrony on the Coding Challenges blog.

You can then test your server by sending it multiple concurrent requests. If you want to see what’s happening, you could make your connection handling thread log its ID and then sleep for 20 seconds before sending the HTML back. This way, you can observe your concurrency in action.

By now, you have a basic web server that can handle multiple concurrent client requests.

#### Checklist:
- [x] Modify the server to handle multiple concurrent connections.
- [x] Implement threading or asynchronous handling for each connection (Go Routines).
- [x] Log the thread ID or connection ID for each request to observe concurrency (Not working in Go).
- [x] Test the server with multiple concurrent requests using curl.


### Step 4: Securing File Access

In this step, your goal is to address some of the issues you hopefully thought of at the end of Step 2. Currently, our server will open a valid file and stream it back to the client. If we’re not careful, clients could provide a path to a file we don’t want to share, such as `/etc/passwd` on a Unix-like system.

Your task for this part of the challenge is to ensure that whatever request path is provided by the client, you only serve documents that are in the `www` directory or its subdirectories. 

Ideally, your server should allow you to specify the location of the `www` folder on startup. Devise some test cases and give it a go!

Once you’ve done that, congratulations 🎉 you’ve built a basic web server that can serve HTML documents!

#### Checklist:
- [ ] Implement path validation to ensure only files within the `www` directory or its subdirectories are served.
- [ ] Allow the server to specify the location of the `www` folder on startup.
- [ ] Create test cases to verify that the server correctly restricts access to files outside the `www` directory.
- [ ] Test the server with valid paths (within `www`) and invalid paths (outside `www`) to confirm correct behavior.