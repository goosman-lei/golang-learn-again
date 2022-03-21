package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("/not-exists-file-path")
	fmt.Printf("%##v\n", err)
	if t, ok := err.(*os.PathError); ok {
		fmt.Printf("%##v %##v\n", t, ok)
	} else {
		fmt.Printf("%##v %##v\n", t, ok)
	}
}