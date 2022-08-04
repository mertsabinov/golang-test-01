package main

import "fmt"

const perfix = "Hello "

func Greeting(name string) string {
	return perfix + name
}

func main() {
	fmt.Printf(Greeting("Mert"))
}
