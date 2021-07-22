package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {

	p := struct {
		first string
		last  string
		age   int
	}{
		first: "Bob",
		last:  "Mccoy",
		age:   43,
	}

	fmt.Println(p)

	p2 := struct {
		first string
		last  string
		age   int
	}{"Bob", "Mccoy", 43}

	fmt.Println(p2)

	type employee struct {
		person
		role string
	}

	em := employee{
		person: person{"Chris", "Janson", 53},
		role:   "internal communications supervisor",
	}

	em2 := employee{person{"Jill", "Simon", 23}, "bot farmer"}
	fmt.Println(em.first)
	fmt.Println(em2.role)

	p4 := newPerson("Tim", "Beans", 23)
	// p4 := person{"Tim", "Beans", 23}
	fmt.Println(p4)

	p3 := new(person)
	p3.first = "coolio"
	p3.last = "jones"
	p3.age = 87
	fmt.Println(p3)
}

func newPerson(first string, last string, age int) *person {
	p := person{
		first: first,
		last:  last,
		age:   age,
	}
	return &p
}

func printNumerical(i int) {
	fmt.Printf("decimal:\t%v\n", i)
	fmt.Printf("binary:\t\t%b\n", i)
	fmt.Printf("hex:\t\t%#x\n", i)
}
