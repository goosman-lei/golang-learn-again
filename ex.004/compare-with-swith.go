package main

import "fmt"

func main() {
	fmt.Printf("compare(%v, %v) = %d\n", "1234", "1234", compare([]byte("1234"), []byte("1234")))
	fmt.Printf("compare(%v, %v) = %d\n", "1234", "12345", compare([]byte("1234"), []byte("12345")))
	fmt.Printf("compare(%v, %v) = %d\n", "12345", "1234", compare([]byte("12345"), []byte("1234")))
	fmt.Printf("compare(%v, %v) = %d\n", "abcdef", "abcfde", compare([]byte("abcdef"), []byte("abcfde")))
}

func compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}
	return 0
}