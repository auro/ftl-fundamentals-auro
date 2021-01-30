package main

import (
	"calculator"
	"fmt"
)

func main() {
	fmt.Println(calculator.Subtract(5, 3))
}

// SingleParam export single param
func SingleParam(p, f float64, s string) {
	fmt.Printf("%T %T %T", p, s, f)
}
