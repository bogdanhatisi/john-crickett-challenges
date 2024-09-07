# Build Your Own `cut` Tool Challenge

This challenge is to build your own version of the Unix command line tool `cut`! The Unix command line tools are excellent examples of good software engineering and they follow the Unix Philosophy:

- Writing simple parts connected by clean interfaces – each tool does one thing and provides a simple CLI that handles text input from files or streams.
- Design programs to be connected to other programs – each tool can be easily connected to other tools via files and streams, allowing powerful compositions.

Following these philosophies has made Unix command line tools some of the most widely used engineering tools, and they can be chained together to create powerful workflows.

You can read more about the Unix Philosophy in the excellent book **The Art of Unix Programming**.

---

## The Challenge: Building `cut`

The functional requirements for `cut` are described concisely in its man page. To see this on your system, try:

```bash
 man cut
 ```

### TL;DR:
`cut` - cuts out selected portions from each line in a file.

---

## Steps

### Step 0: Setup your environment
- [x] Set up your IDE/editor of choice and programming language.
- [x] Download the provided text file from Dropbox.

---

### Step 1: Simple version of `cut`
- [x] Implement a simple version of `cut` that opens the provided tab-separated file and prints the second field (`-f2`) from each line.

Example:
```bash
cut -f2 sample.tsv
```

Expected output:
f1
1
6
11
16
21

### Step 2: Support the `-d` option for custom delimiters
- [ ] Extend your code to support the `-d` option to allow the user to specify what character to use as the delimiter between fields.
- [ ] If no delimiter is provided, default to using a tab (`\t`).

Example with comma as the delimiter:
```bash
cut -f1 -d, fourchords.csv | head -n5
```

Expected output:
Song title
"10000 Reasons (Bless the Lord)"
"20 Good Reasons"
"Adore You"
"Africa"

Example with default tab delimiter:
```bash
cut -f1 sample.tsv
```

Expected output:
f0
0
5
10
15
20