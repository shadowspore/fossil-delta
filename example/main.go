package main

import (
	"fmt"

	fdelta "github.com/axvq/fossil-delta"
)

func main() {
	origin := []byte("abcdefghijklmnopqrstuvwxyz1234567890")
	target := []byte("abcdefghijklmnopqrstuvwxyz0987654321")

	// Create delta
	delta := fdelta.Create(origin, target)

	// Create target by patching origin with delta
	patched, err := fdelta.Apply(origin, delta)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Origin: `%s`\n", origin)
	fmt.Printf("Target: `%s`\n", target)
	fmt.Printf("Delta : `%s`\n", delta)
	fmt.Printf("Patch : `%s`\n", patched)
}
