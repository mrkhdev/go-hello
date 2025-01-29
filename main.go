package main

import "fmt"

var (
	Version = "v0.0.0"
	Commit  = "none"
)

func main() {
	fmt.Printf("Hello, World from version %s (commit: %s)! \n", Version, Commit)

	fmt.Println("Feature change. Antoher!!")
}
