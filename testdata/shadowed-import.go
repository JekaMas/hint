// Test for shadowed import names.

// Package foo ...
package foo

import (
	"fmt"
	"time"
)

func time() int { // MATCH /variable name should not shadow package name/
	var fmt string // MATCH /variable name should not shadow package name/
	fmt.Print(time.Now())
	time := 0 // MATCH /variable name should not shadow package name/

	g := map[string]string{}

	for _, time := range g { // MATCH /variable name should not shadow package name/
	}

	for fmt := range g { // MATCH /variable name should not shadow package name/
	}
	return time
}
