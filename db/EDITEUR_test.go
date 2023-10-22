package db

import (
	"database/sql"
	"testing"
)

func createTestEditeurs(testDB *sql.DB) error {
	_, err := testDB.Exec("INSERT INTO EDITEUR (ID_EDITEUR, NOM_EDITEUR, TEL_EDITEUR, MAIL_EDITEUR, ADRESSE_EDITEUR) VALUES (1, 'nom1', 'tel1', 'mail1', 'adresse1'), (2, 'nom2', 'tel2', 'mail2', 'adresse2'), (3, 'nom3', 'tel3', 'mail3', 'adresse3')")
	if err != nil {
		return err
	}

	return nil
}

func TestGetAllEditeur(t *testing.T) {
	testDB, err := InitTestDB()

	err = createTestEditeurs(testDB)
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	editeurs, err := GetAllEditeur(testDB)
	if err != nil {
		t.Fatalf("Erreur dans GetAllEditeur : %v", err)
	}

	if len(editeurs) != 3 {
		t.Errorf("Le nombre d'éditeurs ne correspond pas à celui attendu.")
	}

	expectedIDs := []int{1, 2, 3}
	for i, editeur := range editeurs {
		if editeur.ID_EDITEUR != expectedIDs[i] {
			t.Errorf("La valeur de l'ID_EDITEUR ne correspond pas à celle attendue.")
		}
	}
}

func TestGetEditeurById(t *testing.T) {
	testDB, err := InitTestDB()

	err = createTestEditeurs(testDB)
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	editeur, err := GetEditeurById(testDB, 2)
	if err != nil {
		t.Fatalf("Erreur dans GetEditeurById : %v", err)
	}

	if editeur.ID_EDITEUR != 2 {
		t.Errorf("La valeur de l'ID_EDITEUR ne correspond pas à celle attendue.")
	}
}

func TestDeleteEditeur(t *testing.T) {
	testDB, err := InitTestDB()

	err = createTestEditeurs(testDB)
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	err = DeleteEditeur(testDB, 2)
	if err != nil {
		t.Fatalf("Erreur dans DeleteEditeur : %v", err)
	}

	editeurs, err := GetAllEditeur(testDB)
	if err != nil {
		t.Fatalf("Erreur dans GetAllEditeur : %v", err)
	}

	if len(editeurs) != 2 {
		t.Errorf("Le nombre d'éditeurs ne correspond pas à celui attendu.")
	}

	expectedIDs := []int{1, 3}
	for i, editeur := range editeurs {
		if editeur.ID_EDITEUR != expectedIDs[i] {
			t.Errorf("La valeur de l'ID_EDITEUR ne correspond pas à celle attendue.")
		}
	}
}

func TestPostEditeur(t *testing.T) {
	testDB, err := InitTestDB()

	editeur := EDITEUR{NOM_EDITEUR: "nom", TEL_EDITEUR: "tel", MAIL_EDITEUR: "mail", ADRESSE_EDITEUR: "adresse"}

	err = PostEditeur(testDB, editeur)
	if err != nil {
		t.Fatalf("Erreur dans PostEditeur : %v", err)
	}

	editeurs, err := GetAllEditeur(testDB)
	if err != nil {
		t.Fatalf("Erreur dans GetAllEditeur : %v", err)
	}

	if len(editeurs) != 1 {
		t.Errorf("Le nombre d'éditeurs ne correspond pas à celui attendu.")
	}

}
