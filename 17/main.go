package main

import (
	"fmt"
	"reflect"
)

func main() {
	interfaceSlice := make([]interface{}, 0)

	interfaceSlice = append(interfaceSlice, true)
	interfaceSlice = append(interfaceSlice, 50)
	interfaceSlice = append(interfaceSlice, make(chan int))
	interfaceSlice = append(interfaceSlice, "hello")

	for i, v := range interfaceSlice {
		elemType := reflect.ValueOf(v).Kind()
		switch elemType {
		case reflect.String:
			fmt.Printf("%d elem of slice is string\n", i)
		case reflect.Int:
			fmt.Printf("%d elem of slice is int\n", i)
		case reflect.Chan:
			fmt.Printf("%d elem of slice is chan\n", i)
		case reflect.Bool:
			fmt.Printf("%d elem of slice is bool\n", i)
		}
	}

}
