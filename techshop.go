package main

import (
	"fmt"
	"strings"
)

// ── STRUCTURES
type Produit struct {
	ID        int
	Nom       string
	Marque    string
	Prix      float64
	Stock     int
	Categorie string
	Actif     bool
}

type Catalogue struct {
	Produits []Produit
}

// ── MÉTHODES DU CATALOGUE

// AjouterProduit ajoute un produit au catalogue.
func (c *Catalogue) AjouterProduit(p Produit) error {
	// Vérifier que l'ID n'existe pas déjà
	for _, existant := range c.Produits {
		if existant.ID == p.ID {
			return fmt.Errorf("erreur : un produit avec l'ID %d existe deja", p.ID)
		}
	}
	c.Produits = append(c.Produits, p)
	return nil
}

// TrouverParID cherche un produit par son ID.
func (c Catalogue) TrouverParID(id int) (Produit, error) {
	for _, p := range c.Produits {
		if p.ID == id {
			return p, nil
		}
	}
	return Produit{}, fmt.Errorf("erreur : produit avec l'ID %d introuvable", id)
}

// TrouverParCategorie retourne tous les produits d'une catégorie.
func (c Catalogue) TrouverParCategorie(cat string) []Produit {
	var resultats []Produit
	for _, p := range c.Produits {
		if strings.EqualFold(p.Categorie, cat) {
			resultats = append(resultats, p)
		}
	}
	return resultats
}

// AppliquerReduction applique une réduction (%) sur tous les produits d'une catégorie.
func (c *Catalogue) AppliquerReduction(categorie string, pct float64) int {
	modifies := 0
	for i, p := range c.Produits {
		if strings.EqualFold(p.Categorie, categorie) {
			c.Produits[i].Prix = p.Prix - p.Prix*pct/100
			modifies++
		}
	}
	return modifies
}

// Vendre réduit le stock d'un produit.
func (c *Catalogue) Vendre(id int, qte int) error {
	for i, p := range c.Produits {
		if p.ID == id {
			if p.Stock < qte {
				return fmt.Errorf("erreur : stock insuffisant (%d disponible, %d demande)", p.Stock, qte)
			}
			c.Produits[i].Stock -= qte
			return nil
		}
	}
	return fmt.Errorf("erreur : produit avec l'ID %d introuvable", id)
}

// Rapport retourne un résumé du catalogue.
func (c Catalogue) Rapport() string {
	total := 0.0
	for _, p := range c.Produits {
		total += p.Prix * float64(p.Stock)
	}
	return fmt.Sprintf("Nombre de produits : %d\nValeur totale du stock : %.2f EUR", len(c.Produits), total)
}

// ── AFFICHAGE D'UN PRODUIT

func afficherProduit(p Produit) {
	fmt.Printf("  [%d] %s %s | Prix: %.2f EUR | Stock: %d | Categorie: %s\n",
		p.ID, p.Marque, p.Nom, p.Prix, p.Stock, p.Categorie)
}

// ── MAIN

func main() {

	// Création du catalogue et ajout de 5 produits de départ
	catalogue := Catalogue{}

	catalogue.AjouterProduit(Produit{1, "iPhone 15", "Apple", 999.99, 10, "Smartphone", true})
	catalogue.AjouterProduit(Produit{2, "MacBook Pro", "Apple", 2499.99, 5, "Laptop", true})
	catalogue.AjouterProduit(Produit{3, "Galaxy S24", "Samsung", 849.99, 8, "Smartphone", true})
	catalogue.AjouterProduit(Produit{4, "ThinkPad X1", "Lenovo", 1599.99, 4, "Laptop", true})
	catalogue.AjouterProduit(Produit{5, "MX Master 3", "Logitech", 99.99, 20, "Accessoire", true})

	fmt.Println("=== Bienvenue sur TechShop ===")

	// Boucle du menu CLI
	for {
		fmt.Println("\n[1] Ajouter  [2] Chercher  [3] Soldes  [4] Vendre  [5] Rapport  [0] Quitter")
		fmt.Print("Votre choix : ")

		var choix int
		fmt.Scan(&choix)

		switch choix {

		case 1: // Ajouter un produit
			var p Produit
			fmt.Print("ID : ")
			fmt.Scan(&p.ID)
			fmt.Print("Nom : ")
			fmt.Scan(&p.Nom)
			fmt.Print("Marque : ")
			fmt.Scan(&p.Marque)
			fmt.Print("Prix : ")
			fmt.Scan(&p.Prix)
			fmt.Print("Stock : ")
			fmt.Scan(&p.Stock)
			fmt.Print("Categorie : ")
			fmt.Scan(&p.Categorie)
			p.Actif = true

			err := catalogue.AjouterProduit(p)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Produit ajoute avec succes !")
			}

		case 2: // Chercher par ID ou catégorie
			fmt.Println("Chercher par [1] ID  [2] Categorie")
			var sousChoix int
			fmt.Scan(&sousChoix)

			if sousChoix == 1 {
				fmt.Print("ID du produit : ")
				var id int
				fmt.Scan(&id)
				p, err := catalogue.TrouverParID(id)
				if err != nil {
					fmt.Println(err)
				} else {
					afficherProduit(p)
				}
			} else {
				fmt.Print("Categorie : ")
				var cat string
				fmt.Scan(&cat)
				produits := catalogue.TrouverParCategorie(cat)
				if len(produits) == 0 {
					fmt.Println("Aucun produit trouve dans cette categorie.")
				} else {
					for _, p := range produits {
						afficherProduit(p)
					}
				}
			}

		case 3: // Appliquer une réduction
			fmt.Print("Categorie : ")
			var cat string
			fmt.Scan(&cat)
			fmt.Print("Reduction (%) : ")
			var pct float64
			fmt.Scan(&pct)
			nb := catalogue.AppliquerReduction(cat, pct)
			fmt.Printf("Reduction de %.0f%% appliquee sur %d produit(s).\n", pct, nb)

		case 4: // Vendre
			fmt.Print("ID du produit : ")
			var id int
			fmt.Scan(&id)
			fmt.Print("Quantite : ")
			var qte int
			fmt.Scan(&qte)
			err := catalogue.Vendre(id, qte)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Vente enregistree !")
			}

		case 5: // Rapport
			fmt.Println(catalogue.Rapport())

		case 0: // Quitter
			fmt.Println("Au revoir !")
			return

		default:
			fmt.Println("Choix invalide, reessayez.")
		}
	}
}
