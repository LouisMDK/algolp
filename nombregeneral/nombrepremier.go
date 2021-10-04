package algolp

import "math"

func EstPremier(val int) (res bool){
	for i :=2; i<=int(math.Floor(math.Sqrt(float64(val)))); i++{
		if val % i  == 0 {
			return res
		}
	}
	return val > 1
} 