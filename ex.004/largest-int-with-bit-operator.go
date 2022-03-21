package main

import "fmt"

func main() {
	fmt.Printf("%016X\n", uint(0))
	fmt.Printf("%016X\n", ^uint(0))
	fmt.Printf("%016X\n", ^uint(0)>>1)
}