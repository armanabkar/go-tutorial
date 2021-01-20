package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

func main() {
	// Declaring Variables
	name, email := "Hello World", "armanabkar@gmail.com"
	var age int32 = 22
	const isCool = true

	fmt.Println(name, email)

	// Print type of variable
	fmt.Printf("%T\n", age)

	// Math Package
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(2.7))

	fmt.Println(greeting("Arman"))

	// Arrays - have fixed length
	var fruitArr [2]string
	fruitArr[0] = "Apple"
	fruitArr[1] = "Banana"
	peoples := [2]string{"Arman", "Alin"}
	fmt.Println(peoples)

	// Slices - Don't have fixed length
	peoplesSlice := []string{"Arman", "Alin"}
	fmt.Println(len(peoplesSlice))

	// Conditionals
	x := 5
	y := 10
	if x < y {
		fmt.Printf("%d is less than %d", x, y)
	} else {
		fmt.Printf("%d is bigger than %d", x, y)
	}

	color := "red"
	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	}

	// Loops
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}

	// FizzBuzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	// Maps - key & Value pairs
	emails := make(map[string]string)
	emails["Arman"] = "armanabkar@gmail.com"
	fmt.Println(emails["Arman"])
	delete(emails, "Arman")
	ages := map[string]int{"Arman": 22}
	fmt.Println(ages["Arman"])

	// Range - and loop through ids
	ids := []int{1, 41, 5, 2, 3}
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}
	// if you don't want index use _
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum: ", sum)

	// Pointers - for better performance
	// Everything in Go is passed by value
	a := 5
	b := &a
	// b will print memory address
	fmt.Println(a, b)
	// Use * to read value from address
	fmt.Println(*b)
	// Change value with pointer
	*b = 10
	fmt.Println(a)

	sumClosure := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(sumClosure(i))
	}

	// similar to javascript classes or interfaces
	person1 := person{firstName: "Arman", lastName: "Abkar", city: "Los Angeles", age: 22}
	// person1 := person{"Arman", "Abkar", "Los Angeles", 22}
	fmt.Println(person1)
	fmt.Println(person1.city)

	person1.hasBirthday()
	person1.hasBirthday()
	fmt.Println(person1.greet())

	// Web
	http.HandleFunc("/", index)
	fmt.Println("Server listening on port 3000")
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

// Interfaces
type shape interface {
	area() float64
}
type circle struct {
	x, y, radius float64
}
type rectangle struct {
	width, height float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s shape) float64 {
	return s.area()
}

// Methods
// Value Receivers
func (p person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " from " + p.city + " and I'm " + strconv.Itoa(p.age)
}

// Pointer Receivers - for changing something
func (p *person) hasBirthday() {
	// Void func - have no return type
	p.age++
}

// Struct
// Exported struct should be UpperCase
type person struct {
	firstName, lastName string
	city                string
	age                 int
}

// Closures
// Anonymous function
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// Delcaring Functions
func greeting(name string) string {
	return "Hello " + name
}

// go build just compiles the executable file and moves it to the destination. go install does a little bit more. It moves the executable file to $GOPATH/bin and caches all non-main packages which are imported to $GOPATH/pkg. The cache will be used during the next compilation provided the source did not change yet.
