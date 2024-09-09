# Build Your Own `wc` Tool

This challenge is to build your own version of the Unix command-line tool `wc`! The Unix command-line tools follow the Unix Philosophies:

- **Writing simple parts connected by clean interfaces**: Each tool does one thing and provides a simple CLI that handles text input from files or file streams.
- **Design programs to be connected to other programs**: Each tool can be easily connected to other tools to create incredibly powerful compositions.

Following these philosophies has made Unix command-line tools some of the most widely used software engineering tools, allowing us to create complex text data processing pipelines from simple commands.

You can read more about the Unix Philosophy in the excellent book **The Art of Unix Programming**.

---

## The Challenge - Building `wc`

The functional requirements for `wc` are concisely described in its man page. You can check it out in your local terminal by running:

```bash
man wc
```

TL;DR:

wc – word, line, character, and byte count.

## Steps

### Step 0: Setting up

Like all good software engineering, we’re zero-indexed! In this step, you’re going to set your environment up, ready to begin developing and testing your solution.

I’ll leave you to set up your IDE/editor and programming language of choice. After that, here’s what I’d like you to do to be ready to test your solution:

-[x] Download the test text and save it as test.txt to use for testing.

### Step 1: Writing a Simple Version of `wc`

In this step, your goal is to write a simple version of `wc`, let's call it **`ccwc`** (cc for Coding Challenges), that takes the command-line option **`-c`** and outputs the number of **bytes** in a file.

#### Checklist:

- [x] Accept the `-c` option from the command line.
- [x] Open and read the file specified by the user.
- [x] Count the total number of **bytes** in the file.
- [x] Output the byte count followed by the filename.


Example:

```bash
ccwc -c test.txt
```

Expected Output:
```
342190 test.txt
```

### Step 2: Counting the Number of Lines

In this step, your goal is to extend **`ccwc`** to support the command-line option **`-l`** that outputs the number of **lines** in a file.

#### Checklist:

- [ ] Accept the `-l` option from the command line.
- [ ] Open and read the file specified by the user.
- [ ] Count the total number of **lines** in the file.
- [ ] Output the line count followed by the filename.

Example:

```bash
ccwc -l test.txt
```

Expected Output:
```
7145 test.txt
```