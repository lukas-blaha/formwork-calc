package main

import "fmt"

func main() {
	fList := NewFormworkList("formwork.csv")

	// w := NewWall()
	w := Wall{
		Type:   "external",
		Width:  1100,
		Height: 305,
		Depth:  25,
	}

	fmt.Println(fList)
	fmt.Println(w)
}
