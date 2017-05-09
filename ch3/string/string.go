// string
package main

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	s := "test 星期二"
	fmt.Println(len(s))                    // "14"
	fmt.Println(utf8.RuneCountInString(s)) // "8"

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	r := []rune(s)
	fmt.Printf("%x\n", r)
	fmt.Println(len(r))
	fmt.Println(string(r))

	fmt.Println(intsToString(r))

	fmt.Println(strconv.Itoa(1234))
	fmt.Println(strconv.FormatInt(1234, 2))
	fmt.Println(strconv.FormatInt(1234, 16))

	x, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(x)
}

func intsToString(values []rune) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
