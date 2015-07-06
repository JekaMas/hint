// Test for bad receiver names.

// Package foo ...
package foo

type foo struct{}

func (this foo) f1() {
}

func (self foo) f2() { // MATCH /receiver name should be 'this'/
}

func (f foo) f3() { // MATCH /receiver name should be 'this'/
}

func (foo) f4() {
}

type bar struct{}

func (b bar) f1() { // MATCH /receiver name should be 'this'/
}

func (b bar) f2() { // MATCH /receiver name should be 'this'/
}

func (a bar) f3() { // MATCH /receiver name should be 'this'/
}

func (a *bar) f4() { // MATCH /receiver name should be 'this'/
}

func (b *bar) f5() { // MATCH /receiver name should be 'this'/
}

func (bar) f6() {
}

func (_ *bar) f7() { // MATCH /receiver name should be 'this'/
}
