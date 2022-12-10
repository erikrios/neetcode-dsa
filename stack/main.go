package main

import (
	"fmt"
	"stack/stack"
)

func main() {
	s := stack.New[string]()
	s.Push("Hello")
	s.Push("World")
	v, err := s.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
	v, err = s.Peek()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
	s.Push("Erik")
	v, err = s.Peek()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
}
