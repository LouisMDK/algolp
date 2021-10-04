package algolp

import "math"

// La fonction EstPremier renvoie vrai si l'entier donné est un nombre premier, et faux sinon
func EstPremier(val int) (res bool){
	for i :=2; i<=int(math.Floor(math.Sqrt(float64(val)))); i++{
		if val % i  == 0 {
			return res
		}
	}
	return val > 1
} 