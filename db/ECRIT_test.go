package db

import (
	"database/sql"
	"testing"
)

func createTestEcrit(testDB *sql.DB) error {
	//create AUTEUR
	_, err := testDB.Exec("INSERT INTO AUTEUR (ID_AUTEUR, NOM_AUTEUR, PRENOM_AUTEUR, TEL_AUTEUR, MAIL_AUTEUR, ADRESSE_AUTEUR) VALUES (1, 'nom1', 'prenom1', 'tel1', 'mail1', 'adresse1'), (2, 'nom2', 'prenom2', 'tel2', 'mail2', 'adresse2'), (3, 'nom3', 'prenom3', 'tel3', 'mail3', 'adresse3')")
	if err != nil {
		return err
	}

	//create EDITEUR
	_, err = testDB.Exec("INSERT INTO EDITEUR (ID_EDITEUR, NOM_EDITEUR, TEL_EDITEUR, MAIL_EDITEUR, ADRESSE_EDITEUR) VALUES (1, 'nom1', 'tel1', 'mail1', 'adresse1'), (2, 'nom2', 'tel2', 'mail2', 'adresse2'), (3, 'nom3', 'tel3', 'mail3', 'adresse3')")
	if err != nil {
		return err
	}

	//create TYPE_OUVRAGE
	_, err = testDB.Exec("INSERT INTO TYPE_OUVRAGE (ID_TYPE) VALUES ('1'), ('2'), ('3')")
	if err != nil {
		return err
	}

	//create OUVRAGE
	_, err = testDB.Exec("INSERT INTO OUVRAGE (ID_OUVRAGE, TITRE_OUVRAGE, PARUTION_OUVRAGE, ID_TYPE, ID_EDITEUR) VALUES (1, 'titre1', '2023-01-02', '1', 1), (2, 'titre2', '2023-01-02', '2', 2), (3, 'titre3', '2023-01-02', '3', 3)")
	if err != nil {
		return err
	}

	//create ECRIT
	_, err = testDB.Exec("INSERT INTO ECRIT (ID_ECRIT, ID_OUVRAGE, ID_AUTEUR) VALUES (1, 1, 1), (2, 2, 2), (3, 3, 3)")
	if err != nil {
		return err
	}

	return nil
}

func TestGetAllEcrit(t *testing.T) {
	testDB, err := InitTestDB()

	err = createTestEcrit(testDB)
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	ecrits, err := GetAllEcrit(testDB)
	if err != nil {
		t.Fatalf("Erreur dans GetAllEcrit : %v", err)
	}

	expectedIDs := []int{1, 2, 3}
	for i, ecrit := range ecrits {
		if ecrit.ID_ECRIT != expectedIDs[i] {
			t.Errorf("La valeur de l'ID_ECRIT ne correspond pas à celle attendue.")
		}
	}
}

func TestGetEcritById(t *testing.T) {
	testDB, err := InitTestDB()

	err = createTestEcrit(testDB)
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	ecrit, err := GetEcritById(testDB, 2)
	if err != nil {
		t.Fatalf("Erreur dans GetEcritById : %v", err)
	}

	if ecrit.ID_ECRIT != 2 {
		t.Errorf("La valeur de l'ID_ECRIT ne correspond pas à celle attendue.")
	}
}

func TestDeleteEcrit(t *testing.T) {
	testDB, err := InitTestDB()

	err = createTestEcrit(testDB)
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	err = DeleteEcrit(testDB, 2)
	if err != nil {
		t.Fatalf("Erreur dans DeleteEcrit : %v", err)
	}

	ecrits, err := GetAllEcrit(testDB)
	if err != nil {
		t.Fatalf("Erreur dans GetAllEcrit : %v", err)
	}

	if len(ecrits) != 2 {
		t.Errorf("Le nombre d'écrits ne correspond pas à celui attendu.")
	}

	expectedIDs := []int{1, 3}
	for i, ecrit := range ecrits {
		if ecrit.ID_ECRIT != expectedIDs[i] {
			t.Errorf("La valeur de l'ID_ECRIT ne correspond pas à celle attendue.")
		}
	}
}

func TestPostEcrit(t *testing.T) {
	testDB, err := InitTestDB()

	//create AUTEUR
	_, err = testDB.Exec("INSERT INTO AUTEUR (ID_AUTEUR, NOM_AUTEUR, PRENOM_AUTEUR, TEL_AUTEUR, MAIL_AUTEUR, ADRESSE_AUTEUR) VALUES (1, 'nom1', 'prenom1', 'tel1', 'mail1', 'adresse1'), (2, 'nom2', 'prenom2', 'tel2', 'mail2', 'adresse2'), (3, 'nom3', 'prenom3', 'tel3', 'mail3', 'adresse3')")
	if err != nil {
		t.Errorf("Erreur lors de l'insertion de données de test : %v", err)
	}

	//create EDITEUR
	_, err = testDB.Exec("INSERT INTO EDITEUR (ID_EDITEUR, NOM_EDITEUR, TEL_EDITEUR, MAIL_EDITEUR, ADRESSE_EDITEUR) VALUES (1, 'nom1', 'tel1', 'mail1', 'adresse1'), (2, 'nom2', 'tel2', 'mail2', 'adresse2'), (3, 'nom3', 'tel3', 'mail3', 'adresse3')")
	if err != nil {
		t.Errorf("Erreur lors de l'insertion de données de test : %v", err)
	}

	//create TYPE_OUVRAGE
	_, err = testDB.Exec("INSERT INTO TYPE_OUVRAGE (ID_TYPE) VALUES ('1'), ('2'), ('3')")
	if err != nil {
		t.Errorf("Erreur lors de l'insertion de données de test : %v", err)
	}

	//create OUVRAGE
	_, err = testDB.Exec("INSERT INTO OUVRAGE (ID_OUVRAGE, TITRE_OUVRAGE, PARUTION_OUVRAGE, ID_TYPE, ID_EDITEUR) VALUES (1, 'titre1', '2023-01-02', '1', 1), (2, 'titre2', '2023-01-02', '2', 2), (3, 'titre3', '2023-01-02', '3', 3)")
	if err != nil {
		t.Errorf("Erreur lors de l'insertion de données de test : %v", err)
	}

	ecrit := ECRIT{
		ID_OUVRAGE: 2,
		ID_AUTEUR:  2,
	}

	err = PostEcrit(testDB, ecrit)
	if err != nil {
		t.Fatalf("Erreur dans PostEcrit : %v", err)
	}

	ecrits, err := GetAllEcrit(testDB)
	if err != nil {
		t.Fatalf("Erreur dans GetAllEcrit : %v", err)
	}

	if len(ecrits) != 1 {
		t.Errorf("Le nombre d'écrits ne correspond pas à celui attendu.")
	}

}

func TestPutEcrit(t *testing.T) {
	testDB, err := InitTestDB()

	err = createTestEcrit(testDB)
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	ecrit := ECRIT{
		ID_ECRIT:   2,
		ID_OUVRAGE: 2,
		ID_AUTEUR:  2,
	}

	err = PutEcrit(testDB, ecrit)
	if err != nil {
		t.Fatalf("Erreur dans PutEcrit : %v", err)
	}

	ecrits, err := GetAllEcrit(testDB)
	if err != nil {
		t.Fatalf("Erreur dans GetAllEcrit : %v", err)
	}

	if len(ecrits) != 3 {
		t.Errorf("Le nombre d'écrits ne correspond pas à celui attendu.")
	}

	expectedIDs := []int{1, 2, 3}
	for i, ecrit := range ecrits {
		if ecrit.ID_ECRIT != expectedIDs[i] {
			t.Errorf("La valeur de l'ID_ECRIT ne correspond pas à celle attendue.")
		}
	}
}
