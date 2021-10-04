package algolp

// Selection renvoie un nouveau tableau d'entier contenant tous les entiers d'un tableau trié inférieurs ou égaux à une valeur
func Selection(tab []int, val int) (res []int){
	fin, _ := Recherche(tab, val)
	if fin == -1{
		return res
	}
	return tab[:fin]
}