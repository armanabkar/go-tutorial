package main

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// Declaring Variables
	var age int32 = 22
	name, email := "Hello World", "armanabkar@gmail.com"
	var complexNum complex128 = complex(2, 3)
	const isCool = true // value is known at compile time
	const (
		A int64 = iota
		B
		c
	) // each constant is assigned to a unique int

	// Declaring Types
	type Celsius float64
	var temperature Celsius = 34.0

	fmt.Println(name, email, temperature, complexNum)

	// Print type of variable
	fmt.Printf("%T\n", age)

	// Convert Type
	var x int32 = 1
	var y int16 = 2
	x = int32(y)

	// 8bit UTF codes are same as ASCII - Unicode is 32bit char
	greet := "Hi There" // string literal

	// Math Package
	fmt.Println(math.Floor(2.7), greet, x, y)
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(2.7))

	fmt.Println(greeting("Arman"))

	// Arrays - have fixed length
	var fruitArr [2]string
	fruitArr[0] = "Apple"
	fruitArr[1] = "Banana"
	peoples := [2]string{"Arman", "Alin"}
	fmt.Println(peoples)

	// Slices - Don't have fixed length, contains a pointer to the array
	peoplesSlice := []string{"Arman", "Alin", "Jessie"}
	sli1 := peoplesSlice[0:1]
	fmt.Println(len(peoplesSlice), cap(peoplesSlice), sli1)
	sli2 := make([]int, 10, 15)

	// Conditionals / Control Structures
	x2 := 5
	y2 := 10
	if x2 < y2 {
		fmt.Printf("%d is less than %d", x2, y2)
	} else {
		fmt.Printf("%d is bigger than %d", x2, y2)
	}
	// Switch Statements - we don't need 'break'
	color := "red"
	switch color {
	case "red":
		fmt.Println("color is red", sli2)
	case "blue":
		fmt.Println("color is blue")
	}
	// tagless switch
	switch {
	case x > 1:
		fmt.Println("case1")
	}

	// Loops
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	// similar to while
	j := 0
	for j < 3 {
		fmt.Println(("hi"))
		j++
		if j == 5 {
			break // will exit the loop
		} else if j == 2 {
			continue // won't print 2
		}
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
	ages := map[string]int{"Arman": 22, "Alin": 16} // Map Literal
	ages["Jessie"] = 3
	delete(ages, "Sogol")
	fmt.Println(ages["Arman"])
	// Iterating through Map
	for key, val := range ages {
		fmt.Println(key, val)
	}

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

	// Pointers - is an address to data in memory, for better performance
	// Everything in Go is passed by value
	a := 5
	// & return addres of
	b := &a
	// b will print memory address
	fmt.Println(a, b)
	// Use * to read value from address
	fmt.Println(*b)
	// Change value with pointer
	*b = 10
	fmt.Println(a)
	// new() creates var and returns a pointer
	ptr := new(int)
	*ptr = 3

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

	// Error Handling
	f, err := os.Open("/text.txt")
	if err != nil {
		fmt.Println((err))
		return
	}
	fmt.Println(f)

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

// Struct -> compose data types
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
	defer fmt.Println("Bye!") // deferred until the surrunding func completes

	return "Hello " + name
}

// go build just compiles the executable file and moves it to the destination. go install does a little bit more. It moves the executable file to $GOPATH/bin and caches all non-main packages which are imported to $GOPATH/pkg. The cache will be used during the next compilation provided the source did not change yet.
