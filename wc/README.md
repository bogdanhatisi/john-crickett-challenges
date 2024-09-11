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

- [x] Accept the `-l` option from the command line.
- [x] Open and read the file specified by the user.
- [x] Count the total number of **lines** in the file.
- [x] Output the line count followed by the filename.

Example:

```bash
ccwc -l test.txt
```

Expected Output:
```
7145 test.txt
```

### Step 3: Counting the Number of Words

In this step, your goal is to extend **`ccwc`** to support the command-line option **`-w`** that outputs the number of **words** in a file.

#### Checklist:

- [x] Accept the `-w` option from the command line.
- [x] Open and read the file specified by the user.
- [x] Count the total number of **words** in the file.
- [x] Output the word count followed by the filename.

Example:

```bash
ccwc -w test.txt
```

Expected Output:
```
58164 test.txt
```

### Step 4: Counting the Number of Characters

In this step, your goal is to extend **`ccwc`** to support the command-line option **`-m`** that outputs the number of **characters** in a file. If the current locale does not support multibyte characters, this will match the **`-c`** option.

You can learn more about programming for locales [here](https://learn.microsoft.com/en-us/globalization/locale/locale).

#### Checklist:

- [x] Accept the `-m` option from the command line.
- [x] Open and read the file specified by the user.
- [x] Count the total number of **characters** (not bytes) in the file.
- [ ] Handle multibyte characters correctly, depending on the locale.
- [x] Output the character count followed by the filename.

Example:

```bash
ccwc -m test.txt
```

Expected Output:
```
339292 test.txt
```

## Step Five: Supporting the Default Option

In this step, your goal is to support the default option — i.e., when no options are provided, the program should behave as if the `-c`, `-l`, and `-w` options were specified. This means that the program will output the number of **lines**, **words**, and **bytes** for the given file.

### Checklist:

- [x] If no command-line options are provided, the program should act as if the `-c`, `-l`, and `-w` options were passed.
- [x] Ensure that the program prints the **line count**, **word count**, and **byte count** followed by the filename.
- [x] Test that the output format matches the following:

Example:

```bash
ccwc test.txt
```

Expected Output:
```
7145   58164  342190 test.txt
```

## The Final Step: Supporting Standard Input

In this final step, your goal is to support reading from standard input if no filename is specified. When the user pipes input into the program via standard input (e.g., using `cat`), the program should process that input and output the result based on the options provided.

### Checklist:

- [x] Support reading from **standard input** when no filename is specified.
- [x] If no filename is provided, the program should read from standard input and behave as if it was provided a file.
- [] Ensure that the program correctly handles input from standard input when options such as `-l`, `-w`, `-c`, or `-m` are provided.
- [x] Ensure that the output is properly formatted, just as it would be when reading from a file.
  
### Example:

```bash
cat test.txt | ccwc -l
```

Expected Output:
```
7145
```