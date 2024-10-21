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
