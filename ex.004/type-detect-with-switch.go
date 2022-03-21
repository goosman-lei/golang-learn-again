package main

import "fmt"

func main() {
	var (
		v_int     int     = 3
		v_bool    bool    = true
		v_float32 float32 = float32(3.1415926)
		v_float64 float64 = float64(1.44)
	)

	printType(v_int)
	printType(v_bool)
	printType(v_float32)
	printType(v_float64)
	printType(&v_int)
	printType(&v_bool)
	printType(&v_float32)
	printType(&v_float64)
	printType(int32(v_int))
}

func printType(v interface{}) {
	switch t := v.(type) {
	default:
		fmt.Printf("unexcepted type %T\n", t)
	case bool:
		fmt.Printf("boolean %T %v\n", t, v)
	case *bool:
		fmt.Printf("*boolean %T %v\n", t, v)
	case int:
		fmt.Printf("int %T %v\n", t, v)
	case *int:
		fmt.Printf("*int %T %v\n", t, v)
	case float32:
		fmt.Printf("float32 %T %v\n", t, v)
	case float64:
		fmt.Printf("float64 %T %v\n", t, v)
	}
}