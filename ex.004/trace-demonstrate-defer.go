package main

import "fmt"

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a")) // 03. entering:a
	fmt.Println("in a")  // 04. in a
	// 05. leaving:a
}

func b() {
	defer un(trace("b")) // 01. entering:b
	fmt.Println("in b")  // 02. in b
	a()
	// 06. leaving:b
}

func main() {
	b()
}