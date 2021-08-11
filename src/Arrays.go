package main

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"math"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/reader"
	"golang.org/x/tour/wc"
)

func arrays() {
	//declares the size of the array twice
	//grades := [3]int{97, 85, 93}
	grades := [...]int{97, 85, 93}
	fmt.Println(grades)

	//example for string array
	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("# Students: %v\n", len(students))

	//example of creating a matrix
	var matrix [3][3]int = [3][3]int{
		{1, 0, 0},
		{1, 1, 0},
		{1, 0, 1},
	}
	fmt.Println(matrix)

	//some fun with pointers
	a := [...]int{1, 2, 3}
	b := &a
	b[1] = 5
	//b points to a
	fmt.Println(a)
	fmt.Println(b)

	//slice fun
	t := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	u := t[:]
	v := t[3:]
	t[5] = 42
	fmt.Println(t)
	fmt.Println(u)
	fmt.Println(v)

	//make, if you know you need 100 elements this runs faster
	//because it never needs to reallocate
	g := make([]int, 3, 100)
	fmt.Println(g)
	fmt.Printf("Capacity: %v, Length: %v\n", len(g), cap(g))
	//spread operator similar to java script
	//capacity is the size of the underlying array
	//length is the size of the slice
	g = append(g, []int{2, 3, 4, 5}...)
	fmt.Println(g)

	h := []int{2, 3, 4, 5}
	fmt.Println(h)
	hs := h[1:]
	fmt.Println(hs)
	hp := h[:len(h)-1]
	fmt.Println(hp)
	hm := append(h[:2], h[3:]...)
	fmt.Println(hm)
	fmt.Println("complete chaos")
	fmt.Println(h)

	//creates an array [n]T n size with type T
	//cannot be resized becuase length is part of the type
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//a slice is a dynamically sized "view" into the elements of the array
	//a slice does not actually store any data, it only desribes a section of the underlying array
	//changing the elements in a slice changes the elements in the underlying array
	//as such, other slices viewing the same underlying array will see the changes
	var s []int = primes[1:4]
	fmt.Println(s)

	//a slice literal is similar to an array literal but it doesnt have length
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	type Test struct {
		i int
		b bool
	}

	//defining a struct and using it in the array
	testArr := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(testArr)

	s = []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
	//notice how you dont lose the data in the underlying
	//until you drop from the front
	//because the address of the array points to the start of the array
	//so removing from the front also shrinks capacity
	s = s[:]
	printSlice(s)

	//null value is nil
	var sli []int
	printSlice(sli)
	if sli == nil {
		fmt.Println("nil!")
	}
	//to create dynamically sized arraysuse the make function
	a2 := make([]int, 5)
	printSlicev2("a2", a2)

	//length 0 capacity 5
	b2 := make([]int, 0, 5)
	printSlicev2("b2", b2)

	c2 := b2[:2]
	printSlicev2("c2", c2)

	d2 := c2[2:5]
	printSlicev2("d2", d2)

	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		//join brings together all values in a slice with the given arg in between
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var s3 []int
	printSlice(s)

	// append works on nil slices.
	s3 = append(s3, 0)
	printSlice(s3)

	// The slice grows as needed.
	s3 = append(s3, 3, 1)
	printSlice(s3)

	// We can add more than one element at a time.
	s3 = append(s3, 2, 3, 4)
	printSlice(s3)

	//for loops work with slices using the range call
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for index, element := range pow {
		fmt.Printf("2**%d = %d\n", index, element)
	}

	pic.Show(Pic)
}

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for i := range pic {
		pic[i] = make([]uint8, dx)
		for j := range pic[i] {
			pic[i][j] = uint8((i + j) / 2)
		}
	}
	return pic
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func printSlicev2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func maps() {
	statePopulations := map[string]int{
		"Iowa": 3000000,
	}
	fmt.Println(statePopulations)

	// check for map keys
	_, isThere := statePopulations["Ohio"]
	fmt.Println(isThere)
	//return size of keys
	fmt.Println(len(statePopulations))

	//can check for existence and use the var at the same time
	if pop, ok := statePopulations["Iowa"]; ok {
		fmt.Println(pop)
	}

	if pop, ok := statePopulations["Ohio"]; ok {
		fmt.Println(pop)
	} else { //and you can use the vars created in the original if block
		fmt.Println(strconv.FormatBool(ok) + " Ohio was not found")
	}
	//but the vars created in the if block are only accessible in the block itself

	rand.Seed(6)
	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Println(split(10))

	sum := 0
	for i := 0; i < 10; i++ {
		sum = sum + i
		fmt.Println(sum)
	}

	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	type Vertex struct {
		Lat, Long float64
	}
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(m["Bell Labs"])
	m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m)

	m2 := make(map[string]int)
	m2["Answer"] = 42
	fmt.Println("The value:", m2["Answer"])
	m2["Answer"] = 48
	fmt.Println("The value:", m2["Answer"])
	delete(m2, "Answer")
	fmt.Println("The value:", m2["Answer"])
	v, ok := m2["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	wc.Test(WordCount)

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

//function adder has no args, and return type is -function with 1 int arg, and return type is int-
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func WordCount(s string) map[string]int {
	var arr = make(map[string]int)
	for _, word := range strings.Fields(s) {
		v, exists := arr[word]
		if exists {
			arr[word] = v + 1
		} else {
			arr[word] = 1
		}
	}
	return arr
}

//cannot use := here, outside a function must use a keyword (var, func, ...)
var i2, j int = 1, 2

//return values defined in the function definition
//naked return, a return statement without values returns the name values
//try to only use in short functions, harm readability
func split(sum int) (x, y int) {
	// can have multiple types in var declarations
	var c, python, java = true, false, "no!"
	fmt.Println(i2, j, c, python, java)

	x = sum + 7
	y = x + 8
	return
}

func swi() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	//checking for the day
	//time now in the playground always returns go's birthday
	//2009-11-10 23:00:00 UTC

	//A Weekday specifies a day of the week (Sunday = 0, ...).
	//defined like this inside of time.go
	//type Weekday int
	//const (
	//	Sunday Weekday = iota
	// 	Monday // 1
	// 	Tuesday // 2
	// 	Wednesday // 3
	// 	Thursday // 4
	// 	Friday // 5
	// 	Saturday // 6
	// )
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	//can also do switches without arguments
	//then proceeds to behave like a long if else
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	//this is a defer, it runs after the function has finished
	//works like a stack, so test defer gets ran first, the print world
	defer fmt.Println("world")
	defer testDefer()

	fmt.Println("hello")
}

func testDefer() {
	fmt.Println("this ran after everything")
	pointers()
}

func pointers() {
	//go has pointers but no pointer arithmetic
	i, j := 42, 64

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 8    // divide j through the pointer
	fmt.Println(j) // see the new value of j

	structs()
}

type Vertex struct {
	x float64
	y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{x: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func structs() {
	v := Vertex{1, 2}
	fmt.Println(v)
	v.x = 10
	fmt.Println(v)

	p := &v
	p.x = 1e9
	fmt.Println(p)

	fmt.Println(v1, p, v2, v3)
}

// Let's have some fun with functions.
// Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	last := 0
	current := 1
	return func() int {
		temp := current + last
		last = current
		current = temp
		return current
	}
}

func testmain2() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

type Point struct {
	x, y float64
}

//this function has a reveiver of type point
//this is a method, BECAUSE it has a receiver argument
func (v Point) DistanceFromZero() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

//same behavior as the function above
func Abs(v Point) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

//creating a type, can then becom a receiver
type MyFloat float64

//passed a reference to the original object
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//can modify values within the object
//object used loosely here because they arent objects persay but
//that gets the idea across
func (v *Point) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func typemain() {
	pt := Point{1, 2}
	fmt.Println(pt.DistanceFromZero())

	float := MyFloat(10.0)
	fmt.Println(float.Abs())

	v := Point{3, 4}
	fmt.Println(v)
	v.Scale(10)
	fmt.Println(v)
}

type Abser interface {
	Abs() float64
}

func interface1main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v

	fmt.Println(a.Abs())
}

//type I must implement the M() method
//but there is no explicit declaration of implments
//this decouples from the implementation
type Inter interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func interface2main() {
	//this lets us create object T and store it as type Inter
	//and we know that M() may be called on it
	var i Inter = T{"hello"}
	i.M()
}

type Student struct {
	Name  string
	Grade string
	Age   int
}

//implement the fmt interface needed to display as string
func (s Student) String() string {
	return fmt.Sprintf("%v (%v grade), %v", s.Name, s.Grade, s.Age)
}

func stringermain() {
	a := Student{"Tyler", "7th", 14}
	b := Student{"Ryan", "16th", 25}
	fmt.Println(a, b)
}

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func ipStringermain() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

type MyError struct {
	When time.Time
	What string
}

//comment this out and get an error below because our type
//wouldn't be implementing the Error() method
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func errormain() {
	//if err is nill then we good
	i, err := strconv.Atoi("42")
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	}
	fmt.Println("Converted integer:", i)

	if err := run(); err != nil {
		fmt.Println(err)
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot sqrt negative number %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0.0, ErrNegativeSqrt(x)
	}
	last := 0.0
	z := x / 2
	for i := 0; i < 200 && last != z; i++ {
		last = z
		z -= (z*z - x) / (2 * z)
		z = math.Floor(z*100) / 100
	}
	return z, nil
}

func sqrttestmain() {
	z, _ := Sqrt(2)
	fmt.Println(z)
	fmt.Println(Sqrt(-2))
}

type MyReader struct{}

func (mr MyReader) Read(b []byte) (int, error) {
	for index := range b {
		b[index] = 'A'
	}
	return len(b), nil
}

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	for i := 0; i < len(b); i++ {
		if b[i] >= 'A' && b[i] < 'Z' {
			b[i] = 65 + (((b[i] - 65) + 13) % 26)
		} else if b[i] >= 'a' && b[i] <= 'z' {
			b[i] = 97 + (((b[i] - 97) + 13) % 26)
		}
	}
	return n, err
}

func readermain() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
	reader.Validate(MyReader{})

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	reader := rot13Reader{s}
	io.Copy(os.Stdout, &reader)
}

type MyImage struct {
	w, h int
}

func (im MyImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, im.w, im.h)
}
func (im MyImage) ColorModel() color.Model {
	return color.RGBAModel
}
func (im MyImage) At(i, j int) color.Color {
	v := uint8((j + i) / 2)
	color := color.RGBA{v, v, 255, 255}
	return color
}

func imagetestingmain() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())

	test := MyImage{800, 800}
	pic.ShowImage(test)
}
