package procedures

import (
	m "Enervon-CProMax/models"
	"fmt"
)

const MAXDATA int = m.MAXDATA

func FillArray(data *m.ProductData, n *int) {
	var determinator int
	for determinator != 2 {
		if determinator == 1 {
			fmt.Println("Fill The Data (every data is nullable):")
			fmt.Print("    Product Description: ")
			fmt.Scanln(&data[*n].ProductDescription)
			fmt.Print("    Amount of Vitamin C Content (mg): ")
			fmt.Scanln(&data[*n].VitaminC)
			fmt.Print("    Amount of Niasinamida Content (mg): ")
			fmt.Scanln(&data[*n].Niasinamida)
			fmt.Print("    Amount of Kalsium Pantotenat (mg): ")
			fmt.Scanln(&data[*n].KalsiumPantotenat)
			fmt.Print("    Amount of Vitamin B1 Content (mg): ")
			fmt.Scanln(&data[*n].VitaminB1)
			//TODO: Tambahin lagi fmt Scanln buat struct field yang lainnya
			fmt.Println("Data Ke-", *n+1, "Selesai Diisi")
			if *n < MAXDATA {
				*n = *n + 1
			} else if *n >= MAXDATA {
				*n = MAXDATA
			}
		}
		// else if determinator == 3 {
		// 	fmt.Println("Update The Data (if null, then the data will be set with previous one):")
		// 	fmt.Print("    Product Description: ")
		// 	fmt.Scanln(&data[*n-1].ProductDescription)
		// 	fmt.Print("    Amount of Vitamin C Content (mg): ")
		// 	fmt.Scanln(&data[*n-1].VitaminC)
		// 	fmt.Print("    Amount of Niasinamida Content (mg): ")
		// 	fmt.Scanln(&data[*n-1].Niasinamida)
		// 	fmt.Print("    Amount of Kalsium Pantotenat (mg): ")
		// 	fmt.Scanln(&data[*n-1].KalsiumPantotenat)
		// 	fmt.Print("    Amount of Vitamin B1 Content (mg): ")
		// 	fmt.Scanln(&data[*n-1].VitaminB1)
		// 	//TODO: Tambahin lagi fmt Scanln buat struct field yang lainnya
		// 	fmt.Println("Data Ke-", *n-1, "Berhasil diubah")
		// }
		OnFillingArray()
		fmt.Scan(&determinator)
	}
}

func ViewAllArray() {

}
