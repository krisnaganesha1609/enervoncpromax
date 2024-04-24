package procedures

import (
	f "Enervon-CProMax/functions"
	m "Enervon-CProMax/models"
	"fmt"
)

func DetailedHighestComposition() {
	fmt.Println()
	fmt.Println("-------------------------------------------------")
	fmt.Println("              High Composition Menu              ")
	fmt.Println("-------------------------------------------------")
	fmt.Println("1. Find detailed Highest Vitamin C Composition   ")
	fmt.Println("2. Find detailed Highest Vitamin B1 Composition  ")
	fmt.Println("3. Find detailed Highest Vitamin B2 Composition  ")
	fmt.Println("4. Find detailed Highest Vitamin B6 Composition  ")
	fmt.Println("5. Find detailed Highest Vitamin B12 Composition ")
	fmt.Println("6. Find detailed Highest Vitamin D Composition   ")
	fmt.Println("7. Go back to main menu						  ")
	fmt.Println("-------------------------------------------------")
	var determinator, n int
	var data m.ProductData
	for determinator != 7 && determinator >= 7 && determinator <= 0 {
		fmt.Print("Please Enter Menu: ")
		fmt.Scan(&determinator)
		if determinator == 1 {
			f.HighestCompositionVitaminC(data, n)
		} else if determinator == 2 {
			f.HighestCompositionVitaminB1(data, n)
		} else if determinator == 3 {
			f.HighestCompositionVitaminB2(data, n)
		} else if determinator == 4 {
			f.HighestCompositionVitaminB6(data, n)
		} else if determinator == 5 {
			f.HighestCompositionVitaminB12(data, n)
		} else if determinator == 6 {
			f.HighestCompositionVitaminD(data, n)
		}
	}
	Menu()
}

func DetailedLowestComposition() {
	fmt.Println()
	fmt.Println("-------------------------------------------------")
	fmt.Println("               Low Composition Menu              ")
	fmt.Println("-------------------------------------------------")
	fmt.Println("1. Find detailed Highest Vitamin C Composition   ")
	fmt.Println("2. Find detailed Highest Vitamin B1 Composition  ")
	fmt.Println("3. Find detailed Highest Vitamin B2 Composition  ")
	fmt.Println("4. Find detailed Highest Vitamin B6 Composition  ")
	fmt.Println("5. Find detailed Highest Vitamin B12 Composition ")
	fmt.Println("6. Find detailed Highest Vitamin D Composition   ")
	fmt.Println("7. Go back to main menu						  ")
	fmt.Println("-------------------------------------------------")
	var determinator, n int
	var data m.ProductData
	for determinator != 7 && determinator >= 7 && determinator <= 0 {
		fmt.Print("Please Enter Menu: ")
		fmt.Scan(&determinator)
		if determinator == 1 {
			f.LowestCompositionVitaminC(data, n)
		} else if determinator == 2 {
			f.LowestCompositionVitaminB1(data, n)
		} else if determinator == 3 {
			f.LowestCompositionVitaminB2(data, n)
		} else if determinator == 4 {
			f.LowestCompositionVitaminB6(data, n)
		} else if determinator == 5 {
			f.LowestCompositionVitaminB12(data, n)
		} else if determinator == 6 {
			f.LowestCompositionVitaminD(data, n)
		}
	}
	Menu()
}
