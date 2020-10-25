package main

import (
	"fmt"
	"strings"
)

func main() {
	hello := "hello"

	// declare the builder
	sb := strings.Builder{}

	// create a simple build
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")
	fmt.Println(sb.String())

	// words example
	words := []string{"hello", "world"}
	sb.Reset()

	sb.WriteString("<ul>")
	for _, v := range words {
		sb.WriteString(v)
	}
	sb.WriteString("</ul>")
	fmt.Println(sb.String())
}
