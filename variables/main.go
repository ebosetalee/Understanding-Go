package main

import (
	"fmt"
	"strconv"
)

var x string = "statename"

func main() {
	var begin string = "Emmanuella" // variable initialization

	// we can do this instead
	name := "Emmanuella"

	// or

	var newName string

	fmt.Println("Hello, World from", begin)
	// Hello, World from Emmanuella (notice the spacing after from)

	// initialize here
	newName = "Ella"

	fmt.Println(name)

	fmt.Println(newName)

	// working with print function (above is print line)
	i := 99

	fmt.Printf("%v, %T", i, i) // %v - value; %T - variable type; variable
	// 99, int%   (wondering why there's a percentage at the back)

	// to convert i to float you can either name the type as float64 or float32 or do this
	k := 34.
	fmt.Printf("%v, %T", k, k)
	// 99, int34, float64%  (it returned this; i'm guessing it has to do with functions)

	// lets try
	var j float32 = 27

	fmt.Printf("%v, %T", j, j)
	// 99, int34, float6427, float32% (i get it now, it's printing without formating to the next line)
	// need to learn to move prints to next line

	// we can declare multiple variables at once by doing:
	var (
		street string  = "Johnson"
		number int     = 4
		season float32 = 23.2
	)

	fmt.Println("") // to seperate print lines
	fmt.Println(number, street, season) // 4 Johnson 23.2

	// lets look at redeclaring a variable
	// i have declared number to redecalre "number" with another value i should do: 

	number = 24

	fmt.Println(number, street, season) // 24 Johnson 23.2

	// i can't do this i'll get an error:

	// var street = "Stables"
	// fmt.Println(number, street, season) 

	//# command-line-arguments
	//variables/main.go:61:6: street redeclared in this block
	//variables/main.go:45:3: other declaration of street

	// nor this
	// street := "Studge"
	// fmt.Println(number, street, season)

	// # command-line-arguments
	// variables/main.go:66:9: no new variables on left side of :=

	// if i declare the variable in the package level like x above i can reassign x here

	fmt.Println(x)// "statename"

	var x int = 16

	fmt.Println(x) // 16

	// casing is important in go (naming conventions)
	// lowercase limits the variable to the package it was declared in
	// uppercase means it can be used outside of the package it was declared in

	// if the variable has a long term use, it's best to write a clear name.
	// working with acronym : eg
	var theURL string = "google.com"
	fmt.Println(theURL) // google.com

	// It is go practice to use uppercase for the acronym "URL", "HTTP"
	 

	// NOW:
	// Variable Conversion: Looking at the two variables below;
	var l int = 42
	fmt.Printf("%v, %T\n", l, l) // 42, int

	var m float32
	fmt.Println() // 

	m = float32(l)
	fmt.Printf("%v, %T\n", m, m) // 42, float32
	// m = float32(l) is a conversion operator
	// note the "\n" is for new line
	// however, doing in vice versa we would lose information eg:

	var n float32 = 42.5
	fmt.Printf("%v, %T\n", n, n) // 42.5, float32

	var o int
	fmt.Println() // 

	o = int(n)
	fmt.Printf("%v, %T\n", o, o) // 42, int

	// as 42.5 isn't 42.

	// now switching from int to string will produce a weird result

	var p int = 42
	fmt.Printf("%v, %T\n", p, p) // 42, int

	var q string
	fmt.Println() // 

	// q = string(p) // we already get a line error saying "cannot convert p (variable of type float32) to type stringcompiler InvalidConversion"
	// but if we run it in the go url tour we get *, string

	// instead we use the "strconv" i.e string conversion package
	q = strconv.Itoa(p)
	fmt.Printf("%v, %T\n", q, q) // 42, string

	// SUMMARY
	// 3 ways to declare a variable

	// can redeclare variables but can shadow them (i.e change values)

	// All variables must be used

	// Visibility:
	// - lowercase first letter for package scope
	// - uppercase first letter is exported globally
	// - no private scope

	// Naming Convensions: Pascal case and camelCase (no hypen, no underscore). The only exception is acronyms

	// As short as possible: longer names for longer lives

	// Type Conversion:
	// - destinationType(variable) i.e float32(32)
	// - use strconv package for strings
}