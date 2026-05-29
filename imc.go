package main

import "fmt"

	// ── Constantes pour les catégories d'IMC
const (
	IMCMaigreur = 18.5 // En dessous --> Maigreur
	IMCNormal   = 25.0 // En dessous --> Normal
	IMCSurpoids = 30.0 // En dessous --> Surpoids, sinon --> Obésité
	Nom         = "Clément"
)

func main() {

	// ── Déclaration des variables
	var (
		poids  float64 = 90.5 // Poids en kilogrammes
		taille float64 = 1.80 // Taille en mètres
	)

	// ── Calcul de l'IMC
	imc := poids / (taille * taille)

	// ──affichage de la catégorie
	if imc < IMCMaigreur {
		fmt.Printf("%s est dans la Catégorie : Maigreur. Car votre IMC est de %.2f\n", Nom, imc)
	} else if imc < IMCNormal {
		fmt.Printf("%s est dans la Catégorie : Normal. Car votre IMC est de %.2f\n", Nom, imc)
	} else if imc < IMCSurpoids {
		fmt.Printf("%s est dans la Catégorie : Surpoids. Car votre IMC est de %.2f\n", Nom, imc)
	} else {
		fmt.Printf("%s est dans la Catégorie : Obésité. Car votre IMC est de %.2f\n", Nom, imc)
	}
}
