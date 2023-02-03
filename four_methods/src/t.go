package main

import "fmt"

func main() {
	t := "this is"
	for _, c := range t {
		fmt.Printf("%v", c)
	}
}
