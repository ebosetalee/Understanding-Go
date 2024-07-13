package main

import "fmt"

// We'd look into the following:
// - Boolean types
// - Numeric Types: Integer, float and complex numbers
// - Text types

func main(){
	var a bool = true // or false
	fmt.Println(a) // true
	fmt.Printf("%v, %T\n", a, a) // true, bool

	// The usual case to use a boolean is when equating an go uses "==" too.
	b := 1 == 2-1 // i did this beacuse i was getting a warning on using same value twice using 1 == 1 (identical expressions on the left and right side of the '==' operator (SA4000))
	c := 1 == 2
	fmt.Printf("%v, %T\n", b, b) // true, bool
	fmt.Printf("%v, %T\n", c, c) // false, bool

	// Note: In go, everytime you initialize a variable it has a zero value and 0 means false so if i do this:
	var d bool
	fmt.Printf("%v, %T\n", d, d) // false, bool
	// Note: Boolean type is not an alias for other types.

	// Note: an integer bit can be specified but usually it'll be a 32 bit e.g 42, int
	// However, you can have the following int
	// - int8 = -127 to 127
	// - int16 = -32768 to 32767
	// - int32 = -2147483648 to 2147483647 
	// - int64 = -9223372036854775808 to 223372036854775807

	// we can also create a unsigned int eg:
	var e uint16 = 42 //42, uint16
	fmt.Printf("%v, %T\n", e, e)

	// You can have the following uint
	// - uint8 = 0 to 255
	// - int16 = 0 to 65535
	// - int32 = 0 to 4294967295

	f := 10
	g := 3
	fmt.Println(f + g) // 13
	fmt.Println(f - g) // 7
	fmt.Println(f * g) // 30
	fmt.Println(f / g) // 3
	fmt.Println(f % g) // 1

	// Note: - Two integers can't give a float as seen in "f / g" 
	// - you can't perform arithmetic across types eg:
	var h int = 16
	var i int8 = 8
	// fmt.Println(h + i) // invalid operation: h + i (mismatched types int and int8)compiler (MismatchedTypes)
	// for this to work you have to do a type inversion on one of the variables:
	fmt.Println(h + int(i)) // 24

	// Next is to look at bit operators. To understand the answers, you need to look at them in binary(f: 1010; g: 0011)
	fmt.Println(f & g) // 2 (and operator) //0010
	fmt.Println(f | g) // 11 (or operator) //1011
	fmt.Println(f ^ g) // 9 (exclusive or operator) //1001
	fmt.Println(f &^ g) // 8 (and not operator) //0100

	// Next is Bit shifting:
	j := 8
	fmt.Println(j << 3) // This willl bitshift "j" left 3 places while; // 64
	fmt.Println(j >> 3)// this willl bitshift "j" right 3 places. // 1
	// To understand the response:  "j" is 2^3 
	// thus j << 3 means 2^3 * 2^3 = 2^6
	// while j >> 3 means 2^3 / 2^3 = 2^0

	// Next data type is "floatting point"  follow IEEE754 standard
	// However, you can have just float 32 or 64
	// - float32 = +/-1.18E-38 to +/- 3.4E38  
	// - float64 = +/-2.23E-308 to +/- 1.8E308 
	// Here's how to create floating point number:
	k := 3.14
	k = 13.7e72 // exponential notation
	k = 2.1E14
	fmt.Printf("%v, %T\n", k, k) // 2.1e+14, float64
	// or you can do var k float32 or float64 = 3.14

	l := 10.2
	m := 3.7
	// here are the arithmetic operation that can be done on float
	fmt.Println(l + m) // 13.899999999999999
	fmt.Println(l - m) // 6.499999999999999
	fmt.Println(l * m) // 37.74
	fmt.Println(l / m) // 2.7567567567567566
	// notice there's no "%" because invalid operation: operator % not defined on l (variable of type float64)compiler (UndefinedOp)
	// There are no bit operators nor bit shifting

	// Next data type is "Complex type":
	// Two types; complex64 (float 32 + float 32) and Complex128 (float64 + float64)
	var n complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", n, n) // (1+2i), complex64
	// go understand the "i" literal as imaginary number

	// if you do this you get:
	var o complex64 = 2i
	fmt.Printf("%v, %T\n", o, o) // (0+2i), complex64; it sees i as special

	// Thew operations available are:
	p := 1 + 2i
	q := 2 + 5.2i

	fmt.Println(p + q) // (3+7.2i)
	fmt.Println(p - q) // (-1-3.2i)
	fmt.Println(p * q) // (-8.4+9.2i)
	fmt.Println(p / q) // (0.3994845360824742-0.038659793814433i)

	// to decompose the imaginary part to real part, there are two functions that can be used
	// we'd use variable n for this: 1. real fn
	fmt.Printf("%v, %T\n", real(n), real(n)) // 1, float32
	// 2. imag fn
	fmt.Printf("%v, %T\n", imag(n), imag(n)) // 2 , float32
	// The functions work in complex64 or 128. They'll pull out the real or imaginary part

	// If we run them using complex128; we get "float 64":
	var r complex128 = 1 + 2i
	fmt.Printf("%v, %T\n", real(r), real(r)) // 1, float64
	fmt.Printf("%v, %T\n", imag(r), imag(r)) // 2, float64

	// To make a complex number instead of using a literal, there's the complex fn
	var s complex128 = complex(5, 12) // this takes 2 params: 1st is real part, 2nd is imaginary part 
	fmt.Printf("%v, %T\n", s, s) // (5+12i), complex128

	// The last data type is "TEXT" type: This falls into two basic categories: 
	// 1. String: any UTF8 character:
	t := "this is a string";
	fmt.Printf("%v, %T\n", t, t) // this is a string, string
	// just like any string, you can print out a specific letter; such as:
	fmt.Printf("%v, %T\n", t[2], t[2]) // 105, uint8; this returns the bit of "i"
	// but to get the actual letter; we'd do this
	fmt.Printf("%v, %T\n", string(t[2]), t) // i, string
	// we cannot do this 
	// t[2]: "u" // illegal label declaration (syntax)
	// but we can concatenate (add two strings together)
	t2 := "trying this string too";
	fmt.Printf("%v, %T\n", t + t2, t + t2) // this is a stringtrying this string too, string

	// We can convert string to collections of bytes; which is called "slice of bytes" e.g using "t":
	u := []byte(t) // this is converting t to a collection of bytes
	fmt.Printf("%v, %T\n", u, u) // 116 104 105 115 32 105 115 32 97 32 115 116 114 105 110 103], []uint8
	// The result comes out as te ASCII / UTF value for each character in "t"
	// Nb: []unit8 is a type alias for bytes.
	// This is useful as we'd mostly work with collections of bytes in go. 

	// 2. RUNE: (different from string) it is any UTF32 character
	// in go "" (quotes) matter; while "" (double quotes) is string; rune is '' (single quotes)
	// to declare a rune, we do the following:
	v := 'a';
	fmt.Printf("%v, %T\n", v, v) // 97, int32 
	// The result is because runes are just a type alias for int32.
	// Thus, when strings can be converted back and forth between collections of bytes, 
	// runes are a true type alias i.e it is just talking about int32
	// even if you do this
	var w rune = 'a';
	fmt.Printf("%v, %T\n", w, w) // 97, int32; still get the same thing
}