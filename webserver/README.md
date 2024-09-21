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
- [ ] Familiarize yourself with the basics of handling HTTP requests in your chosen language.