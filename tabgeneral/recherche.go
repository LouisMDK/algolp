package algolp

// Recherche renvoie l'indice du premier élément dans un tableau d'entiers et s'il le contient ou non 
// si le tableau ne contient pas l'élément, la fonction retourne -1 en indice
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