package algolp

// Recherche renvoie s'il contient un élément val et l'indice de la première occurence
// si le tableau est vide ou ne contient pas l'élément, la fonction retourne -1 en indice
func Recherche(t []int, v int) (indice int, res bool){
	indice = -1
	for i, val := range t {
		if val == v{
			indice = i
			res = true
			break	
		}
	}
	return indice, res
}