# ascii-art

# Understanding the Core Idea
This program does 3 main things:
```
1. Read the banner file (standard.txt or shadow.txt or thinkertoy.txt)

2. Convert the input string characters to ASCII art lines

3. Print the result
```

Each character in the banner file has:

```
height = 8 lines
```

So, printing `"HELLO"` means:

```
print line1 of H + line1 of E + line1 of L + line1 of L + line1 of O
print line2 of H + line2 of E + Line2 of L + line2 of L + line2 of O
...
Repeat until line8
```

This is the **core algorithm**

## Folder Structure
```console
ascii-art/
├── main.go
├── standard.txt
├── shadow.txt
├── thinkertoy.txt
└── ascii/
    └── render.go
```

**Explanation:**
* main.go: handles input (os.Args)
* banner text files: has the ASCII templates
* ascii/: contains the logic for rendering the ASCII art
* render.go: contains the functions for processing

## How the banner files work
The banner file contains ASCII Characters from:
```
ASCII 32 -> ASCII 126
```

Total characters:

```
95 characters. 
```

Confirm this from the man ascii table [here](https://www.man7.org/linux/man-pages/man7/ascii.7.html).

So, the file structure is:
```(blank line)

8 lines -> space
8 lines -> !
8 lines -> "
8 lines -> #
...
```

To find a character position:
```
index = (Ascii_code - 32) * 9
```

Why 9?
```
8 lines
+1 seperator line
```
---

## Undertsanding this program in simple terms
- The program reads the input string from `os.Args`.
- It then loads the ASCII banner file.
- Each character in the banner file is represented by 8 lines.
- The position of each charater is calculated using its ASCII. code minus 32.
- Then it prints the ASCII representation line by line for the entire string.

## Simple Visual of Program Flow
```console
           +-------------+
Input ---> |  main.go    |
           +-------------+
                   |
                   v
           +-------------+
           | ascii.Render|
           +-------------+
                   |
        +----------------------+
        | read banner file     |
        +----------------------+
                   |
        +----------------------+
        | build ascii mapping  |
        +----------------------+
                   |
        +----------------------+
        | render each line     |
        +----------------------+
                   |
                   v
              OUTPUT
```
