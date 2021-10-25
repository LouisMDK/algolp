package algolp


func TriInsertion(liste []int) ([]int) {
	
	var t2 []int
	var i int
	for i < len(liste) {


		var j int
		var v int = liste[i]
		for j < len(t2) && t2[j] <= v {
			j++
		}
		var tmp int
		for j < len(t2) {
			tmp = t2[j]
			t2[j] = v
			v = tmp
			j++
		}
		t2 = append(t2, v)




		i++
	}
	return t2
}
