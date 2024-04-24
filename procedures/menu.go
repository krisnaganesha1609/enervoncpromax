package procedures

import "fmt"

func Start() {
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("                              Enervon-C Pro Max                               ")
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("          Inspired By: Enervon-C, Programming, and Find Min-Max               ")
	fmt.Println("                         Developed By: Kelompok 4                             ")
	fmt.Println(" Find And Compare The Highest/Lowest Composition In Your Consumables Products ")
	fmt.Println("------------------------------------------------------------------------------")
}

func Menu() {
	fmt.Println()
	fmt.Println("-------------------------------------")
	fmt.Println("              Main Menu              ")
	fmt.Println("-------------------------------------")
	fmt.Println("1. Add Product Data (Up To 20 Data)  ")
	fmt.Println("2. View All Product Data             ")
	fmt.Println("3. Find Highest Composition          ")
	fmt.Println("4. Find Lowest Composition           ")
	fmt.Println("5. Find Detailed Highest Composition ")
	fmt.Println("6. Find Detailed Lowest Composition  ")
	fmt.Println("7. Exit Program                      ")
	fmt.Println("-------------------------------------")
}

func OnFillingArray() {
	fmt.Println("-----------------------------------")
	fmt.Println("          What To Do Next?         ")
	fmt.Println("-----------------------------------")
	fmt.Println("1. Continue Adding More Data       ")
	fmt.Println("2. Finish and Back To Main Menu    ")
	fmt.Println("3. Edit Current Data               ")
	fmt.Println("-----------------------------------")
}

func End() {
	fmt.Println()
	fmt.Println("----------------------------------")
	fmt.Println(" Thank You For Using This Program ")
	fmt.Println("----------------------------------")
}
