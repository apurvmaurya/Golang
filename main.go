package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{"Apurv", "Maurya", 18}
	p2 := person{"Apurv", "Maurya", 18}

	p := []person{p1, p2}

	b, err := json.Marshal(p)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))
}
