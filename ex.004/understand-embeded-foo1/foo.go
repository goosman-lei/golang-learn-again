package foo1

import (
	"fmt"
	"reflect"
)

type Foo struct {
	IVal int
}

func (f *Foo) WhoAmI() {
	var (
		r_f   reflect.Value = reflect.ValueOf(f)
		r_f_t reflect.Type  = r_f.Type()
	)
	fmt.Printf("foo1.WhoAmI() %#v\n", r_f)
	fmt.Printf("foo1.WhoAmI() %#v\n", r_f_t)
	fmt.Printf("foo1.WhoAmI() %s\n", r_f_t)
	fmt.Printf("foo1.WhoAmI() %s.%s\n", r_f_t.PkgPath(), r_f_t.Name())
}