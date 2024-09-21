<!-- TOC --><a name="-solving-john-cricketts-coding-challenges-in-go"></a>
# 🛠️ Solving John Crickett's Coding Challenges in Go

This repository contains my solutions to John Crickett's challenges. Each challenge focuses on creating a tool that performs a specific task, implemented in Go.

## Table of Contents

<!-- TOC start (generated with https://github.com/derlin/bitdowntoc) -->

- [📜 Solved Challenges](#-solved-challenges)
   * [1. Cut Challenge](#1-cut-challenge)
   * [2. Wc Challenge](#2-wc-challenge)
- [🚧 In Progress Challenges](#-in-progress-challenges)
   * [1. Basic Web Server Challenge](#1-basic-web-server-challenge)
- [⚙️ How to Run](#-how-to-run)
- [💬 Feedback](#-feedback)
- [📚 References](#-references)
- [🙏 Thank You](#-thank-you)

<!-- TOC end -->

<!-- TOC --><a name="-solved-challenges"></a>
## 📜 Solved Challenges

<!-- TOC --><a name="1-cut-challenge"></a>
### 1. [Cut Challenge](./cut)
   - A Go implementation of the `cut` Unix command.
   - Supports selecting portions of each line from a file based on field numbers and delimiters.
   - **Features**:
     - Field selection (`-f` flag)
     - Custom delimiter support (`-d` flag)
     - Handles both file input and standard input (`stdin`)

<!-- TOC --><a name="2-wc-challenge"></a>
### 2. [Wc Challenge](./wc)
   - A Go implementation of the `wc` Unix command.
   - Supports counting lines, words, bytes, and characters.
   - **Features**:
     - Line count (`-l` flag)
     - Word count (`-w` flag)
     - Byte count (`-c` flag)
     - Character count (`-m` flag)
     - Handles both file input and standard input (`stdin`)
     - Default behavior when no flags are provided (counts lines, words, and bytes)

<!-- TOC --><a name="-in-progress-challenges"></a>
## 🚧 In Progress Challenges

<!-- TOC --><a name="1-web-server-challenge"></a>
### 1. [Basic Web Server Challenge](./webserver)
   - Developing a simple web server that supports a subset of HTTP/1.1.

### 🔁 Adding tests to all challenges
   - **Goal**: Gradually add a comprehensive set of tests to ensure each challenge is fully covered.
   - **Why**: Writing tests helps make the code more reliable and easier to maintain. With good test coverage, I can refactor or add new features without worrying about breaking existing functionality.
   - **Status**: Ongoing, with tests being added progressively.
  

<!-- TOC --><a name="-how-to-run"></a>
## ⚙️ How to Run

Each challenge has its own folder and README with detailed instructions on how to run the solution.

<!-- TOC --><a name="-feedback"></a>
## 💬 Feedback

I’d love to hear your thoughts, suggestions, or improvements on these solutions! Feel free to reach out through the following channels:

- 📧 Email: [hatisi.bogdan@gmail.com](mailto:hatisi.bogdan@gmail.com)
- 💼 LinkedIn: [Bogdan Hatisi](https://linkedin.com/in/bogdan-hatisi)

Your feedback is greatly appreciated!


<!-- TOC --><a name="-references"></a>
## 📚 References

- [John Crickett's Challenges](https://codingchallenges.fyi/challenges/intro)  
   The source of these coding challenges, aimed at improving software engineering skills.

- [Go Documentation](https://golang.org/doc/)  
   Official documentation for the Go programming language, covering everything from basic syntax to advanced topics.


<!-- TOC --><a name="-thank-you"></a>
## 🙏 Thank You

A special thank you to **John Crickett** for providing these fantastic coding challenges! These exercises are a great way to hone your software engineering skills.

