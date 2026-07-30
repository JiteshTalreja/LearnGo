/*

Go Strings vs Runes (Mental Model)
1. A Go string is always stored as UTF-8 encoded bytes.
A string is a sequence of bytes, not a sequence of runes.
Even "A" is stored as bytes.

Example:

"A"

Memory:

65
2. A rune represents a single Unicode code point.
A rune is not a string.
rune is an alias for int32.

Example:

'A'

is stored as:

int32(65)
3. Double quotes (" ") create strings, single quotes (' ') create rune literals.
"A"   // string
'A'   // rune (int32)

Although both represent the same character, they are completely different types.

4. String indexing returns a byte, not a rune.
text := "Hello"

text[0]

returns

byte  // uint8

because indexing means:

"Give me the byte stored at this position."

It does not decode UTF-8.

5. range over a string decodes UTF-8 into runes.
for _, ch := range text {
}

Here:

ch

is a

rune

because Go decodes the UTF-8 byte sequence into Unicode code points.

6. Strings store bytes; runes store Unicode values.

Example:

"H"

Stored as UTF-8 bytes:

65

Whereas

'H'

Stored as:

int32(65)

No UTF-8 encoding is involved for a rune.

7. Why does Go have runes?

Go separates text from individual characters.

A string represents text.

A rune represents one Unicode character.

This makes parsing, lexing, tokenizing, and character processing simple, type-safe, and efficient.

8. Why not use one-character strings like Python?

Python treats a single character as a one-character string.

Go intentionally distinguishes between:

Text → string
Single Unicode character → rune

This gives better type safety and makes character-oriented code (parsers, compilers, tokenizers, validators) more explicit.

9. Character comparisons use runes.

Instead of:

if ch == "a"

Go uses:

if ch == 'a'

which compares two Unicode code points (int32 values).

10. The most important mental model
            "Hello"
                │
                ▼
        UTF-8 encoded bytes
                │
      +----+----+----+----+----+
      |72  |101 |108 |108 |111 |
      +----+----+----+----+----+
                │
      text[i] returns one byte
                │
                ▼
          byte (uint8)

--------------------------------------

        range over string
                │
                ▼
      Decode UTF-8 bytes
                │
                ▼
       Unicode Code Point
                │
                ▼
          rune (int32)
⭐ The one sentence I want you to remember

A Go string stores UTF-8 bytes. A rune stores one Unicode code point. Indexing returns bytes; range decodes bytes into runes.
*/

package main

import "fmt"

func main() {
	// var someText = "hello"
	// fmt.Println("%C\n", someText[0])

	var name, age = "Kim", 31

	fmt.Printf("%s is %d years old", name, age)

	fmt.Println("\n \t this is some text with emojies 😀🥸")

	fmt.Println(`this is a some raw string
                     spread out in multiple 
                     linesssssssssssss`)

	something := "56"
	fmt.Println(something)

	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5

	fmt.Printf("%b, %b, %b, %b, %b, %b", a, b, c, d, e, f)

	fmt.Printf("\n%x, %x, %x, %x, %x, %x", a, b, c, d, e, f)
}
