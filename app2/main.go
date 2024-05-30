package main

import "fmt"

func main() {
	name := "SHIVAJI BURGULA"
	id := "500224628"
	age := "25"

	greeting(name, id, age)
	//fmt.Println("Hello from SHIVAJI BURGULA 500224628")
}
func greeting(name string, id string, age string) {
	fmt.Println(name, id, "you are", age, "years old.")
}
