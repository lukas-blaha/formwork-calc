package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Wall struct {
	Type   string
	Width  int
	Height int
	Depth  int
}

func (w *Wall) getDimension(sc *bufio.Scanner, prop string) error {
	fmt.Printf("Wall %s (cm): ", prop)
	for sc.Scan() {
		if d, err := strconv.Atoi(sc.Text()); err != nil {
			return err
		} else {
			if prop == "width" {
				w.Width = d
			} else if prop == "height" {
				w.Height = d
			} else if prop == "depth" {
				w.Depth = d
			}
			break
		}
	}

	return nil
}

func NewWall() Wall {
	var wall Wall

	sc := bufio.NewScanner(os.Stdin)

	fmt.Print("Please select type of the wall [external/internal]: ")
	for sc.Scan() {
		if strings.ToLower(sc.Text()) != "internal" && strings.ToLower(sc.Text()) != "external" {
			fmt.Println("Incorrect option!")
			fmt.Print("Please select type of the wall [external/internal]: ")
		} else {
			wall.Type = sc.Text()
			break
		}
	}

	wall.getDimension(sc, "width")
	wall.getDimension(sc, "height")
	wall.getDimension(sc, "depth")

	return wall
}
