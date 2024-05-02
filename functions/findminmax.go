package functions

import (
	m "Enervon-CProMax/models"
)

const MAXDATA int = m.MAXDATA

/* Highest Composition Functions */
func HighestCompositionVitaminC(data m.ProductData, n int) int {
	var idxmaxVitaminC int = 0
	var k int = 1
	for k < n {
		if data[idxmaxVitaminC].VitaminC < data[k].VitaminC {
			idxmaxVitaminC = k
		}
		k++
	}
	return idxmaxVitaminC
}

func HighestCompositionVitaminB1(data m.ProductData, n int) int {
	var idxmaxVitaminB1 int
	var k int = 1
	for k < n {
		if data[idxmaxVitaminB1].VitaminB1 < data[k].VitaminB1 {
			idxmaxVitaminB1 = k
		}
		k++
	}
	return idxmaxVitaminB1
}

func HighestCompositionVitaminB2(data m.ProductData, n int) int {
	var idxmaxVitaminB2 int
	var k int = 1
	for k < n {
		if data[idxmaxVitaminB2].VitaminB2 < data[k].VitaminB2 {
			idxmaxVitaminB2 = k
		}
		k++
	}
	return idxmaxVitaminB2
}

func HighestCompositionVitaminB6(data m.ProductData, n int) int {
	var idxmaxVitaminB6 int
	var k int = 1
	for k < n {
		if data[idxmaxVitaminB6].VitaminB6 < data[k].VitaminB6 {
			idxmaxVitaminB6 = k
		}
		k++
	}
	return idxmaxVitaminB6
}

func HighestCompositionVitaminB12(data m.ProductData, n int) int {
	var idxmaxVitaminB12 int
	var k int = 1
	for k < n {
		if data[idxmaxVitaminB12].VitaminB12 < data[k].VitaminB12 {
			idxmaxVitaminB12 = k
		}
		k++
	}
	return idxmaxVitaminB12
}

func HighestCompositionVitaminD(data m.ProductData, n int) int {
	var idxmaxVitaminD int
	var k int = 1
	for k < n {
		if data[idxmaxVitaminD].VitaminD < data[k].VitaminD {
			idxmaxVitaminD = k
		}
		k++
	}
	return idxmaxVitaminD
}

/* Lowest Composition Functions */
func LowestCompositionVitaminC(data m.ProductData, n int) int {
	var idxlowVitaminC int = 0
	var k int = 1
	for k < n {
		if data[idxlowVitaminC].VitaminC > data[k].VitaminC {
			idxlowVitaminC = k
		}
		k++
	}
	return idxlowVitaminC
}

func LowestCompositionVitaminB1(data m.ProductData, n int) int {
	var idxlowVitaminB1 int
	var k int = 1
	for k < n {
		if data[idxlowVitaminB1].VitaminB1 > data[k].VitaminB1 {
			idxlowVitaminB1 = k
		}
		k++
	}
	return idxlowVitaminB1
}

func LowestCompositionVitaminB2(data m.ProductData, n int) int {
	var idxlowVitaminB2 int
	var k int = 1
	for k < n {
		if data[idxlowVitaminB2].VitaminB2 > data[k].VitaminB2 {
			idxlowVitaminB2 = k
		}
		k++
	}
	return idxlowVitaminB2
}

func LowestCompositionVitaminB6(data m.ProductData, n int) int {
	var idxlowVitaminB6 int
	var k int = 1
	for k < n {
		if data[idxlowVitaminB6].VitaminB6 > data[k].VitaminB6 {
			idxlowVitaminB6 = k
		}
		k++
	}
	return idxlowVitaminB6
}

func LowestCompositionVitaminB12(data m.ProductData, n int) int {
	var idxlowVitaminB12 int
	var k int = 1
	for k < n {
		if data[idxlowVitaminB12].VitaminB12 > data[k].VitaminB12 {
			idxlowVitaminB12 = k
		}
		k++
	}
	return idxlowVitaminB12
}

func LowestCompositionVitaminD(data m.ProductData, n int) int {
	var idxlowVitaminD int
	var k int = 1
	for k < n {
		if data[idxlowVitaminD].VitaminD > data[k].VitaminD {
			idxlowVitaminD = k
		}
		k++
	}
	return idxlowVitaminD
}
