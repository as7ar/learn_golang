package main

import (
	"LearnGo/geometry"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Number interface {
	int | float64
}

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

	// func

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

	// Pointer

	x := 10
	p := &x
	fmt.Println(x, p)
	change(p)
	fmt.Println(x, p)

	// Struct

	person := Person{
		Name: "Steve",
		Age:  20,
		Job:  "Student",
	}
	fmt.Println(person.Name)

	fmt.Println(genPerson("Alex").Name)

	rectangle := geometry.Rect{Width: 6, Height: 12}
	fmt.Println("rec width, height:", rectangle.Height, rectangle.Width, "\narea:", rectangle.Area(), "\nperim:", rectangle.Perim())

	var color Color = 1
	fmt.Println(color.String())

	for file := range FindLogFiles("../") { // hmmm. Can't Understand
		fmt.Println("found:", file)

		if strings.Contains(file, "error") {
			break
		}
	}

	for x := range Count(5) {
		fmt.Println(x)
	}

	// Goroutine

	f("direct")

	go f("goroutine")

	go func(msg string) { fmt.Println(msg) }("TEST")

	time.Sleep(time.Second)

	// Channel
	done := make(chan bool)

	message := make(chan string, 3)
	go func() {
		message <- "hello"
		done <- true
	}()
	message <- "?"
	fmt.Println(<-message)
	fmt.Println(<-message)

	<-done

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "pass")
	pong(pings, pongs)
	fmt.Println(<-pongs)

	// select

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	i2 := 0
	for {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
			i2 += 1
			fmt.Println(i2)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
			i2 += 1
			fmt.Println(i2)
		}
		if i2 >= 2 {
			break
		}
	}

}

func ping(pings chan<- string, msg string) {
	pings <- msg
}
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}
