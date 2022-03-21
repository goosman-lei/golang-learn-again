package foo2

import (
	"fmt"
	"reflect"
)

type Foo struct {
}

func (f *Foo) WhoAmI() {
	var (
		f_t reflect.Type = reflect.TypeOf(f)
	)
	fmt.Printf("foo2.WhoAmI() %s.%s\n", f_t.PkgPath(), f_t.Name())
}