package algolp


func TriInsertionPlace(liste []int) {
	var i int = 1
	var v int
	var j int
	for i < len(liste) {
		v = liste[i]
		j = i - 1

		for j >= 0 && liste[j] > v {
			liste[j + 1] = liste[j]
			j--
		}
		liste[j + 1] = v
		i++
	}
}
