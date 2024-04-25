package functions

import (
	m "Enervon-CProMax/models"
)

const MAXDATA int = m.MAXDATA

/* Highest Composition Functions */
func HighestCompositionVitaminC(data m.ProductData, n int) int {
	var idxmaxVitaminC int
	for n < len(data) && n+1 <= MAXDATA {
		if data[n].VitaminC > data[idxmaxVitaminC].VitaminC {
			idxmaxVitaminC = n
		}
		n++
	}
	return idxmaxVitaminC
}

func HighestCompositionVitaminB1(data m.ProductData, n int) int {
	var idxmaxVitaminB1 int
	for n < len(data) && n+1 <= MAXDATA {
		if data[n].VitaminB1 > data[idxmaxVitaminB1].VitaminB1 {
			idxmaxVitaminB1 = n
		}
		n++
	}
	return idxmaxVitaminB1
}

func HighestCompositionVitaminB2(data m.ProductData, n int) int {
	var idxmaxVitaminB2 int
	for n < len(data) && n+1 <= MAXDATA {
		if data[n].VitaminB2 > data[idxmaxVitaminB2].VitaminB2 {
			idxmaxVitaminB2 = n
		}
		n++
	}
	return idxmaxVitaminB2
}

func HighestCompositionVitaminB6(data m.ProductData, n int) int {
	var idxmaxVitaminB6 int
	for n < len(data) && n+1 <= MAXDATA {
		if data[n].VitaminB6 > data[idxmaxVitaminB6].VitaminB6 {
			idxmaxVitaminB6 = n
		}
		n++
	}
	return idxmaxVitaminB6
}

func HighestCompositionVitaminB12(data m.ProductData, n int) int {
	var idxmaxVitaminB12 int
	for n < len(data) && n+1 <= MAXDATA {
		if data[n].VitaminB12 > data[idxmaxVitaminB12].VitaminB12 {
			idxmaxVitaminB12 = n
		}
		n++
	}
	return idxmaxVitaminB12
}

func HighestCompositionVitaminD(data m.ProductData, n int) int {
	var idxmaxVitaminD int
	for n < len(data) && n+1 <= MAXDATA {
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
	var idxlowVitaminC int = 0
	var k int = 1
	// for n < len(data) && n+1 <= MAXDATA {
	// 	if data[idxlowVitaminC].VitaminC > data[n].VitaminC {
	// 		idxlowVitaminC = n
	// 	}
	// 	n = n + 1
	// }
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
	// for n < len(data) && n+1 <= MAXDATA {
	// 	if data[idxlowVitaminC].VitaminC > data[n].VitaminC {
	// 		idxlowVitaminC = n
	// 	}
	// 	n = n + 1
	// }
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
	// for n < len(data) && n+1 <= MAXDATA {
	// 	if data[idxlowVitaminC].VitaminC > data[n].VitaminC {
	// 		idxlowVitaminC = n
	// 	}
	// 	n = n + 1
	// }
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
	// for n < len(data) && n+1 <= MAXDATA {
	// 	if data[idxlowVitaminC].VitaminC > data[n].VitaminC {
	// 		idxlowVitaminC = n
	// 	}
	// 	n = n + 1
	// }
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
	// for n < len(data) && n+1 <= MAXDATA {
	// 	if data[idxlowVitaminC].VitaminC > data[n].VitaminC {
	// 		idxlowVitaminC = n
	// 	}
	// 	n = n + 1
	// }
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
	// for n < len(data) && n+1 <= MAXDATA {
	// 	if data[idxlowVitaminC].VitaminC > data[n].VitaminC {
	// 		idxlowVitaminC = n
	// 	}
	// 	n = n + 1
	// }
	for k < n {
		if data[idxlowVitaminD].VitaminD > data[k].VitaminD {
			idxlowVitaminD = k
		}
		k++
	}
	return idxlowVitaminD
}
