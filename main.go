package main

import "fmt"

var (
	Version = "v0.0.0"
	Commit  = "none"
)

func main() {
	fmt.Printf("Hello, World from version %s (%s)\n", Version, Commit)

	fmt.Println("bugfix feature")
	fmt.Println("Nieuwe feature")

	fmt.Println("bugfix feature")
	fmt.Println("Nieuwe feature")
	fmt.Println("Nieuwe feadfture")
	fmt.Println("Nieusswe feature")
}
