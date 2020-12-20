package main

import "fmt"

func main() {
	SingleParam(1, 3, "a")
}

// SingleParam export single param
func SingleParam(p, f float64, s string) {
	fmt.Printf("%T %T %T", p, s, f)
}
