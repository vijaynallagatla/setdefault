package main

import (
	"fmt"
	. "github.com/vijaynallagatla/bind-default"
)

type Example struct {
	Name string `default:"John Doe"`
	Age  int    `default:"30"`
}

func main() {
	example := Example{}
	BindDefault(&example)
	fmt.Println(example.Name) // Output: John Doe
	fmt.Println(example.Age)  // Output: 30
}
