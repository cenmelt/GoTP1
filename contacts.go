package main

import "fmt"

// ── STRUCT Adresse
type Adresse struct {
	Rue        string
	Ville      string
	CodePostal string
}

// Format retourne l'adresse en une seule ligne lisible
func (a Adresse) Format() string {
	return a.Rue + ", " + a.CodePostal + " " + a.Ville
}

// ── STRUCT Personne
type Personne struct {
	Prenom string
	Nom    string
	Age    int
	Email  string
}

// NomComplet retourne le prénom et le nom collés ensemble
func (p Personne) NomComplet() string {
	return p.Prenom + " " + p.Nom
}

// Presentation retourne une courte description de la personne
func (p Personne) Presentation() string {
	return fmt.Sprintf("%s, %d ans — %s", p.NomComplet(), p.Age, p.Email)
}

// ── STRUCT Employe 
type Employe struct {
	Personne // embedded : on peut écrire e.Prenom directement
	Adresse  // embedded : on peut écrire e.Ville directement
	Poste    string
	Salaire  float64
}

// FicheEmploye affiche toutes les infos de l'employé
func (e Employe) FicheEmploye() string {
	return fmt.Sprintf(
		"%s\n   Poste   : %s\n   Adresse : %s\n   Salaire : %.2f EUR\n   Contact : %s",
		e.Presentation(),
		e.Poste,
		e.Format(),
		e.Salaire,
		e.Email,
	)
}

// AugmenterSalaire augmente le salaire d'un pourcentage donné
func (e *Employe) AugmenterSalaire(pct float64) {
	e.Salaire = e.Salaire + e.Salaire*pct/100
}

// ── STRUCT Etudiant 
type Etudiant struct {
	Personne        // embedded
	Promo    string
	Moyenne  float64
}

// MentionObtenue retourne la mention selon la moyenne
func (et Etudiant) MentionObtenue() string {
	switch {
	case et.Moyenne >= 16:
		return "Très Bien"
	case et.Moyenne >= 14:
		return "Bien"
	case et.Moyenne >= 12:
		return "Assez Bien"
	case et.Moyenne >= 10:
		return "Passable"
	default:
		return "Insuffisant"
	}
}

// FicheEtudiant affiche toutes les infos de l'étudiant
func (et Etudiant) FicheEtudiant() string {
	return fmt.Sprintf(
		"%s\n   Promo   : %s\n   Moyenne : %.2f --> %s",
		et.Presentation(),
		et.Promo,
		et.Moyenne,
		et.MentionObtenue(),
	)
}

func main() {

	// Slice d'employés
	employes := []Employe{
		{
			Personne: Personne{"Sophie", "Martin", 34, "sophie.martin@corp.fr"},
			Adresse:  Adresse{"12 rue de la Paix", "Paris", "75001"},
			Poste:    "Developpeuse Go",
			Salaire:  3800.00,
		},
		{
			Personne: Personne{"Lucas", "Dupont", 41, "lucas.dupont@corp.fr"},
			Adresse:  Adresse{"5 avenue des Roses", "Lyon", "69002"},
			Poste:    "Chef de projet",
			Salaire:  4500.00,
		},
	}

	// Slice d'étudiants
	etudiants := []Etudiant{
		{
			Personne: Personne{"Emma", "Petit", 21, "emma.petit@univ.fr"},
			Promo:    "M1 Informatique 2025",
			Moyenne:  15.5,
		},
		{
			Personne: Personne{"Noah", "Bernard", 22, "noah.bernard@univ.fr"},
			Promo:    "M2 Data Science 2025",
			Moyenne:  11.0,
		},
	}

	// Affichage des fiches employés avec for range sur le slice
	fmt.Println("= EMPLOYES =")
	for i := range employes {
		// On utilise &employes[i] (pointeur) pour que AugmenterSalaire modifie le vrai élément
		employes[i].AugmenterSalaire(10)
		fmt.Println(employes[i].FicheEmploye())
		fmt.Println()
	}

	// Affichage des fiches étudiants avec for range sur le slice
	fmt.Println("= ETUDIANTS =")
	for _, et := range etudiants {
		fmt.Println(et.FicheEtudiant())
		fmt.Println()
	}
}
