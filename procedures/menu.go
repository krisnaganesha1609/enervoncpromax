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
	fmt.Println("3. Find Detailed Highest Composition ")
	fmt.Println("4. Find Detailed Lowest Composition  ")
	fmt.Println("5. Exit Program                      ")
	fmt.Println("-------------------------------------")
}

func OnFillingArray() {
	fmt.Println("-----------------------------------")
	fmt.Println("          What To Do Next?         ")
	fmt.Println("-----------------------------------")
	fmt.Println("1. Continue Adding More Data       ")
	fmt.Println("2. Finish and Back To Main Menu    ")
	fmt.Println("-----------------------------------")
}

func MenuDetailedHighestComposition() {
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
	fmt.Println("7. Go back to main menu				")
	fmt.Println("-------------------------------------------------")
}

func MenuDetailedLowestComposition() {
	fmt.Println()
	fmt.Println("-------------------------------------------------")
	fmt.Println("               Low Composition Menu              ")
	fmt.Println("-------------------------------------------------")
	fmt.Println("1. Find detailed Lowest Vitamin C Composition   ")
	fmt.Println("2. Find detailed Lowest Vitamin B1 Composition  ")
	fmt.Println("3. Find detailed Lowest Vitamin B2 Composition  ")
	fmt.Println("4. Find detailed Lowest Vitamin B6 Composition  ")
	fmt.Println("5. Find detailed Lowest Vitamin B12 Composition ")
	fmt.Println("6. Find detailed Lowest Vitamin D Composition   ")
	fmt.Println("7. Go back to main menu				 ")
	fmt.Println("-------------------------------------------------")
}

func End() {
	fmt.Println()
	fmt.Println("----------------------------------")
	fmt.Println(" Thank You For Using This Program ")
	fmt.Println("----------------------------------")
}
