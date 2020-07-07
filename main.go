package main

import (
	"fmt"
)

type Golang struct {
	golang string
	cpp string
}

func main() {
	var G Golang
	G.golang = "Schnell"
	G.cpp = "Schneller"
	Ausgabe(G)
}

func Ausgabe (b Golang) {
	fmt.Printf("Golang ist: %s\n", b.golang)
	fmt.Printf("C++ ist: %s\n", b.cpp)
}

