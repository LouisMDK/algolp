package algolp

// Selection renvoie un nouveau tableau d'entier contenant tous les entiers d'un tableau inférieurs ou égaux à une valeur
func Selection(tab []int, val int) (res []int){
	for _, nombre := range tab{
		if nombre <= val{
			res = append(res, nombre)
		}
	}
	return res
}