package db

import (
	"database/sql"
	"testing"
)

func createTestOuvrages(testDB *sql.DB) error {
	//create TYPE_OUVRAGE
	_, err := testDB.Exec("INSERT INTO TYPE_OUVRAGE (ID_TYPE) VALUES ('type1'), ('type2'), ('type3')")
	if err != nil {
		return err
	}

	//create EDITEUR
	_, err = testDB.Exec("INSERT INTO EDITEUR (ID_EDITEUR, NOM_EDITEUR, TEL_EDITEUR, MAIL_EDITEUR, ADRESSE_EDITEUR) VALUES (1, 'nom1', 'tel1', 'mail1', 'adresse1'), (2, 'nom2', 'tel2', 'mail2', 'adresse2'), (3, 'nom3', 'tel3', 'mail3', 'adresse3')")
	if err != nil {
		return err
	}

	//create OUVRAGE
	_, err = testDB.Exec("INSERT INTO OUVRAGE (ID_OUVRAGE, TITRE_OUVRAGE, PARUTION_OUVRAGE, ID_TYPE, ID_EDITEUR) VALUES (1, 'titre1', '2023-02-01', 'type1', 1), (2, 'titre2', '2023-02-01', 'type2', 2), (3, 'titre3', '2023-02-01', 'type3', 3)")
	if err != nil {
		return err
	}

	return nil
}

func TestGetAllOuvrage(t *testing.T) {
	testDB, err := InitTestDB()

	err = createTestOuvrages(testDB)
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	ouvrages, err := GetAllOuvrage(testDB)
	if err != nil {
		t.Fatalf("Erreur dans GetAllOuvrage : %v", err)
	}

	if len(ouvrages) != 3 {
		t.Errorf("Le nombre d'ouvrages ne correspond pas à celui attendu.")
	}

	expectedIDs := []int{1, 2, 3}
	for i, ouvrage := range ouvrages {
		if ouvrage.ID_OUVRAGE != expectedIDs[i] {
			t.Errorf("La valeur de l'ID_OUVRAGE ne correspond pas à celle attendue.")
		}
	}
}

func TestGetOuvrageById(t *testing.T) {
	testDB, err := InitTestDB()

	err = createTestOuvrages(testDB)
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	ouvrage, err := GetOuvrageById(testDB, 2)
	if err != nil {
		t.Fatalf("Erreur dans GetOuvrageById : %v", err)
	}

	if ouvrage.ID_OUVRAGE != 2 {
		t.Errorf("La valeur de l'ID_OUVRAE ne correspond pas à celle attendue.")
	}
}

func TestDeleteOuvrage(t *testing.T) {
	testDB, err := InitTestDB()

	err = createTestOuvrages(testDB)
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	err = DeleteOuvrage(testDB, 2)
	if err != nil {
		t.Fatalf("Erreur dans DeleteOuvrage : %v", err)
	}

	ouvrages, err := GetAllOuvrage(testDB)
	if err != nil {
		t.Fatalf("Erreur dans GetAllOuvrage : %v", err)
	}

	if len(ouvrages) != 2 {
		t.Errorf("Le nombre d'ouvrages ne correspond pas à celui attendu.")
	}

	expectedIDs := []int{1, 3}
	for i, ouvrage := range ouvrages {
		if ouvrage.ID_OUVRAGE != expectedIDs[i] {
			t.Errorf("La valeur de l'ID_OUVRAGE ne correspond pas à celle attendue.")
		}
	}
}

func TestPostOuvrage(t *testing.T) {
	testDB, err := InitTestDB()

	//create TYPE_OUVRAGE
	_, err = testDB.Exec("INSERT INTO TYPE_OUVRAGE (ID_TYPE) VALUES ('type1'), ('type2'), ('type3')")
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	//create EDITEUR
	_, err = testDB.Exec("INSERT INTO EDITEUR (ID_EDITEUR, NOM_EDITEUR, TEL_EDITEUR, MAIL_EDITEUR, ADRESSE_EDITEUR) VALUES (1, 'nom1', 'tel1', 'mail1', 'adresse1'), (2, 'nom2', 'tel2', 'mail2', 'adresse2'), (3, 'nom3', 'tel3', 'mail3', 'adresse3')")
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	ouvrage := OUVRAGE{
		ID_OUVRAGE:       4,
		TITRE_OUVRAGE:    "titre4",
		PARUTION_OUVRAGE: "2023-02-01",
		ID_TYPE:          "type1",
		ID_EDITEUR:       1,
	}

	err = PostOuvrage(testDB, ouvrage)
	if err != nil {
		t.Fatalf("Erreur dans PostOuvrage : %v", err)
	}

	ouvrages, err := GetAllOuvrage(testDB)
	if err != nil {
		t.Fatalf("Erreur dans GetAllOuvrage : %v", err)
	}

	if len(ouvrages) != 1 {
		t.Errorf("Le nombre d'ouvrages ne correspond pas à celui attendu.")
	}
}
