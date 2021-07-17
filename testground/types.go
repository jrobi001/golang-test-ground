package main

import "fmt"

func main() {
	a := 2
	b := 2.
	c := 'a'
	d := "a"
	e := '日'
	f := "elephants"
	g := '⌘'
	fmt.Printf("%T\n", a)
	fmt.Println(c)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Println(e)
	e += 24
	fmt.Println(string(e))
	c -= 32
	fmt.Println(string(c))
	fmt.Println(len(f))
	fmt.Println(g)

	fmt.Printf("%v + %v = %v\n", a, b, a+int(b))
	fmt.Printf("%c\n", 26085)
	fmt.Printf("%v\t%#x\t%b\t%x\n", e, e, e, e)
	s := fmt.Sprint("cake is good")
	fmt.Println(s)

	type counter int
	var z counter = 23
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	j := 23
	j = int(z)
	fmt.Println(j)

}
