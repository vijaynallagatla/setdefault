package main

import (
	"fmt"
	"github.com/vijaynallagatla/setdefault/v1"
)

type Example struct {
	Name       string `default:"John Doe"`
	Age        int    `default:"30"`
	Profession string `default:"Engineer"`
}

func main() {
	example := Example{}
	setdefault.Bind(&example)
	fmt.Println(example.Name)       // Output: John Doe
	fmt.Println(example.Age)        // Output: 30
	fmt.Println(example.Profession) // Output: Engineer
}
