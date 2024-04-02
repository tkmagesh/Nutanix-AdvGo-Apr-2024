package main

import "fmt"

type MyStr string

func (s MyStr) Length() int {
	return len(s)
}

func main() {
	str := MyStr("Non culpa minim ex quis eu enim. Mollit qui sint velit pariatur pariatur. Aliqua mollit nostrud do aliquip in duis id consectetur ullamco cupidatat enim excepteur. Esse laboris Lorem consequat dolor et id laborum dolore Lorem. Amet nulla elit velit ex officia proident minim adipisicing.")
	fmt.Println(str.Length())
}
