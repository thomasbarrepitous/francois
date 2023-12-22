package main

import (
	"bufio"
	"fmt"
	"francois/parser"
	"log"
	"os"
)

const (
	ExecType string = "REPL"
)

func main() {
	if ExecType == "FILE" {
		execFile("test.txt")
	}
	if ExecType == "REPL" {
		execRepl()
	}
}

func execFile(path string) {
	lines, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	parser := parser.Parser{}
	program := parser.ProduceAST(string(lines))
	for _, statement := range program.Body {
		fmt.Printf("%+v\n", statement)
	}
}

func execRepl() {
	for {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if line == "exit\n" {
			break
		}
		parser := parser.Parser{}
		program := parser.ProduceAST(line)
		for _, statement := range program.Body {
			fmt.Printf("%+v\n", statement)
		}
	}
}
