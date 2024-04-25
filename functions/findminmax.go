package functions

import (
	m "Enervon-CProMax/models"
)

const MAXDATA int = m.MAXDATA

/* Highest Composition Functions */
func HighestCompositionVitaminC(data m.ProductData, n int) int {
	var idxmaxVitaminC int
	for n < len(data) && n + 1 <= MAXDATA {
		if data[n].VitaminC > data[idxmaxVitaminC].VitaminC {
			idxmaxVitaminC = n
		}
		n++
	}
	return idxmaxVitaminC
}

func HighestCompositionVitaminB1(data m.ProductData, n int) int {
	var idxmaxVitaminB1 int
	for n < len(data) && n + 1 <= MAXDATA {
		if data[n].VitaminB1 > data[idxmaxVitaminB1].VitaminB1 {
			idxmaxVitaminB1 = n
		}
		n++
	}
	return idxmaxVitaminB1
}

func HighestCompositionVitaminB2(data m.ProductData, n int) int {
	var idxmaxVitaminB2 int
	for n < len(data) && n + 1 <= MAXDATA {
		if data[n].VitaminB2 > data[idxmaxVitaminB2].VitaminB2 {
			idxmaxVitaminB2 = n
		}
		n++
	}
	return idxmaxVitaminB2
}

func HighestCompositionVitaminB6(data m.ProductData, n int) int {
	var idxmaxVitaminB6 int
	for n < len(data) && n + 1 <= MAXDATA {
		if data[n].VitaminB6 > data[idxmaxVitaminB6].VitaminB6 {
			idxmaxVitaminB6 = n
		}
		n++
	}
	return idxmaxVitaminB6
}

func HighestCompositionVitaminB12(data m.ProductData, n int) int {
	var idxmaxVitaminB12 int
	for n < len(data) && n + 1 <= MAXDATA {
		if data[n].VitaminB12 > data[idxmaxVitaminB12].VitaminB12 {
			idxmaxVitaminB12 = n
		}
		n++
	}
	return idxmaxVitaminB12
}

func HighestCompositionVitaminD(data m.ProductData, n int) int {
	var idxmaxVitaminD int
	for n < len(data) && n + 1 <= MAXDATA {
		if data[n].VitaminD > data[idxmaxVitaminD].VitaminD {
			idxmaxVitaminD = n
		}
		n++
	}
	return idxmaxVitaminD
}

/* Lowest Composition Functions */
/* fix this plox */
func LowestCompositionVitaminC(data m.ProductData, n int) int {
	var idxlowVitaminC int
	for n < len(data) && n + 1 <= MAXDATA {
		if data[idxlowVitaminC].VitaminC > data[n].VitaminC {
			idxlowVitaminC = n
		}
		n++
	}
	return idxlowVitaminC
}

func LowestCompositionVitaminB1(data m.ProductData, n int) int {
	var idxlowVitaminB1 int
	for n < len(data) && n + 1 <= MAXDATA {
		if data[idxlowVitaminB1].VitaminB1 > data[n].VitaminB1 {
			idxlowVitaminB1 = n
		}
		n++
	}
	return idxlowVitaminB1
}

func LowestCompositionVitaminB2(data m.ProductData, n int) int {
	var idxlowVitaminB2 int
	for n < len(data) && n + 1 <= MAXDATA {
		if data[idxlowVitaminB2].VitaminB2 > data[n].VitaminB2 {
			idxlowVitaminB2 = n
		}
		n++
	}
	return idxlowVitaminB2
}

func LowestCompositionVitaminB6(data m.ProductData, n int) int {
	var idxlowVitaminB6 int
	for n < len(data) && n + 1 <= MAXDATA {
		if data[idxlowVitaminB6].VitaminB6 > data[n].VitaminB6 {
			idxlowVitaminB6 = n
		}
		n++
	}
	return idxlowVitaminB6
}

func LowestCompositionVitaminB12(data m.ProductData, n int) int {
	var idxlowVitaminB12 int
	for n < len(data) && n + 1 <= MAXDATA {
		if data[idxlowVitaminB12].VitaminB12 > data[n].VitaminB12 {
			idxlowVitaminB12 = n
		}
		n++
	}
	return idxlowVitaminB12
}

func LowestCompositionVitaminD(data m.ProductData, n int) int {
	var idxlowVitaminD int
	for n < len(data) && n + 1 <= MAXDATA {
		if data[idxlowVitaminD].VitaminD > data[n].VitaminD {
			idxlowVitaminD = n
		}
		n++
	}
	return idxlowVitaminD
}

