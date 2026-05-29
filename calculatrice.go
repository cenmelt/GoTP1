package main

import "fmt"

// operer calcule a op b et retourne le résultat + une éventuelle erreur
func operer(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("erreur : division par zéro")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("erreur : opération '%s' inconnue", op)
	}
}

// creerOperation retourne une closure qui applique l'opération donnée
func creerOperation(op string) func(float64, float64) float64 {
	return func(a, b float64) float64 {
		resultat, _ := operer(a, b, op)
		return resultat
	}
}

func main() {
	fmt.Println("= Calculatrice Go =")
	fmt.Println("Format : <nombre> <nombre> <opération>  |  'quit' pour quitter")

	for {
		fmt.Print(">> ")

		var a, b float64
		var op string

		// Lire les 3 valeurs d'un coup
		fmt.Scan(&a, &b, &op)

		// Quitter si l'opération est "quit"
		if op == "quit" {
			fmt.Println("Au revoir !")
			break
		}

		// Calculer et afficher le résultat
		resultat, err := operer(a, b, op)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%.2f %s %.2f = %.2f\n", a, op, b, resultat)
		}
	}
}
