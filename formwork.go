package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Formwork struct {
	Type   string
	Width  int
	Height int
}

func NewFormworkList(filename string) []Formwork {
	var l []Formwork

	f, err := os.Open(filename)
	if err != nil {
		log.Panicln(err)
	}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.Split(sc.Text(), ", ")
		t := line[0]
		h, err := strconv.Atoi(line[1])
		if err != nil {
			log.Panic(err)
		}
		w, err := strconv.Atoi(line[2])
		if err != nil {
			log.Panic(err)
		}

		l = append(l, Formwork{t, h, w})
	}

	return l
}
