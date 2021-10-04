package algolp


func contient(t []int, v int) (indice int, res bool){
	for i, val := range t {
		if val == v{
			indice = i
			res = true
			break	
		}
	}
	return indice, res
}