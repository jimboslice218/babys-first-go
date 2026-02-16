package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)
// import "rsc.io/quote"

func main() {
	// fmt.Println("Hello World!")
	names := []string{"Gladys", "Samantha", "Darrin"}
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}