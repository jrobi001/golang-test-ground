package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByLast []user

func (b ByLast) Len() int           { return len(b) }
func (b ByLast) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByLast) Less(i, j int) bool { return b[i].Last < b[j].Last }

type UserByAge []user

func (b UserByAge) Len() int           { return len(b) }
func (b UserByAge) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b UserByAge) Less(i, j int) bool { return b[i].Age < b[j].Age }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	// fmt.Println(users)

	// your code goes here
	sort.Sort(UserByAge(users))
	for _, v := range users {
		fmt.Println(v.Age)
	}
	sort.Sort(ByLast(users))
	for _, v := range users {
		fmt.Println(v.Last)
	}
	// fmt.Println(users)

	for _, user := range users {
		sort.Strings(user.Sayings)
		fmt.Printf("%s %s's Favorite sayings:\n", user.First, user.Last)
		for _, saying := range user.Sayings {
			fmt.Printf("\t %v\n", saying)
		}
	}
}
