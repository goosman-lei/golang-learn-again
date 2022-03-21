package understandinit

import "fmt"

func init() {
	fmt.Printf("First init func in b\n")
}
func init() {
	fmt.Printf("Second init func in b\n")
}

func B() {
	fmt.Printf("understandinit.b\n")
}
