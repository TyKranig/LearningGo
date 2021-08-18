package main

import (
	"github.com/tylerkranig/life"
)

// cannot use colon syntax here
// vars can be declared in packages
var (
	actorName string = "Elisabeth Sladen"
	companion string = "sarah jane smith"
)

// lower case vars scoped to the package
var i int = 42

// upper case vars scoped global
var I int = 42

// length of var name reflects the life of the variable
var seasonNumber int = 11

// package level long and verbose
var versionOfTheThing int = 1

// block scope
func testMain() {
	// go figures out what the variable type is
	// inner most scope is used here, called shadowing
	// fmt.Printf("%v, %T\n", i, i)
	// var i int = 42
	// var j float32 = 27
	// k := 99.0
	// fmt.Printf("%v, %T\n", i, i)

	// var m int = 42
	// fmt.Println(m)
	// var j float32
	// j = float32(m)
	// fmt.Println(j)

	// var k string
	// k = strconv.Itoa(m)
	// fmt.Println(k + " k is this")

	// maps()
}

func main() {
	life.Run(400, 300)
}
