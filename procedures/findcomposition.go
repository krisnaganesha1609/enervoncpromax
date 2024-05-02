package procedures

import (
	f "Enervon-CProMax/functions"
	m "Enervon-CProMax/models"
	"fmt"
)

func DetailedHighestComposition(data m.ProductData, n int) {
	var determinator int
	MenuDetailedHighestComposition()
	fmt.Print("Please Enter Menu: ")
	fmt.Scan(&determinator)
	for determinator != 7 {
		if determinator == 1 {
			idx := f.HighestCompositionVitaminC(data, n)
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin C: ", data[idx].VitaminC, "mg")
			}
		} else if determinator == 2 {
			idx := f.HighestCompositionVitaminB1(data, n)
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin B1: ", data[idx].VitaminB1, "mg")
			}
		} else if determinator == 3 {
			idx := f.HighestCompositionVitaminB2(data, n)
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin B2: ", data[idx].VitaminB2, "mg")
			}
		} else if determinator == 4 {
			idx := f.HighestCompositionVitaminB6(data, n)
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin B6: ", data[idx].VitaminB6, "mg")
			}
		} else if determinator == 5 {
			idx := f.HighestCompositionVitaminB12(data, n)
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin B12: ", data[idx].VitaminB12, "mg")
			}
		} else if determinator == 6 {
			idx := f.HighestCompositionVitaminD(data, n)
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin D: ", data[idx].VitaminD, "mg")
			}
		} else {
			fmt.Print("Incorrect Menu. Please Enter Menu: ")
			fmt.Scan(&determinator)
		}
		MenuDetailedHighestComposition()
		fmt.Print("Please Enter Menu: ")
		fmt.Scan(&determinator)
	}
}

func DetailedLowestComposition(data m.ProductData, n int) {
	var determinator int
	MenuDetailedLowestComposition()
	fmt.Print("Please Enter Menu: ")
	fmt.Scan(&determinator)
	for determinator != 7 {
		if determinator == 1 {
			idx := f.LowestCompositionVitaminC(data, n)
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin C: ", data[idx].VitaminC, "mg")
			}
		} else if determinator == 2 {
			idx := f.LowestCompositionVitaminB1(data, n)
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin B1: ", data[idx].VitaminB1, "mg")
			}
		} else if determinator == 3 {
			idx := f.LowestCompositionVitaminB2(data, n)
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin B2: ", data[idx].VitaminB2, "mg")
			}
		} else if determinator == 4 {
			idx := f.LowestCompositionVitaminB6(data, n)
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin B6: ", data[idx].VitaminB6, "mg")
			}
		} else if determinator == 5 {
			idx := f.LowestCompositionVitaminB12(data, n)
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin B12: ", data[idx].VitaminB12, "mg")
			}
		} else if determinator == 6 {
			idx := f.LowestCompositionVitaminD(data, n)
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin D: ", data[idx].VitaminD, "mg")
			}
		} else {
			fmt.Print("Incorrect Menu. Please Enter Menu: ")
			fmt.Scan(&determinator)
		}
		MenuDetailedLowestComposition()
		fmt.Print("Please Enter Menu: ")
		fmt.Scan(&determinator)
	}
}
