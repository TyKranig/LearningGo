package main

import (
	"fmt"
)

const a int16 = 27
const (
	w = iota
	x
	z
)

const (
	w2 = iota
)

const (
	// dont care about 0 :(
	_ = iota + 5
	cat
	dog
	snake
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

// defaults set to 1 to check against
const (
	admin = 1 << iota
	hq
	financials
	africa
	asia
	europe
	na
	sa
)

func constants() {
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)
	const a int = 12
	fmt.Printf("%v, %T\n", a, a)

	var b int = 27
	fmt.Printf("%v, %T\n", a+b, a+b)

	// works because t is replaced with 42 in compiling
	const t = 42
	var u int16 = 27
	// compiler sees      42+u, 42+u
	fmt.Printf("%v, %T\n", t+u, t+u)

	fmt.Printf("%v, %T\n", w, w)
	fmt.Println(w)
	fmt.Println(x)
	fmt.Println(z)
	fmt.Println(w2)

	var typeAnimal int = dog
	fmt.Printf("%v\n", typeAnimal)

	fileSize := 4000000000.

	fmt.Printf("%.2fMB\n", fileSize/MB)

	// the roles for this "employee"
	var roles byte = admin | financials | europe
	fmt.Printf("Is admin? %v\n", admin&roles == admin)
	fmt.Printf("Is at the hq? %v\n", hq&roles == admin)
}
