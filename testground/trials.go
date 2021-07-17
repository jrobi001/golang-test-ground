package main

import "fmt"

func main() {
	b := 13
	b <<= b
	printNumerical(b)

}

func printNumerical(i int) {
	fmt.Printf("%v\t%b\t%#x\n", i, i, i)
}
