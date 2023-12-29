package main

import (
	"bufio"
	"fmt"
	"francois/parser"
	"francois/runtime"
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

func initEnv() *runtime.Environment {
	env := runtime.NewEnvironment(nil)
	env.DeclareVariable("a", runtime.NumericValue{Value: 100})
	env.DeclareVariable("b", runtime.NumericValue{Value: 200})
	env.DeclareVariable("true", runtime.BooleanValue{Value: true})
	env.DeclareVariable("false", runtime.BooleanValue{Value: false})
	return env
}

func execFile(path string) {
	env := initEnv()
	lines, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	exec(string(lines), env)
}

func execRepl() {
	env := initEnv()
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
		exec(line, env)
	}
}

func exec(sourceCode string, env *runtime.Environment) {
	parser := parser.Parser{}
	program := parser.ProduceAST(sourceCode)
	runtime := program.Evaluate(env)
	fmt.Printf("%+v\n", runtime)
}
