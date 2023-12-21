package main

import (
	"fmt"
	"francois/lexer"
)

var src string = `a = 1 + 2 * 3`

func main() {
	l := lexer.Tokenize(src)
	fmt.Println(l)
}
