package main

import "fmt"

type Incrementer struct {
	val int
}

func (incr *Incrementer) incrWithPointer() {
	fmt.Printf("In incrWithPointer. Before ++: %p %d\n", incr, incr.val)
	incr.val++
	fmt.Printf("In incrWithPointer. After ++: %p %d\n", incr, incr.val)
}

func (incr Incrementer) incrWithValue() {
	fmt.Printf("In incrWithValue. Before ++: %p %d\n", &incr, incr.val)
	incr.val++
	fmt.Printf("In incrWithValue. After ++: %p %d\n", &incr, incr.val)
}

func main() {
	var (
		incrValue   Incrementer  = Incrementer{val: 1}
		incrPointer *Incrementer = &incrValue
	)

	fmt.Printf("Init state\n")
	fmt.Printf("%p %v\n", &incrValue, incrValue)
	fmt.Printf("%p %v\n", incrPointer, *incrPointer)

	fmt.Printf("\nCall incrWithValue By value\n")
	incrValue.incrWithValue()
	fmt.Printf("%p %v\n", &incrValue, incrValue)
	fmt.Printf("%p %v\n", incrPointer, *incrPointer)

	fmt.Printf("\nCall incrWithValue By pointer\n")
	incrPointer.incrWithValue()
	fmt.Printf("%p %v\n", &incrValue, incrValue)
	fmt.Printf("%p %v\n", incrPointer, *incrPointer)

	fmt.Printf("\nCall incrWithPointer By pointer\n")
	incrPointer.incrWithPointer()
	fmt.Printf("%p %v\n", &incrValue, incrValue)
	fmt.Printf("%p %v\n", incrPointer, *incrPointer)

	fmt.Printf("\nCall incrWithPointer By value\n")
	incrValue.incrWithPointer()
	fmt.Printf("%p %v\n", &incrValue, incrValue)
	fmt.Printf("%p %v\n", incrPointer, *incrPointer)
}