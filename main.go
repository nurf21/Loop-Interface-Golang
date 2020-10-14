package main

import (
	"fmt"
)

type print interface {
	rataKiri()
	rataKanan()
}

type param struct {
	value int
}

func main() {
	var num int
	fmt.Print("Masukkan angka : ")
	fmt.Scanf("%d", &num)

	var pattern print
	pattern = param{num}
	pattern.rataKiri()
	fmt.Println("\n ")
	pattern.rataKanan()
}

func (c param) rataKiri() {
	for i := 1; i <= c.value; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func (c param) rataKanan() {
	for i := 1; i <= c.value; i++ {
		for j := 1; j <= i; j++ {
			if j > 1 {
				fmt.Print(" ")
			}
		}
		for k := i; k <= c.value; k++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
