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
	fmt.Print("Please Enter Menu: ")
	fmt.Scan(&determinator)
	for determinator != 7 {

		if determinator == 1 {
			//TODO: SOMEONE PLS FIX THIS BUGG ARRGHHHHH
			idx := f.HighestCompositionVitaminC(data, n)
			fmt.Println(idx)
			fmt.Println(data[idx].ProductDescription)
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin C in: ", data[idx].VitaminC, "mg")
			}

		} else if determinator == 2 {
			if f.HighestCompositionVitaminB1(data, n) == 0 {
				fmt.Println("Data Kosong")
			} else {
				fmt.Println(f.HighestCompositionVitaminB1(data, n))
			}

		} else if determinator == 3 {
			if f.HighestCompositionVitaminB2(data, n) == 0 {
				fmt.Println("Data Kosong")
			} else {
				fmt.Println(f.HighestCompositionVitaminB2(data, n))
			}
		} else if determinator == 4 {
			if f.HighestCompositionVitaminB6(data, n) == 0 {
				fmt.Println("Data Kosong")
			} else {
				fmt.Println(f.HighestCompositionVitaminB6(data, n))
			}
		} else if determinator == 5 {
			if f.HighestCompositionVitaminB12(data, n) == 0 {
				fmt.Println("Data Kosong")
			} else {
				fmt.Println(f.HighestCompositionVitaminB12(data, n))
			}
		} else if determinator == 6 {
			if f.HighestCompositionVitaminD(data, n) == 0 {
				fmt.Println("Data Kosong")
			} else {
				fmt.Println(f.HighestCompositionVitaminD(data, n))
			}
		} else {
			fmt.Print("Incorrect Menu. Please Enter Menu: ")
			fmt.Scan(&determinator)
		}
		fmt.Print("Please Enter Menu: ")
		fmt.Scan(&determinator)
	}
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
	fmt.Print("Please Enter Menu: ")
	fmt.Scan(&determinator)
	for determinator != 7 {
		if determinator == 1 {
			if f.LowestCompositionVitaminC(data, n-1) == 0 {
				fmt.Println("Data Kosong")
			} else {
				fmt.Println(f.LowestCompositionVitaminC(data, n-1))
			}
		} else if determinator == 2 {
			if f.LowestCompositionVitaminB1(data, n-1) == 0 {
				fmt.Println("Data Kosong")
			} else {
				fmt.Println(f.LowestCompositionVitaminB1(data, n-1))
			}
		} else if determinator == 3 {
			if f.LowestCompositionVitaminB2(data, n-1) == 0 {
				fmt.Println("Data Kosong")
			} else {
				fmt.Println(f.LowestCompositionVitaminB2(data, n-1))
			}
		} else if determinator == 4 {
			if f.LowestCompositionVitaminB6(data, n-1) == 0 {
				fmt.Println("Data Kosong")
			} else {
				fmt.Println(f.LowestCompositionVitaminB6(data, n-1))
			}
		} else if determinator == 5 {
			if f.LowestCompositionVitaminB12(data, n-1) == 0 {
				fmt.Println("Data Kosong")
			} else {
				fmt.Println(f.LowestCompositionVitaminB12(data, n-1))
			}
		} else if determinator == 6 {
			if f.LowestCompositionVitaminD(data, n-1) == 0 {
				fmt.Println("Data Kosong")
			} else {
				fmt.Println(f.LowestCompositionVitaminD(data, n-1))
			}
		} else {
			fmt.Print("Incorrect Menu. Please Enter Menu: ")
			fmt.Scan(&determinator)
		}
		fmt.Print("Please Enter Menu: ")
		fmt.Scan(&determinator)
	}
}
