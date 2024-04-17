package main

import (
	"fmt"
)

var (
	Number string
	Commit string
	Date   string
)

func main() {
	fmt.Printf("Version: %s\n", Number)
	fmt.Printf("Commit: %s\n", Commit)
	fmt.Printf("Date: %s\n", Date)
}
