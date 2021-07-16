// Copyright 2021 Arman Abkar. All rights reserved.

// Package is workspace (or project)
package main // 'main' is executable package, other packages are reusable
// use import for using other packages

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Declaring Variables and Constants
	var age int8 = 22
	name, email := "Arman Abkar", "armanabkar@gmail.com" // shorthand - only for initialization
	var complexNum complex128 = complex(2, 3)
	const isCool = true // value is known at compile time
	const (
		A int64 = iota
		B
		c
	) // each constant is assigned to a unique int

	// Declaring Types
	type Celsius float64
	type Deck []string
	var temperature Celsius = 34.0

	fmt.Println(name, email, temperature, complexNum)

	// Print type of variable
	fmt.Printf("%T\n", age)

	// 8bit UTF codes are same as ASCII - Unicode is 32bit char
	greet := "Hi There!" // string literal

	// Type Conversion
	var x int32 = 1
	var y int16 = 2
	x = int32(y)
	// Byte Slice -> represents a string! [72 105]
	fmt.Println([]byte(greet))

	// Math Package
	fmt.Println(math.Floor(2.7), greet, x, y)
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(2.7))

	// extracting multiple return values
	greetName, greetSentence := greeting("Arman")
	fmt.Println(greetName, greetSentence)

	// Arrays - have fixed length
	var fruitArr [2]string
	fruitArr[0] = "Apple"
	fruitArr[1] = "Banana"
	peoples := [2]string{"Arman", "Alin"}
	fmt.Println(peoples)
	fmt.Println(fruitArr[0])

	// Slices - dynamic length, contains a pointer to the array(Go will create a Array in memory)
	copyFruitArr := fruitArr[0:1]
	copyFruitArr = append(copyFruitArr, "Cherry")
	fmt.Println(copyFruitArr)
	peoplesSlice := []string{"Arman", "Alin", "Jessie"}
	peoplesSlice = append(peoplesSlice, "Sogol")
	sli1 := peoplesSlice[:1] // range in slice
	fmt.Println(len(peoplesSlice), cap(peoplesSlice), sli1)
	sli2 := make([]int, 10, 15)
	// Iterating through Slice
	for i, name := range peoplesSlice {
		fmt.Println(i, name)
	}
	peoplesSlice = append(peoplesSlice, copyFruitArr...) // unpack list values
	fmt.Println(peoplesSlice)

	// Strings Package
	joinedPeople := strings.Join([]string(peoplesSlice), ",")
	fmt.Println(joinedPeople)

	// Conditionals / Control Structures
	x2 := 5
	y2 := 10
	if x2 < y2 && y2 > 0 {
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
	for i := 1; i <= 30; i++ {
		if i%15 == 0 {
			fmt.Print("FizzBuzz, ")
		} else if i%3 == 0 {
			fmt.Print("Fizz, ")
		} else if i%5 == 0 {
			fmt.Print("Buzz, ")
		} else {
			fmt.Print(i, ", ")
		}
	}

	// Maps - key & Value pairs (and same type), dynamic
	colors1 := map[string]string{
		"red":   "#ff0000",
		"white": "#ffffff",
		"black": "#000000",
	}
	var colors2 map[string]string
	fmt.Println(colors1, colors2)
	emails := make(map[string]string)
	// use [] to access values of Maps
	emails["Arman"] = "armanabkar@gmail.com"
	fmt.Println(emails["Arman"])
	delete(emails, "Arman")
	ages := map[string]int{"Arman": 22, "Alin": 16} // Map Literal
	ages["Jessie"] = 3
	delete(ages, "Sogol")
	fmt.Println(ages["Arman"])
	// Iterating through Map
	for key, val := range colors1 {
		fmt.Println("Hex code for", key, "is", val)
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
	// RAM = Addresses + Values
	// Everything in Go is passed by value, so everytime Go find new address and assign the value to it (without pointers)
	a := 5
	// Turn value into address with '&value'
	b := &a
	// b will print memory address of a
	fmt.Println(a, b)
	// Turn address into value with '*address'
	fmt.Println(*b)
	// Change value with pointer
	*b = 10
	fmt.Println(a)
	// new() creates var and returns a pointer
	ptr := new(int)
	*ptr = 3

	// Value Types: int, float, string, bool, struct, Arrays -> Use pointers to change these things in a function

	// Reference Types: slices, maps, channels, pointers, functions -> Don't worry about pointers

	sumClosure := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(sumClosure(i))
	}

	// similar to TypeScript classes (or interfaces) or objects in JS
	person1 := person{firstName: "Arman", lastName: "Abkar", city: "Los Angeles", age: 22}
	// person1 := person{"Arman", "Abkar", "Los Angeles", 22}
	fmt.Println(person1)
	fmt.Println(person1.city)

	person1.hasBirthday()
	person1.hasBirthday()
	fmt.Println(person1.greet())

	saveToFile("HelloWorld.txt")
	fmt.Println(readFile("HelloWorld.txt"))

	// Random Number Generation
	source := rand.NewSource(time.Now().UnixNano()) // with using time package
	random := rand.New(source)
	randomNumber := random.Intn(len(greet) - 1)
	fmt.Println(randomNumber)

	var person2 person
	fmt.Println(person2)       // Default values: string="", int=0, bool=false
	person2.firstName = "Alin" // updating values
	fmt.Printf("%+v", person2)
	employee1 := employee{salary: 80000, person: person{firstName: "Sogol", lastName: "Abkar", city: "Seattle", age: 35}}
	fmt.Printf("%+v", employee1)

	person3 := *newPerson("Sogol", "Safieddin", "Isfahan", 30)
	fmt.Println(person3.lastName)

	eb := englishBot{}
	gb := germanBot{}
	printGreeting(eb)
	printGreeting(gb)

	// Web & 'http' package
	// http.HandleFunc("/", index)
	// fmt.Println("Server listening on port 3000")
	// http.ListenAndServe(":3000", nil)
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// Reader Interface
	byteSlice2 := make([]byte, 32*1024)
	resp.Body.Read(byteSlice2)

	// io.Copy function; will pass byte slice to writer interface and then source of output that implements writer interface
	io.Copy(os.Stdout, resp.Body)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

// Interfaces
// implicit, are not generic types, contract to help us manage types
type bot interface {
	getGreeting() string
	// getGoodbye(string) string, error
}
type englishBot struct{}
type germanBot struct{}

func (englishBot) getGreeting() string {
	return "\nHi There!"
}
func (germanBot) getGreeting() string {
	return "Hallo!"
}

// use interface in this function, now any type that have getGreeting function receiver, can use this function!
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// englishBot, string, struct, int, ... are Concrete Types -> can create value directly

// Methods (Receiver Function)
// by convension use first letter of type in receivers (person -> p)
// Value Receivers
func (p person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " from " + p.city + " and I'm " + strconv.Itoa(p.age)
}

// Pointer Receivers - for changing something
// *person is type description; we're working with a pointer to a person
func (pointerToPerson *person) hasBirthday() {
	// Void func - have no return type
	// *pointerToPerson -> operator; we want to manipulate the value the pointer is referring
	pointerToPerson.age++
}

// Creating person instance with functions
func newPerson(firstName string, lastName string, city string, age int) *person {
	newPerson := person{firstName, lastName, city, age}
	return &newPerson
}

// Struct -> compose data types, Data Structure/Collection of properties that are related together
// no indexing, fields (can contain different types) should be known at complie time, use to represent things, value type!
// Exported struct should be UpperCase
type person struct {
	firstName, lastName string
	city                string
	age                 int
}

// Embedding struct into another struct
type employee struct {
	person
	salary int
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
// func funcName(arg argType) returnType {}
func greeting(name string) (string, string) {
	defer fmt.Println("Deferred Bye!") // deferred until the surrunding func completes

	// return multiple values
	return "Hello " + name, "Good Morning!"
}
func add(num1 int, num2 int) (sum int) {
	// named return value (sum)
	sum = num1 + num2
	return
}

// Read/Write from HardDrive
func saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte("Hello World from HelloWorld.txt"), 0666)
}
func readFile(filename string) string {
	byteSlice, err := ioutil.ReadFile(filename)

	// Error Handling
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1) // quit the program
	}

	return string(byteSlice)
}

// Using 3rd party modules
// install using "go get github.com/common-nighthawk/go-figure" or paste this in import() and run "go mod tidy" to fetch

// go run (main.go) -> compile and execute one or two files
// go build -> just compiles the executable file and moves it to the destination. go install does a little bit more. It moves the executable file to $GOPATH/bin and caches all non-main packages which are imported to $GOPATH/pkg. The cache will be used during the next compilation provided the source did not change yet.
// go install -> compile and install a package
// go get -> downloads the raw source code of a package
// Other commands: go test, go fmt
