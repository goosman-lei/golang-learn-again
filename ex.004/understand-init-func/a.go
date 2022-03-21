package understandinit

import "fmt"

func init() {
	fmt.Printf("First init func in a\n")
}
func init() {
	fmt.Printf("Second init func in a\n")
}

func A() {
	fmt.Printf("understandinit.a\n")
}