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
```
f1
1
6
11
16
21
```

### Step 2: Support the `-d` option for custom delimiters
- [x] Extend your code to support the `-d` option to allow the user to specify what character to use as the delimiter between fields.
- [x] If no delimiter is provided, default to using a tab (`\t`).

Example with comma as the delimiter:
```bash
cut -f1 -d, fourchords.csv | head -n5
```

Expected output:
```
Song title
"10000 Reasons (Bless the Lord)"
"20 Good Reasons"
"Adore You"
"Africa"
```

Example with default tab delimiter:
```bash
cut -f1 sample.tsv
```

Expected output:
```
f0
0
5
10
15
20
```


### Step 3: Support for the `-f` Option

In this step, your goal is to add support for the `-f` option. This option specifies a **list of fields** to be printed out. The output fields are separated by a single occurrence of the field delimiter character.

The field list is a **comma-separated** or **whitespace-separated** list of fields, i.e., `-f1,2` or `-f "1 2"`. The first field is field number 1.

#### Checklist:

- [x] Add support for the `-f` option to specify a list of fields.
- [x] The field list should be **comma-separated** or **whitespace-separated**.
- [x] Ensure that the output fields are separated by a single occurrence of the **field delimiter**.
- [x] Ensure correct handling of field ranges (i.e., print specified fields for each line).

Example 1: Comma-Separated Field List

```bash
cut -f1,2 sample.tsv
```

Expected Output:
```
f0      f1
0       1
5       6
10      11
15      16
20      21
```

In this case, the fields 1 and 2 are printed for each line, separated by a tab (default delimiter).

Example 2: Whitespace-Separated Field List with Custom Delimiter

```bash
cut -d, -f"1 2" fourchords.csv | head -n5
```

Expected Output:
```
Song title,Artist
"10000 Reasons (Bless the Lord)",Matt Redman and Jonas Myrin
"20 Good Reasons",Thirsty Merc
"Adore You",Harry Styles
"Africa",Toto
```

### Step 4: Support for Standard Input Stream (`stdin`)

In this step, your goal is to support reading input from the **standard input stream** if no filename is provided or if the single dash `-` is provided as the filename.

## Checklist:
- [ ] Add support for reading from the standard input (`stdin`) when no filename is provided.
- [ ] Add support for using `-` to represent standard input.
- [ ] Ensure that data is processed correctly from the input stream when reading from `stdin`.

Example 1: Using Piped Input from `tail`

```bash
tail -n5 fourchords.csv | cut -d, -f"1 2"
```

Expected Output:
```
"Young Volcanoes",Fall Out Boy
"You Found Me",The Fray
"You'll Think Of Me",Keith Urban
"You're Not Sorry",Taylor Swift
"Zombie",The Cranberries
```

Example 2: Using - for Standard Input

```bash
tail -n5 fourchords.csv | cut -d, -f"1 2" -
```

Expected Output:
```
"Young Volcanoes",Fall Out Boy
"You Found Me",The Fray
"You'll Think Of Me",Keith Urban
"You're Not Sorry",Taylor Swift
"Zombie",The Cranberries
```
