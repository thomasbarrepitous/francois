package runtime

import (
	"log"
)

type Environment struct {
	Parent *Environment
	Values map[string]RuntimeValue
}

func NewEnvironment(parent *Environment) *Environment {
	return &Environment{
		Parent: parent,
		Values: make(map[string]RuntimeValue),
	}
}

func (env *Environment) DeclareVariable(name string, value RuntimeValue) RuntimeValue {
	if _, ok := env.Values[name]; ok {
		log.Fatalf("Variable %s already declared", name)
	}
	env.Values[name] = value
	return value
}

func (env *Environment) SetVariable(name string, value RuntimeValue) RuntimeValue {
	resEnv := env.resolve(name)
	resEnv.Values[name] = value
	return value
}

func (env *Environment) GetVariable(name string) RuntimeValue {
	resEnv := env.resolve(name)
	if resEnv == nil {
		log.Fatalf("Variable %s not declared", name)
	}
	return resEnv.Values[name]
}

func (env *Environment) resolve(name string) *Environment {
	if _, ok := env.Values[name]; ok {
		return env
	}
	if env.Parent == nil {
		return nil
	}
	return env.Parent.resolve(name)
}
