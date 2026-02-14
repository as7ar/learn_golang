package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	fmt.Println("Hello", "world")
	a := 2
	a = 3

	b, c := 4.0, "a"
	fmt.Println(a, b, c)

	const e = "abc"
	b = 1 / 3e4
	fmt.Println(e, b)

	for i := 0; i <= 10; i++ {
		fmt.Print(i)
	}

	test := ""
	for i := range 10 {
		test += strconv.Itoa(i)
	}
	fmt.Println(test)

	for range 5 {
		fmt.Println("aa")
	}

	fmt.Println("\033[41mTEST\033[0m")

	str1 := "a"
	if str1 == "a" {
		fmt.Println("str1 is a")
	} else {
		fmt.Println("str1 is not a")
	}

	switch str1 {
	case "a":
		fmt.Println("A!")
	default:
		fmt.Println("WHAT?")
	}

	var list [3]int
	list = [...]int{1, 2, 3}
	fmt.Println(list)

	list2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", list2, "length:", len(list2))

	list3 := [2][2]float32{
		{1, 2},
		{3, 4},
	}
	fmt.Println(list3)

	slice := make([]string, 3)
	fmt.Println(slice)
	slice[0] = "one"
	//slice[3] = "four"
	fmt.Println(slice)

	map1 := make(map[string]float32)
	map1["one"] = 1.0
	map1["two"] = 2.0
	fmt.Println(map1)

	fmt.Println(plus(1, 2.0, 3.0))

	_, vals2 := vals()
	fmt.Println(vals2)

	fmt.Println(sum(1, 2, 3))

	fmt.Println("5!:", fact(5))

	var fibonacci func(n int) int
	fibonacci = func(n int) int {
		if n <= 1 {
			return n
		}
		return fibonacci(n-1) + fibonacci(n-2)
	}

	fmt.Println("Fibonacci 10st term of the  sequence", fibonacci(7))

	for i, c := range "go" {
		fmt.Println(i, c)
	}

	for k, v := range map1 {
		fmt.Println(k, v)
	}

	x := 10
	p := &x
	fmt.Println(x, p)
	change(p)
	fmt.Println(x, p)
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

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
