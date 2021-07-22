package main

import "fmt"

func main() { //nolint:typecheck
	regions := []string{"Auvergne-Rhone-Aples","Bourgogne-Franche-Comté","Bretagne","Centre-Val de Loire","Corse","Grand Est","Hauts-de-France","ile-de-France","Normandie","Nouvelle-Aquitaine","Occitanie","Pays de la loire","Provence Alpes","Guadeloupe","Guyane",
		"La Réunion","Martinique","Mayotte"}

	fmt.Println("slice capacity: ",cap(regions))
	fmt.Println("slice length:",len(regions))
	for i := 0; i < len(regions); i++ {
		fmt.Println(i,regions[i])
	}





}
