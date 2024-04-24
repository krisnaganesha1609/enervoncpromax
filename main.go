package main

import (
	m "Enervon-CProMax/models"
	p "Enervon-CProMax/procedures"
	"fmt"
)

func main() {
	var determinator, n int
	var data m.ProductData
	fmt.Println("Way Down We Go - KALEO") //CAMEO :P
	p.Start()
	for determinator != 7 {
		p.Menu()
		fmt.Print("Please Enter Menu: ")
		fmt.Scan(&determinator)
		if determinator == 1 {
			p.FillArray(&data, &n)
		} else if determinator == 2 {
			p.ViewAllArray()
		}
	}
}
