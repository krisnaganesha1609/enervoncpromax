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
	if len(data) == 0 || n <= 0 {
		return -1 // return -1 or handle it as per your requirement
	}

	idxlowVitaminC := 0
	minVitaminC := data[0].VitaminC

	for i := 1; i < n && i < len(data); i++ {
		if data[i].VitaminC < minVitaminC {
			idxlowVitaminC = i
			minVitaminC = data[i].VitaminC
		}
	}
	return idxlowVitaminC
}

/*
func LowestCompositionVitaminB1(data m.ProductData, n int) float64 {
	var minVitaminB1 float64 = data[0].VitaminB1
	var k int = 1
	for k < n {
		if minVitaminB1 < data[k].VitaminB1 {
			minVitaminB1 = data[k].VitaminB1
		}
		k = k + 1
	}
	return minVitaminB1
}

func LowestCompositionVitaminB2(data m.ProductData, n int) float64 {
	var minVitaminB2 float64 = data[0].VitaminB2
	var k int = 1
	for k < n {
		if minVitaminB2 < data[k].VitaminB2 {
			minVitaminB2 = data[k].VitaminB2
		}
		k = k + 1
	}
	return minVitaminB2
}

func LowestCompositionVitaminB6(data m.ProductData, n int) float64 {
	var minVitaminB6 float64 = data[0].VitaminB6
	var k int = 1
	for k < n {
		if minVitaminB6 < data[k].VitaminB6 {
			minVitaminB6 = data[k].VitaminB6
		}
		k = k + 1
	}
	return minVitaminB6
}

func LowestCompositionVitaminB12(data m.ProductData, n int) float64 {
	var minVitaminB12 float64 = data[0].VitaminB12
	var k int = 1
	for k < n {
		if minVitaminB12 < data[k].VitaminB12 {
			minVitaminB12 = data[k].VitaminB12
		}
		k = k + 1
	}
	return minVitaminB12
}

func LowestCompositionVitaminD(data m.ProductData, n int) float64 {
	var minVitaminD float64 = data[0].VitaminD
	var k int = 1
	for k < n {
		if minVitaminD < data[k].VitaminD {
			minVitaminD = data[k].VitaminD
		}
		k = k + 1
	}
	return minVitaminD
}
*/
