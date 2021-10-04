package algolp

// Recherche renvoie s'il contient un élément val et l'indice de la première occurence qu'il trouve par dichotomie 
// si le tableau est vide ou ne contient pas l'élément, la fonction retourne -1 en indice
func Recherche(tab []int, val int) (indice int, res bool){
	if len(tab) == 0{
		return -1, false
	}

	var debut int = 0
	var fin int = len(tab)
	var milieu int
	for debut < fin{

		milieu = (debut + fin) / 2
		if tab[milieu] == val{
			return milieu, true
		}
		if tab[milieu] > val{
			fin = milieu
		}else{
			debut = milieu + 1
		}
	}
	return -1, false
}