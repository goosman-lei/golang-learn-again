package main

import (
	"tec-inf.com/foo1"
)

type Bar struct {
	*foo1.Foo
	Val int
}

func main() {
	var b Bar = Bar{Val: 1, Foo: &foo1.Foo{IVal: 2}}

	b.WhoAmI()
}