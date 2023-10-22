package db

import (
	"database/sql"
	"testing"
)

func createTestAuteur(testDB *sql.DB) error {
	_, err := testDB.Exec("INSERT INTO AUTEUR (ID_AUTEUR, NOM_AUTEUR, PRENOM_AUTEUR, TEL_AUTEUR, MAIL_AUTEUR, ADRESSE_AUTEUR) VALUES (1, 'nom1', 'prenom1', 'tel1', 'mail1', 'adresse1'), (2, 'nom2', 'prenom2', 'tel2', 'mail2', 'adresse2'), (3, 'nom3', 'prenom3', 'tel3', 'mail3', 'adresse3')")
	if err != nil {
		return err
	}

	return nil
}

func TestGetAllAuteur(t *testing.T) {
	testDB, err := InitTestDB()

	err = createTestAuteur(testDB)
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	auteurs, err := GetAllAuteur(testDB)
	if err != nil {
		t.Fatalf("Erreur dans GetAllAuteur : %v", err)
	}

	if len(auteurs) != 3 {
		t.Errorf("Le nombre d'auteurs ne correspond pas à celui attendu.")
	}

	expectedIDs := []int{1, 2, 3}
	for i, auteur := range auteurs {
		if auteur.ID_AUTEUR != expectedIDs[i] {
			t.Errorf("La valeur de l'ID_AUTEUR ne correspond pas à celle attendue.")
		}
	}
}

func TestGetAuteurById(t *testing.T) {
	testDB, err := InitTestDB()

	err = createTestAuteur(testDB)
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	auteur, err := GetAuteurById(testDB, 2)
	if err != nil {
		t.Fatalf("Erreur dans GetAuteurById : %v", err)
	}

	if auteur.ID_AUTEUR != 2 {
		t.Errorf("La valeur de l'ID_AUTEUR ne correspond pas à celle attendue.")
	}
}

func TestDeleteAuteur(t *testing.T) {
	testDB, err := InitTestDB()

	err = createTestAuteur(testDB)
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	err = DeleteAuteur(testDB, 2)
	if err != nil {
		t.Fatalf("Erreur dans DeleteAuteur : %v", err)
	}

	auteurs, err := GetAllAuteur(testDB)
	if err != nil {
		t.Fatalf("Erreur dans GetAllAuteur : %v", err)
	}

	if len(auteurs) != 2 {
		t.Errorf("Le nombre d'auteurs ne correspond pas à celui attendu.")
	}

	expectedIDs := []int{1, 3}
	for i, auteur := range auteurs {
		if auteur.ID_AUTEUR != expectedIDs[i] {
			t.Errorf("La valeur de l'ID_AUTEUR ne correspond pas à celle attendue.")
		}
	}
}

func TestPostAuteur(t *testing.T) {
	testDB, err := InitTestDB()

	auteur := AUTEUR{ID_AUTEUR: 1, NOM_AUTEUR: "nom1", PRENOM_AUTEUR: "prenom1", TEL_AUTEUR: "tel1", MAIL_AUTEUR: "mail1", ADRESSE_AUTEUR: "adresse1"}

	err = PostAuteur(testDB, auteur)
	if err != nil {
		t.Fatalf("Erreur dans PostAuteur : %v", err)
	}

	auteurs, err := GetAllAuteur(testDB)
	if err != nil {
		t.Fatalf("Erreur dans GetAllAuteur : %v", err)
	}

	if len(auteurs) != 1 {
		t.Errorf("Le nombre d'auteurs ne correspond pas à celui attendu.")
	}

}

func TestPutAuteur(t *testing.T) {
	testDB, err := InitTestDB()

	err = createTestAuteur(testDB)
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	auteur := AUTEUR{ID_AUTEUR: 2, NOM_AUTEUR: "nom2", PRENOM_AUTEUR: "prenom2", TEL_AUTEUR: "tel2", MAIL_AUTEUR: "mail2", ADRESSE_AUTEUR: "adresse2"}

	err = PutAuteur(testDB, auteur)
	if err != nil {
		t.Fatalf("Erreur dans PutAuteur : %v", err)
	}

	auteurs, err := GetAllAuteur(testDB)
	if err != nil {
		t.Fatalf("Erreur dans GetAllAuteur : %v", err)
	}

	if len(auteurs) != 3 {
		t.Errorf("Le nombre d'auteurs ne correspond pas à celui attendu.")
	}

	expectedIDs := []int{1, 2, 3}
	for i, auteur := range auteurs {
		if auteur.ID_AUTEUR != expectedIDs[i] {
			t.Errorf("La valeur de l'ID_AUTEUR ne correspond pas à celle attendue.")
		}
	}
}
