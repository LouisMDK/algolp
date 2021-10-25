package algolp


func TriRapide(liste []int){
	
	var tab *[]int = &liste
	tri(tab, 0, len(liste))
}

func tri(liste *[]int, p, r int) {
	if len((*liste)[p:r]) > 1 {
		var q int
		q = partitionner(liste, p, r)
		tri(liste, p, q)
		tri(liste, q+1, r)
	}
}

func partitionner(liste *[]int, p, r int) (int){
	var pivot int = (*liste)[p]
	var i int = p
	var j int = r - 1

	for i < j {

		for (*liste)[j] >= pivot && i < j {
			j--
		}
		if i < j {
			(*liste)[i], (*liste)[j] = (*liste)[j], (*liste)[i]
		}

		for (*liste)[i] <= pivot && i < j {
			i++
		}
		if i < j {
			(*liste)[i], (*liste)[j] = (*liste)[j], (*liste)[i]
		}
	}
	return i
}