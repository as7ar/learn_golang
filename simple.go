package main

import (
	"fmt"
	"reflect"
)

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func genPerson(name string) *Person {
	p := Person{Name: name}
	p.Age = 32
	return &p
}

func change(x *int) {
	*x = 99
}

func plus(a int, b, c float32) float32 {
	fmt.Println(reflect.TypeOf(b))

	return float32(a) + b + c
}

func vals() (int, int) {
	return 3, 7
}

func sum[T Number](nums ...T) T {
	var total T = 0.0
	for _, num := range nums {
		total += num
	}
	return total
}
