package main

import "fmt"

type Namen struct {
	vorname  string
	nachname string
	x        [5]int
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
func main() {
	var Nm Namen
	Nm.vorname = "Rujbin"
	Nm.nachname = "Nassereslam"
	Nm.x = [5]int{20, 30, 40, 50, 60}
	falls(Nm)
	Array(Nm)
}
