package main

// We'd look into the following:
// - Naming Convention
// - Typed Constants
// - Untyped Constants
// - The difference between both
// - Enumerated Constants (generating constants)
// - Enumerated Expressions

import (
	"fmt"
	// "math"
)

const e int16 = 27

const i = iota

func main() {
	// - Naming a Constant: Starts with a const keyword
	// Note that in go, exporting a variable is starting the first letter with uppercase; e.g MyConst
	// else using const is enough to name a constant and can start with a lowercase myConst

	// A typed constant is created similarly to a typed variable eg
	const myConst int = 24
	fmt.Println(myConst) // 24
	// Note: a constant is immutable; thus cannot be reassigned

	// Another characteristics of a constant is, it has to be assignable at compile time. e.g:
	// const myNewConst float64 = math.Sin(1.57) // got "math.Sin(1.57) (value of type float64) is not constantcompiler (InvalidConstInit)"

	// Note constant can be made up of any of the primitive types.
	const a int = 14
	const b string = "foo"
	const c float32 = 3.14
	const d bool = true

	fmt.Printf("%v\n", a) // 14
	fmt.Printf("%v\n", b) // foo
	fmt.Printf("%v\n", c) // 3.14
	fmt.Printf("%v\n", d) // true

	// Inner constants shadows outer constant e.g: I named a const "e" at package level
	// in we run the print before reassigning, we get:
	fmt.Printf("%v, %T\n", e, e) // 27, int16
	const e int = 16
	fmt.Printf("%v, %T\n", e, e) // 16, int
	// Note: the collection types are mutuable

	// apart from the characteristics of being immutable, a constant and variables can be done the same way eg
	var f int = 27
	fmt.Printf("%v, %T\n", a+f, a+f) // 41, int

	// a sweet charateristics of constant is that without stating a type, the compiler will use the value literally e.g
	const g = 42
	var h int16 = 27
	fmt.Printf("%v, %T\n", g+h, g+h) // 69, int16
	// what it does is it replaces "g" with 42 as "42 + h2"

	// Next is Enumerated Constant
	// at package level i named a const "i"
	fmt.Printf("%v, %T\n", i, i) // 0, int
	// what is iota? iota is a counter used when creating an enumerated constant
	// another feature of a constant is that we can assign them in blocks

	const (
		j = iota
		k = iota
	)
	fmt.Printf("%v\n", j) // 0
	fmt.Printf("%v\n", k) // 1
	// we see that the values increase in the block
	// Another feature of const block is if you don't assign a value to the next const
	// go compiler will infer the value. e.g:
	const (
		l = iota
		m
	)
	fmt.Printf("%v\n", l) // 0
	fmt.Printf("%v\n", m) // 1
	// Note: An iota is scoped to a constant block and with each block, it resets to 0

	// _ is go write only variable symbol
	const (
		_  = iota             // ignore the first value by assigning to blank identifier
		KB = 1 << (10 * iota) // 2^0 * 2^10  // 1024
		MG                    // 2^0 * 2^ 10
		GB
		TB
		PB
		EB
		ZB
		YB
	)
	fileSize := 4000000000.

	fmt.Printf("%v\n", KB)   // 1024
	fmt.Printf("%v\n", MG)   // 1048576

	fmt.Printf("%.2fGB\n", fileSize/GB) // 3.73GB
	// "%.2f" says format is floatting point with 2 decimal places
	// "GB" is literal GB that'll be printed

	// one of the best ways of using iota, bit shifting and constant is in determining user roles 
	// i.e setting boolean flags inside a simple byte i.e
	const (
		isAdmin = 1 << iota // 1 << 0 i.e 2^0 * 2^0 = 1
		isHeadquaters // 1 << 1 =  2 (2^0 * 2^1)
		canSeeFinancials // 1 << 2 = 4 (2^0 * 2^2)

		canSeeAfrica
		canSeeAsia
		canSeeEurope
		canSeeNorthAmerica
		canSeeSouthAmerica
	)
	// in this block, I've defined the various roles 
	// Each constant will occupy one location in a byte

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles) //100101
	// to check if a user is an admin, i can do this:
	fmt.Printf("Is Admin? %v\n", isAdmin & roles == isAdmin) // Is Admin? true
	// to check for a role a user doesnt have
	fmt.Printf("Is HQ? %v\n", isHeadquaters & roles == isHeadquaters) // Is HQ? false

	// SUMMARY: Constants are
	// Immutable, but can be shadowed in value and type
	// Replaced by compiler at compile time:
	// - value must be calculable at compile time
	// Named like variables such as: PascalCase for exports or camelCase for internal
	// Typed constants work like immutable variables:
	// - can interoperate only with the same type
	// Untyped constants work like literal:
	// - can interoperate with similar types
	// Enumerated Constants:
	// - special symbol iota allows related constants to be created easily
	// - Iota starts at 0 in each const block and increment by 1
	// Enumerated Expressions - using iota:
	// - Operations that can be determined at compile time are allowed:
	//   - Arithmetic
	//   - Bitwise operations
	//   - Bitshifting
}
