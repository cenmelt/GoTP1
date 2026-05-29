package main

import "fmt"

const (
	Debutant     = iota // 0
	Intermediaire       // 1
	Expert              // 2
)

func main() {

	joueurs := []string{"Alice", "Bob", "Charlie"}
	scores := []int{42, 87, 55}

	fmt.Println("= Classement des joueurs =")


	for i := 0; i < len(joueurs); i++ {
		fmt.Printf("%d. %s --> score : %d\n", i+1, joueurs[i], scores[i])
	}

	fmt.Println()
	// Score de chaque joueur via un for range
	fmt.Println("=== Scores ===")
	for i, joueur := range joueurs {
		fmt.Printf("%s a marqué %d points\n", joueur, scores[i])
	}

	fmt.Println()
	// un compte à rebours via un while
	fmt.Println("= Compte à rebours =")
	compteur := 3
	for compteur > 0 {
		fmt.Println(compteur, "...")
		compteur--
	}
	fmt.Println("GO !")

	fmt.Println()
	// système d'inventaire
	inventaire := make([]string, 2, 5)

	inventaire[0] = "Épée"     // poche 0 : remplie
	inventaire[1] = "Bouclier" // poche 1 : remplie
	//                            poches 2, 3, 4 : vides (réservées)

	fmt.Println("=== Inventaire du joueur ===")
	fmt.Printf("Objets       : %v\n", inventaire)
	fmt.Printf("len = %d  --> 2 poches utilisées\n", len(inventaire))
	fmt.Printf("cap = %d  --> 5 poches au total dans le sac\n", cap(inventaire))


	inventaire = append(inventaire, "Potion")

	fmt.Println()
	fmt.Println("Après avoir ajouté une Potion :")
	fmt.Printf("Objets       : %v\n", inventaire)
	fmt.Printf("len = %d  --> 3 poches utilisées\n", len(inventaire))

	fmt.Println()

	// switch pour afficher les niveaux d'accès d'un joueur
	fmt.Println("= Niveau du joueur Alice =")
	niveau := Debutant // valeur iota = 0

	switch niveau {
	case Debutant:
		fmt.Println("Niveau : Débutant")
		fmt.Println("  --> Accès au tutoriel")
		fallthrough // on continue vers le cas suivant malgré tout
	case Intermediaire:
		fmt.Println("  --> Accès aux niveaux de base")
		fallthrough
	case Expert:
		fmt.Println("  --> Accès aux règles du jeu")
	}
}
