package main

import "fmt"

type A struct {
}

func (a A) Name() {
	fmt.Println("A")
}

type B struct {
	A
}

func (b B) Name() {
	fmt.Println("B")
}

func main() {
	b := B{}
	b.Name()
}
