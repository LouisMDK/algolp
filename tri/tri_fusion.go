package algolp
func TriFusion(liste []int) (res []int){
	if len(liste) == 0 || liste == nil {
		return res
	}
	if len(liste) == 1 {
		return liste
	}

	var mid int = len(liste) / 2
	var t1 []int = liste[:mid]
	var t2 []int = liste[mid:]
	t1 = TriFusion(t1)
	t2 = TriFusion(t2)
	res = fusion(t1, t2)
	return res
}

func fusion(liste1, liste2 []int) (res []int) {
	
	for len(liste1) > 0 && len(liste2) > 0 {
		if liste1[0] < liste2[0] {
			res = append(res, liste1[0])
			liste1 = liste1[1:]
		}else{
			res = append(res, liste2[0])
			liste2 = liste2[1:]
		}
	}

	if len(liste1) > len(liste2) {
		res = append(res, liste1...)
	}else{
		res = append(res, liste2...)
	}

	return res
}

