package main

import (
	"fmt"
	"github.com/hassanjawwad12/betterjson"
)

func main() {
	// Example usage of MinifyJ
	jsonStr := `{
		"name": "Alice",
		"age": 25,
		"city": "Wonderland"
	}`

	minified, err := betterj.MinifyJ(jsonStr)
	if err != nil {
		fmt.Println("Error minifying JSON:", err)
		return
	}
	fmt.Println("Minified JSON:", minified)

	// Example usage of BeautifyJ
	beautified, err := betterj.BeautifyJ(jsonStr, "  ")
	if err != nil {
		fmt.Println("Error beautifying JSON:", err)
		return
	}
	fmt.Println("Beautified JSON:\n", beautified)
}
