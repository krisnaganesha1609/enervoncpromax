package procedures

import (
	m "Enervon-CProMax/models"
	"fmt"
)

const MAXDATA int = m.MAXDATA

func FillArray(data *m.ProductData, n *int) {
	var determinator int
	OnFillingArray()
	fmt.Print("Please Enter Menu: ")
	fmt.Scan(&determinator)
	for determinator != 2 {
		if determinator == 1 {
			fmt.Println("Fill The Data (every data is nullable):")
			fmt.Print("    Product name: ")
			fmt.Scan(&data[*n].ProductDescription)
			for data[*n].ProductDescription == "" {
				fmt.Print("Product name empty! Please enter product name: ")
				fmt.Scan(&data[*n].ProductDescription)
			}
			fmt.Print("    Amount of Vitamin C Content (mg): ")
			fmt.Scan(&data[*n].VitaminC)
			if data[*n].VitaminC == 0 {
				data[*n].VitaminC = 0
			}
			fmt.Print("    Amount of Vitamin B1 Content (mg): ")
			fmt.Scan(&data[*n].VitaminB1)
			if data[*n].VitaminB1 == 0 {
				data[*n].VitaminB1 = 0
			}
			fmt.Print("    Amount of Vitamin B2 Content (mg): ")
			fmt.Scan(&data[*n].VitaminB2)
			if data[*n].VitaminB2 == 0 {
				data[*n].VitaminB2 = 0
			}
			fmt.Print("    Amount of Vitamin B6 Content (mg): ")
			fmt.Scan(&data[*n].VitaminB6)
			if data[*n].VitaminB6 == 0 {
				data[*n].VitaminB6 = 0
			}
			fmt.Print("    Amount of Vitamin B12 Content (mg): ")
			fmt.Scan(&data[*n].VitaminB12)
			if data[*n].VitaminB12 == 0 {
				data[*n].VitaminB12 = 0
			}
			fmt.Print("    Amount of Vitamin D Content (mg): ")
			fmt.Scan(&data[*n].VitaminD)
			if data[*n].VitaminD == 0 {
				data[*n].VitaminD = 0
			}
			// TODO: Tambahin lagi fmt Scanln buat struct field yang lainnya
			// Done lol
			fmt.Println("Data index [", *n, "] has been filled.")
			*n++
			if *n >= MAXDATA {
				*n = MAXDATA
			}
		}
		OnFillingArray()
		fmt.Print("Please Enter Menu: ")
		fmt.Scan(&determinator)
	}

}

func ViewAllArray(data m.ProductData, n int) {
	if n == 0 {
		fmt.Println("Data is empty")
	} else {
		for i := 0; i < n; i++ {
			fmt.Print("Data index [", i, "] :", data[i], "\n")
		}
	}
}
