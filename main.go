package main

import (
	"fmt"

	"github.com/gitApurv/puppy"
)

func main() {
	fmt.Println("Hello Gophers")
	fmt.Println(puppy.PuppyName)

	fmt.Println(puppy.Family("temp"))

	fmt.Println("Version 1.2.0")
}
