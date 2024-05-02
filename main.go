package main

import (
	m "Enervon-CProMax/models"
	p "Enervon-CProMax/procedures"
	"fmt"
)

func main() {
	var determinator, n int
	var data m.ProductData
	p.Start()
	p.Menu()
	fmt.Print("Please Enter Menu: ")
	fmt.Scan(&determinator)
	fmt.Println()
	for determinator != 5 {
		if determinator == 1 {
			p.FillArray(&data, &n)
		} else if determinator == 2 {
			p.ViewAllArray(data, n)
		} else if determinator == 3 {
			p.DetailedHighestComposition(data, n)
		} else if determinator == 4 {
			p.DetailedLowestComposition(data, n)
		}
		p.Menu()
		fmt.Print("Please Enter Menu: ")
		fmt.Scan(&determinator)
		fmt.Println()
	}
	p.End()
}
