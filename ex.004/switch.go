package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		switch {
		case i < 5:
			fmt.Printf("Haha %v\n", i)
			fallthrough
		case i%2 != 0:
			if i >= 5 {
				break
			}
			fmt.Printf("Hello %v\n", i)
		case i%2 == 0:
			if i < 5 {
				break
			}
			fmt.Printf("Hello %v\n", i)
		}
	}
}