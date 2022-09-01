package main

import "fmt"

func printThis(s string) {
	fmt.Println(s)
}

func main() {
	go printThis("This")
	printThis("That")
}
