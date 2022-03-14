package main

import (
	"fmt"
	"log"

	"tec-inf.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	messages, err := greetings.Hellos([]string{"Gladys", "Samantha", "Darrin"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}