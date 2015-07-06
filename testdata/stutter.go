// Test of stuttery names.

// Package donut ...
package donut

// DonutMaker makes donuts.
type DonutMaker struct{} // MATCH /donut\.DonutMaker.*stutter/

// DonutRank computes the ranking of a donut.
func DonutRank(d Donut) int { // MATCH /donut\.DonutRank.*stutter/
	return 0
}

// Donut is a delicious treat.
type Donut interface{} // MATCH /interface name must have prefix 'I'/

// Donuts are great, aren't they?
type Donuts []Donut // ok because it didn't start a new word

type donutGlaze int // ok because it is unexported
