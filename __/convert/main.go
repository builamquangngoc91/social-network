package main

import "fmt"

type A struct{}

func (a A) Error() string {
	return ""
}

type B struct{}

func (b B) Error() string {
	return ""
}

func main() {
	var a error
	a = A{}

	b, _ := a.(B)
	fmt.Println(b)
}
