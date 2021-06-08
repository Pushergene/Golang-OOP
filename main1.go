package main

import "fmt"

type Namen struct {
	vorname  string
	nachname string
	x        [5]int
	i        int
	y        int
}

func falls(Nm Namen) {
	if Nm.vorname == "Rujbin" && Nm.nachname == "Nassereslam" {
		fmt.Println(Nm.vorname, Nm.nachname)
	} else {
		fmt.Println(Nm.nachname)
	}
}
func Array(Nm Namen) {
	var i int
	for i = 0; i < 5; i++ {
		fmt.Println(Nm.x[i])
	}
}
func Variablen(Nm Namen) {
	fmt.Println(Nm.i, "+", Nm.y, "=", Nm.i+Nm.y)
}
func main() {
	var Nm Namen
	Nm.vorname = "Rujbin"
	Nm.nachname = "Nassereslam"
	Nm.x = [5]int{20, 30, 40, 50, 60}
	falls(Nm)
	Array(Nm)
	Nm.i = 60
	Nm.y = 90
	Variablen(Nm)
}
