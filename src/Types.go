package main

import "fmt"

func typesTesting() {
	n := 1 == 1
	m := 1 == 2
	var l bool
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", m, m)
	fmt.Printf("%v, %T\n", l, l)

	b := 3.14
	fmt.Printf("%v, %T\n", b, b)

	var c complex64 = complex(5, 12)
	var d complex64 = 4 + 7i
	fmt.Printf("%v, %T\n", c, c)
	fmt.Println(c + d)
	fmt.Println(c - d)
	fmt.Println(c / d)
	fmt.Println(c * d)

	s := "this is a string"
	t := []byte(s)
	fmt.Printf("%v, %T\n", string(s[2]), string(s[2]))
	fmt.Printf("%v, %T\n", t, t)

	// rune
	r := 'a'
	fmt.Printf("%v, %T\n", r, r)

}
