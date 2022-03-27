package main

import (
	"fmt"
	"reflect"
)

func main() {
	var t interface{}
	t = true
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	}

	m := make(map[string]interface{})
	m["a"] = [4]int{1, 2, 3, 4}
	m["b"] = []int{1, 2, 3, 4}

	for k, v := range m {
		rt := reflect.TypeOf(v)
		switch rt.Kind() {
		case reflect.Array:
			fmt.Printf("m[%s]: ARRAY\n", k)
		case reflect.Slice:
			fmt.Printf("m[%s]: SLICE\n", k)
		default:
			fmt.Printf("NONT")
		}
	}
}
