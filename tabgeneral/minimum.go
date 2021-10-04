package algolp

import "errors"
var errPasTrouve error = errors.New("Aucun minimum trouvé")
var errTabVide error = errors.New("Tableau vide")


// Minimum renvoie l'indice de la plus petite valeur d'un tableau d'entiers
// si le tableau est vide, la fonction renvoie une erreur
// si le minimum apparaît plusieurs fois, la fonction renvoie l'indice de la première occurence
func Minimum(tab []int) (err error, indice int){
	if len(tab) == 0{
		return errTabVide, 0
	}
	var min int = tab[0]
	for i:=1;i<len(tab);i++{
		if tab[i] < min{
			min = tab[i]
			indice = i
		}
	}
	return err, indice
}