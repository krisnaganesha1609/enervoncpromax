package models

type ProductComposition struct {
<<<<<<< HEAD
	ProductID         uuid.UUID
	ProductName       string
	VitaminC          float64
	Niasinamida       float64
	KalsiumPantotenat float64
	VitaminB1         float64
	VitaminB2         float64
	VitaminB6         float64
	VitaminB12        float64
=======
	ProductDescription string
	VitaminC           float64
	VitaminB1          float64
	VitaminB2          float64
	VitaminB6          float64
	VitaminB12         float64
	VitaminD           float64
>>>>>>> 45465f4 (treewide: huge update)
	//TODO: Add More Product Composition Field E.g :(Fosfor, Vitamin D, Protein, Cholesterol, etc.)
}

const MAXDATA int = 20

type ProductData [MAXDATA]ProductComposition
