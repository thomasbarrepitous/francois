package main

import (
	"bufio"
	"fmt"
	"francois/lexer"
	"log"
	"os"
)

func main() {
	fmt.Print("Enter string : ")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	l := lexer.Tokenize(line)
	fmt.Println(l)
}
