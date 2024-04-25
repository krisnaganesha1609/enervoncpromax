package procedures

import (
	f "Enervon-CProMax/functions"
	m "Enervon-CProMax/models"
	"fmt"
)

func DetailedHighestComposition(data m.ProductData) {
	var determinator, n int
	n = 0
	MenuDetailedHighestComposition()
	fmt.Print("Please Enter Menu: ")
	fmt.Scan(&determinator)
	// fmt.Print("CProMaxDebug: data[determinator-1].ProductDescription value @ procedures/composition.go before entering for determinator: ", data[determinator-1].ProductDescription, "\n")
	for determinator != 7 {
		if determinator == 1 {
			// fmt.Print("CProMaxDebug: determinator, n value @ procedures/composition.go: ", determinator, n, "\n")
			idx := f.HighestCompositionVitaminC(data, n)
			// fmt.Print("CProMaxDebug: idx value @ procedures/composition.go: ", idx, "\n")
			// fmt.Print("CProMaxDebug: data[idx].ProductDescription value @ procedures/composition.go: ", data[idx].ProductDescription, "\n")
			// fmt.Print("CProMaxDebug: data[idx+1].ProductDescription value @ procedures/composition.go: ", data[idx+1].ProductDescription, "\n")
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin C: ", data[idx].VitaminC, "mg")
			}
		} else if determinator == 2 {
			//fmt.Print("CProMaxDebug: determinator, n value @ procedures/composition.go: ", determinator, n, "\n")
			idx := f.HighestCompositionVitaminB1(data, n)
			// fmt.Print("CProMaxDebug: idx value @ procedures/composition.go: ", idx, "\n")
			// fmt.Print("CProMaxDebug: data[idx].ProductDescription value @ procedures/composition.go: ", data[idx].ProductDescription, "\n")
			// fmt.Print("CProMaxDebug: data[idx+1].ProductDescription value @ procedures/composition.go: ", data[idx+1].ProductDescription, "\n")
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin B1: ", data[idx].VitaminB1, "mg")
			}
		} else if determinator == 3 {
			// fmt.Print("CProMaxDebug: determinator, n value @ procedures/composition.go: ", determinator, n, "\n")
			idx := f.HighestCompositionVitaminB2(data, n)
			// fmt.Print("CProMaxDebug: idx value @ procedures/composition.go: ", idx, "\n")
			// fmt.Print("CProMaxDebug: data[idx].ProductDescription value @ procedures/composition.go: ", data[idx].ProductDescription, "\n")
			// fmt.Print("CProMaxDebug: data[idx+1].ProductDescription value @ procedures/composition.go: ", data[idx+1].ProductDescription, "\n")
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin B2: ", data[idx].VitaminB2, "mg")
			}
		} else if determinator == 4 {
			// fmt.Print("CProMaxDebug: determinator, n value @ procedures/composition.go: ", determinator, n, "\n")
			idx := f.HighestCompositionVitaminB6(data, n)
			// fmt.Print("CProMaxDebug: idx value @ procedures/composition.go: ", idx, "\n")
			// fmt.Print("CProMaxDebug: data[idx].ProductDescription value @ procedures/composition.go: ", data[idx].ProductDescription, "\n")
			// fmt.Print("CProMaxDebug: data[idx+1].ProductDescription value @ procedures/composition.go: ", data[idx+1].ProductDescription, "\n")
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin B6: ", data[idx].VitaminB6, "mg")
			}
		} else if determinator == 5 {
			//fmt.Print("CProMaxDebug: determinator, n value @ procedures/composition.go: ", determinator, n, "\n")
			idx := f.HighestCompositionVitaminB12(data, n)
			// fmt.Print("CProMaxDebug: idx value @ procedures/composition.go: ", idx, "\n")
			// fmt.Print("CProMaxDebug: data[idx].ProductDescription value @ procedures/composition.go: ", data[idx].ProductDescription, "\n")
			// fmt.Print("CProMaxDebug: data[idx+1].ProductDescription value @ procedures/composition.go: ", data[idx+1].ProductDescription, "\n")
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin B12: ", data[idx].VitaminB12, "mg")
			}
		} else if determinator == 6 {
			// fmt.Print("CProMaxDebug: determinator, n value @ procedures/composition.go: ", determinator, n, "\n")
			idx := f.HighestCompositionVitaminD(data, n)
			// fmt.Print("CProMaxDebug: idx value @ procedures/composition.go: ", idx, "\n")
			// fmt.Print("CProMaxDebug: data[idx].ProductDescription value @ procedures/composition.go: ", data[idx].ProductDescription, "\n")
			// fmt.Print("CProMaxDebug: data[idx+1].ProductDescription value @ procedures/composition.go: ", data[idx+1].ProductDescription, "\n")
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

func DetailedLowestComposition(data m.ProductData) {
	var determinator, n int
	n = 0
	MenuDetailedLowestComposition()
	fmt.Print("Please Enter Menu: ")
	fmt.Scan(&determinator)
	// fmt.Print("CProMaxDebug: data[determinator-1].ProductDescription value @ procedures/composition.go before entering for determinator: ", data[determinator-1].ProductDescription, "\n")
	for determinator != 7 {
		if determinator == 1 {
			// fmt.Print("CProMaxDebug: determinator, n value @ procedures/composition.go: ", determinator, n, "\n")
			idx := f.LowestCompositionVitaminC(data, n)
			// fmt.Print("CProMaxDebug: idx value @ procedures/composition.go: ", idx, "\n")
			// fmt.Print("CProMaxDebug: data[idx].ProductDescription value @ procedures/composition.go: ", data[idx].ProductDescription, "\n")
			// fmt.Print("CProMaxDebug: data[idx+1].ProductDescription value @ procedures/composition.go: ", data[idx+1].ProductDescription, "\n")
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin C: ", data[idx].VitaminC, "mg")
			}
		} else if determinator == 2 {
			// fmt.Print("CProMaxDebug: determinator, n value @ procedures/composition.go: ", determinator, n, "\n")
			idx := f.LowestCompositionVitaminB1(data, n) - 1
			// fmt.Print("CProMaxDebug: idx value @ procedures/composition.go: ", idx, "\n")
			// fmt.Print("CProMaxDebug: data[idx].ProductDescription value @ procedures/composition.go: ", data[idx].ProductDescription, "\n")
			// fmt.Print("CProMaxDebug: data[idx+1].ProductDescription value @ procedures/composition.go: ", data[idx+1].ProductDescription, "\n")
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin B1: ", data[idx].VitaminB1, "mg")
			}
		} else if determinator == 3 {
			// fmt.Print("CProMaxDebug: determinator, n value @ procedures/composition.go: ", determinator, n, "\n")
			idx := f.LowestCompositionVitaminB2(data, n) - 1
			// fmt.Print("CProMaxDebug: idx value @ procedures/composition.go: ", idx, "\n")
			// fmt.Print("CProMaxDebug: data[idx].ProductDescription value @ procedures/composition.go: ", data[idx].ProductDescription, "\n")
			// fmt.Print("CProMaxDebug: data[idx+1].ProductDescription value @ procedures/composition.go: ", data[idx+1].ProductDescription, "\n")
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin B2: ", data[idx].VitaminB2, "mg")
			}
		} else if determinator == 4 {
			// fmt.Print("CProMaxDebug: determinator, n value @ procedures/composition.go: ", determinator, n, "\n")
			idx := f.LowestCompositionVitaminB6(data, n) - 1
			// fmt.Print("CProMaxDebug: idx value @ procedures/composition.go: ", idx, "\n")
			// fmt.Print("CProMaxDebug: data[idx].ProductDescription value @ procedures/composition.go: ", data[idx].ProductDescription, "\n")
			// fmt.Print("CProMaxDebug: data[idx+1].ProductDescription value @ procedures/composition.go: ", data[idx+1].ProductDescription, "\n")
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin B6: ", data[idx].VitaminB6, "mg")
			}
		} else if determinator == 5 {
			// fmt.Print("CProMaxDebug: determinator, n value @ procedures/composition.go: ", determinator, n, "\n")
			idx := f.LowestCompositionVitaminB12(data, n) - 1
			// fmt.Print("CProMaxDebug: idx value @ procedures/composition.go: ", idx, "\n")
			// fmt.Print("CProMaxDebug: data[idx].ProductDescription value @ procedures/composition.go: ", data[idx].ProductDescription, "\n")
			// fmt.Print("CProMaxDebug: data[idx+1].ProductDescription value @ procedures/composition.go: ", data[idx+1].ProductDescription, "\n")
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin B12: ", data[idx].VitaminB12, "mg")
			}
		} else if determinator == 6 {
			// fmt.Print("CProMaxDebug: determinator, n value @ procedures/composition.go: ", determinator, n, "\n")
			idx := f.LowestCompositionVitaminD(data, n) - 1
			// fmt.Print("CProMaxDebug: idx value @ procedures/composition.go: ", idx, "\n")
			// fmt.Print("CProMaxDebug: data[idx].ProductDescription value @ procedures/composition.go: ", data[idx].ProductDescription, "\n")
			// fmt.Print("CProMaxDebug: data[idx+1].ProductDescription value @ procedures/composition.go: ", data[idx+1].ProductDescription, "\n")
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
