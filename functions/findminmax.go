package functions

import (
	m "Enervon-CProMax/models"
)

//TODO: PLS FIX THE FUNCTIONS FROM RETURNING FLOAT64 INTO RETURNING INT Index of array

func HighestCompositionVitaminC(data m.ProductData, n int) int {
	var idxmaxVitaminC int = 0
	var k int = 1
	for k < n {
		if data[idxmaxVitaminC].VitaminC < data[k].VitaminC {
			idxmaxVitaminC = k
		}
		k = k + 1
	}
	return idxmaxVitaminC
}

func HighestCompositionVitaminB1(data m.ProductData, n int) float64 {
	var maxVitaminB1 float64 = data[0].VitaminB1
	var k int = 1
	for k < n {
		if maxVitaminB1 < data[k].VitaminB1 {
			maxVitaminB1 = data[k].VitaminB1
		}
		k = k + 1
	}
	return maxVitaminB1
}

func HighestCompositionVitaminB2(data m.ProductData, n int) float64 {
	var maxVitaminB2 float64 = data[0].VitaminB2
	var k int = 1
	for k < n {
		if maxVitaminB2 < data[k].VitaminB2 {
			maxVitaminB2 = data[k].VitaminB2
		}
		k = k + 1
	}
	return maxVitaminB2
}

func HighestCompositionVitaminB6(data m.ProductData, n int) float64 {
	var maxVitaminB6 float64 = data[0].VitaminB6
	var k int = 1
	for k < n {
		if maxVitaminB6 < data[k].VitaminB6 {
			maxVitaminB6 = data[k].VitaminB6
		}
		k = k + 1
	}
	return maxVitaminB6
}

func HighestCompositionVitaminB12(data m.ProductData, n int) float64 {
	var maxVitaminB12 float64 = data[0].VitaminB12
	var k int = 1
	for k < n {
		if maxVitaminB12 < data[k].VitaminB12 {
			maxVitaminB12 = data[k].VitaminB12
		}
		k = k + 1
	}
	return maxVitaminB12
}

func HighestCompositionVitaminD(data m.ProductData, n int) float64 {
	var maxVitaminD float64 = data[0].VitaminD
	var k int = 1
	for k < n {
		if maxVitaminD < data[k].VitaminD {
			maxVitaminD = data[k].VitaminD
		}
		k = k + 1
	}
	return maxVitaminD
}

func LowestCompositionVitaminC(data m.ProductData, n int) float64 {
	var minVitaminC float64 = data[0].VitaminC
	var k int = 1
	for k < n {
		if minVitaminC < data[k].VitaminC {
			minVitaminC = data[k].VitaminC
		}
		k = k + 1
	}
	return minVitaminC
}

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
