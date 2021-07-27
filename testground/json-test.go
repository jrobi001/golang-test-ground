package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name   string  `json:"PersonName"`
	Age    int     `json:"PersonAge"`
	Height float64 `json:"PersonHeight"`
}

func main() {

	bs := []byte(`[{"PersonName":"Bob Mccoy","PersonAge":32,"PersonHeight":1.78},{"PersonName":"Jill Jefferson","PersonAge":65,"PersonHeight":1.58}]`)

	var people []person

	err := json.Unmarshal(bs, &people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people[0].Name)
}
