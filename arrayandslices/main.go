package main
// We'd look into the following:
// Array:
// - creation; built in function; working with Arrays
// Slices:
// - creation; built in function; working with slices (exercises)

import "fmt"

func main() {
	grades := [3]int{93, 85, 97}
	fmt.Printf("Grades: %v\n", grades) // Grades: [93 85 97]

	// or for strings
	names := [4]string{"Jane", "John", "doe", "william"}
	fmt.Printf("Names: %v\n", names) // Names [Jane John doe william]

	// now instead of numbers you can use "." dots:
	grades2 := [...]int{93, 85, 97}
	fmt.Printf("Grades: %v\n", grades2)

	// we can create an array without specifying the fields like this:
	var students [3]string
	fmt.Printf("Students: %v\n", students) // Students: [  ]
	// to specify a value in the array
	students[0] = "James"
	fmt.Printf("Students: %v\n", students) //  Students: [James  ]

	// In situation where the length is unknown, you can get the length by doing:
	fmt.Printf("Number of Names: %v\n", len(names)) // 4

	// Another form of arrays are:
	// var identityMatrix [3][3]int = [3][3]int{ [3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}} // stores 3by3 identity matrix 
	// getting: redundant type from array, slice, or map composite literalsimplifycompositelitdefault
	// fmt.Println(identityMatrix) // [[1 0 0] [0 1 0] [0 0 1]]

	// we can also do this:
	var identityMatrix1 [3][3]int
	identityMatrix1[0] = [3]int{1, 0, 0}
	identityMatrix1[1] = [3]int{0, 1, 0}
	identityMatrix1[2] = [3]int{0, 0, 1} // stores 3by3 identity matrix 
	// getting: redundant type from array, slice, or map composite literalsimplifycompositelitdefault
	fmt.Println(identityMatrix1) // [[1 0 0] [0 1 0] [0 0 1]]

	// Note: In copying an array, go is creating a literal copy of the array i.e pointing to a different set of data e.g:
	a:= [...]int{1, 2, 3}
	b:= a
	b[1] = 5
	fmt.Println(a) // [1 2 3]
	fmt.Println(b) // [1 5 3]
	// if we want "b" to point to the same data "a" has instead of copying, we do:
	c := &a
	fmt.Println(c) // &[1 2 3]
	c[1] = 6
	fmt.Println(a) // [1 6 3] // a now has the same value
	fmt.Println(c) // &[1 6 3]

	// ----------------------------- | SLICE | -----------------------------
	// a slice is intialized the same way but as a literal by using "[]":
	d := []int{4, 5, 6,}
	fmt.Println(d) // [4 5 6]
	// It looks just like an array and anything an array can do, it can; with 1 or 2 exceptions.

	// An additional function to slice other than len()is capacity - cap()	
	fmt.Printf("Capacity: %v\n", cap(d)) // Capacity: 3
	// unlike an array; slices refer to the same data instead of creating a literal copy; eg:
	e := d
	e[1] = 8
	fmt.Println(d) // [4 8 6]
	fmt.Println(e) // [4 8 6]
	// we can see they point to the same thing

	// There are several other ways to create a slice
	f := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	g := f[:] // slice of all elements
	h := f[3:] // slice from 4th elements to the end
	i := f[:6] // slice first 6 elements
	j := f[3:6] // slice the 4th, 5th, 6th elements
	fmt.Println(f) // [1 2 3 4 5 6 7 8 9 10]
	fmt.Println(g) // [1 2 3 4 5 6 7 8 9 10]
	fmt.Println(h) // [4 5 6 7 8 9 10]
	fmt.Println(i) // [1 2 3 4 5 6]
	fmt.Println(j) // [4 5 6]

	// Another way to create a slice is using the makefn: takes 2 to 3 params (type, length, capacity)
	k := make([]int, 3)
	fmt.Println(k) // [0 0 0]
	fmt.Printf("Length: %v\n", len(k)) // Length: 3
	fmt.Printf("Capacity: %v\n", cap(k)) // Capacity: 3 

	// this means in slice you can have a higher capacity as the slice wont be the same size always:
	l := make([]int, 3, 100)
	fmt.Println(l) // [0 0 0]
	fmt.Printf("Length: %v\n", len(l)) // Length: 3
	fmt.Printf("Capacity: %v\n", cap(l)) // Capacity: 100

	// the way a slice works is we can add and remove elements from it:
	m := []int{}
	fmt.Println(m) // []
	fmt.Printf("Length: %v\n", len(m)) // Length: 0
	fmt.Printf("Capacity: %v\n", cap(m)) // Capacity: 0
	// we can add to the slice by using: append() // this function can take 2 or more elements
	m = append(m, 1)
	fmt.Println(m) // [1]
	fmt.Printf("Length: %v\n", len(m)) // Length: 1
	fmt.Printf("Capacity: %v\n", cap(m)) // Capacity: 1 // in the tutorial he got 2
	// append can take in more params and it is called "variadic function"
	m = append(m, 2, 3, 4, 5)
	fmt.Println(m) // [1 2 3 4 5]
	fmt.Printf("Length: %v\n", len(m)) // Length: 5
	fmt.Printf("Capacity: %v\n", cap(m)) // Capacity: 6 // in the tutorial he got 8

	// you can also a spread a slice like this:
	m = append(m, []int{6, 7, 8, 9}...)
	fmt.Println(m) // [1 2 3 4 5 6 7 8 9]
 
	// Another common operation in slice is "stack"
	n := []int{1, 2, 3, 4, 5}
	// shift operation: to remove an element from the end of the slice
	o := n[:len(n)-1]
	fmt.Println(n) // [1 2 3 4 5]
	fmt.Println(o) // [1 2 3 4]

	// Note that even when appending elements it slice, it affects the original slice
	// as stated earlier slice is just references and not create a copy when assigning slice to another variable. eg:
	fmt.Println(n) // [1 2 3 4 5]
	p := append(n[:2], n[3:]...)
	fmt.Println(p) // [1 2 4 5]
	fmt.Println(n) // [1 2 4 5 5] // we can see that 3 was removed from n also and 5 duplicated.

	// ----------------------------- | SUMMARY | -----------------------------
	// Array:
	// - collection of same type
	// - fixed type
	// - declared in 3 ways: 
	//    (a) grades := [3]int{93, 85, 97}
	//    (b) grades := [...]int{93, 85, 97}
	//    (c) var grades [3]int
	// - access via zero-based index i.e grades[1] == 85
	// - len function returns the size of the array
	// - copies underlying data not refrences
	// Slices:
	// - Backed by array
	// - creation style:
	//    (a) slice existing array or slice
	//    (b) literal style
	//    (c) use make() : make([]int, 10) or make([]int, 10, 100) i.e make(type, length, capacity)
	// - len function
	// - cap function
	// - append function
	// copies refer to same array
}