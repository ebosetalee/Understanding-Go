package main

// We'd look into the following:
// Maps:
// - What they are
// - creation
// - manipulation
// Structs:
// - What they are
// - creation
// - Naming Convention
// - Embedding
// - Tags

import "fmt"

func main() {
	statePopulation := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	// this is map["key type"]"value type"
	fmt.Println((statePopulation)) // map[California:39250017 Florida:20612439 Illinois:12801539 New York:19745289 Ohio:11614373 Pennsylvania:12802503 Texas:27862596]
	// not a great output
	//you cant use slices, map or other functions as key or value type. eg:
	// a := map[[]int]string{} // we get invalid map key type []int (compiler) IncomparableMapKey
	// however if we have this
	a := map[[3]int]string{} // it works because an array is a valid key type but not a slice
	fmt.Println((a))         // map[] // empty cause no fields in it

	// Another way to create a map is using make()
	// b := make(map[string]int) // I'm getting this "value of b is never used (SA4006) go-staticcheck"
	b := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	fmt.Println(b) // map[California:39250017 Florida:20612439 Illinois:12801539 New York:19745289 Ohio:11614373 Pennsylvania:12802503 Texas:27862596]
	// In other languages, map works like objects in adding and getting key values
	statePopulation["Georgia"] = 10310371
	fmt.Println((statePopulation)) // map[California:39250017 Florida:20612439 Georgia:10310371 Illinois:12801539 New York:19745289 Ohio:11614373 Pennsylvania:12802503 Texas:27862596]

	// Note: the return order of a map is not guaranteed. (not sure about this but once there's a change yes.)
	// we can delete using: delete()
	delete(statePopulation, "Georgia")
	fmt.Println(statePopulation) // map[California:39250017 Florida:20612439 Illinois:12801539 New York:19745289 Ohio:11614373 Pennsylvania:12802503 Texas:27862596]
	// however, the delete function will return a value:
	fmt.Println((statePopulation["Georgia"])) // 0
	// now look at this
	fmt.Println((statePopulation["Oho"])) // 0
	// it returned 0 same as Georgia, whether misspelt or doesnt exist, why? till we know.
	// Another way to verify the details in our map we can use the comma ok syntax
	// what we've been doing so far is:
	pop := statePopulation["Oho"]
	fmt.Println((pop)) // 0
	// but if we do:
	c, ok := statePopulation["Oho"]
	fmt.Println(c, ok) // 0 false
	// The false states if it exists or not
	// if we ask for an existing key we get true:
	d, ok := statePopulation["Ohio"]
	fmt.Println(d, ok) // 11614373 true
	// if you just want to check if it exists without the value you can do:
	_, available := statePopulation["Ohio"]
	fmt.Println(available) // true

	// you can also use len here
	// Just like slice, a map references and not copy so change to a is change to b

	// -------------------- | STRUCTS | ----------------------

	type doctor struct {
		number     int
		actorName  string
		companions []string
	}

	aDoctor := doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoctor) // {3 Jon Pertwee [Liz Shaw Jo grant Sarah Jane Smith]}
	// in a struct, we can mix any type of data and not restricted to consistent types
	// we can get specific values using the dot syntax
	fmt.Println(aDoctor.actorName) // Jon Pertwee
	// we can also do this to the slice
	fmt.Println(aDoctor.companions[2]) // Sarah Jane Smith

	// you can also create another without using the syntax
	bDoctor := doctor{
		3,
		"Jon Pertwee",
		[]string{
			"Liz Shaw",
			"Jo grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(bDoctor) // {3 Jon Pertwee [Liz Shaw Jo grant Sarah Jane Smith]}
	// best to use the first and avoid the second in case of adding new fields as postion syntax won't work.

	// Another way to declare struct is: (literal syntax)
	cDoctor := struct{name string}{name: "Jon Pertwee"}
	fmt.Println(cDoctor) // {Jon Pertwee}

	// Unlike map and slices, struct copies the data not referencing it. e.g:
	dDoctor := cDoctor
	dDoctor.name = "Liz Shaw"
	fmt.Println(dDoctor) // {Liz Shaw}
	fmt.Println(cDoctor) // {Jon Pertwee}

	// Just like arrays, to use referencing instead of copying you use &
	eDoctor := &cDoctor
	eDoctor.name = "Tom Baker"
	fmt.Println(cDoctor) // {Tom Baker}
	fmt.Println(eDoctor) // &{Tom Baker}

	// -------------------- | EMBEDDING | ----------------------
	// NB: Go doesnt have inheritance but has composition; which is similar to inheritance: eg:

	type animal struct {
		name string
		origin string
	}

	type bird struct {
		animal // this is called compostion through embedding.
		speedKPH float32
		canFly bool
	}
	//
	f := bird{}
	f.name = "Emu"
	f.origin = "Australia"
	f.speedKPH = 48
	f.canFly = false
	fmt.Println(f) // {{Emu Australia} 48 false}
	// i can do this and it'll work
	fmt.Println(f.name) // Emu

	// if we want to declare bird with literal syntax then we have to explictly talk about the internal animal structure:
	g := bird{
		animal: animal{name: "Emu", origin: "Australia"},
		speedKPH: 48,
		canFly: false,
	}
	fmt.Println(g) // {{Emu Australia} 48 false}
	// i can do this and it'll work
	fmt.Println(g.name) // Emu
	// -------------------- | TAGS | ----------------------
	// The format of tag is to have backticks "`" as the delimiters. eg:
	type person struct {
		name string `required, max: "100"` // tags
		origin string
	}

	// you need to use reflect

	// marshal the struct to bytes

	// -------------------- | SUMMARY | ----------------------
}
