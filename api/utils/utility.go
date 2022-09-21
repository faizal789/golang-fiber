package utils


import "math"

func GetEnd(pageNumber int, size int) int {
	if pageNumber == 0 {
		return 0
	}

	return int(math.Abs(float64(((pageNumber - 1)*size + 1) - 1)))
	
}