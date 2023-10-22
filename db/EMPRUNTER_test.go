package db

import (
	"database/sql"
	"testing"
)

func CreateTestEmprunter(testDB *sql.DB) error {

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

	//create CATEGORY_ABONNE
	_, err = testDB.Exec("INSERT INTO CATEGORIE_ABONNE (ID_CATEGORIE) VALUES ('1'), ('2'), ('3')")
	if err != nil {
		return err
	}

	//create ABONNE
	_, err = testDB.Exec("INSERT INTO ABONNE (ID_ABONNE, NOM_ABONNE, PRENOM_ABONNE, TEL_ABONNE, MAIL_ABONNE, ADRESSE_ABONNE, ID_CATEGORIE) VALUES (1, 'nom1', 'prenom1', 'tel1', 'mail1', 'adresse1', '1'), (2, 'nom2', 'prenom2', 'tel2', 'mail2', 'adresse2', '2'), (3, 'nom3', 'prenom3', 'tel3', 'mail3', 'adresse3', '3')")
	if err != nil {
		return err
	}

	//create EMPRUNTER
	_, err = testDB.Exec("INSERT INTO EMPRUNTER (ID_EMPRUNT, ID_OUVRAGE, ID_ABONNE, DATE_SORTIE) VALUES (1, 1, 1, '2023-02-01'), (2, 2, 2, '2023-02-01'), (3, 3, 3, '2023-02-01')")
	if err != nil {
		return err
	}

	return nil
}

func TestGetAllEmprunter(t *testing.T) {
	testDB, err := InitTestDB()

	err = CreateTestEmprunter(testDB)
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	emprunter, err := GetAllEmprunter(testDB)
	if err != nil {
		t.Fatalf("Erreur dans GetAllEmprunter : %v", err)
	}

	if len(emprunter) != 3 {
		t.Errorf("Le nombre d'emprunter ne correspond pas à celui attendu.")
	}

	expectedIDs := []int{1, 2, 3}
	for i, emprunter := range emprunter {
		if emprunter.ID_EMPRUNT != expectedIDs[i] {
			t.Errorf("La valeur de l'ID_EMPRUNT ne correspond pas à celle attendue.")
		}
	}
}

func TestGetEmprunterById(t *testing.T) {
	testDB, err := InitTestDB()

	err = CreateTestEmprunter(testDB)
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	emprunter, err := GetEmprunterById(testDB, 2)
	if err != nil {
		t.Fatalf("Erreur dans GetEmprunterById : %v", err)
	}

	if emprunter.ID_EMPRUNT != 2 {
		t.Errorf("La valeur de l'ID_EMPRUNT ne correspond pas à celle attendue.")
	}
}

func TestDeleteEmprunter(t *testing.T) {
	testDB, err := InitTestDB()

	err = CreateTestEmprunter(testDB)
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	err = DeleteEmprunter(testDB, 2)
	if err != nil {
		t.Fatalf("Erreur dans DeleteEmprunter : %v", err)
	}

	emprunter, err := GetAllEmprunter(testDB)
	if err != nil {
		t.Fatalf("Erreur dans GetAllEmprunter : %v", err)
	}

	if len(emprunter) != 2 {
		t.Errorf("Le nombre d'emprunter ne correspond pas à celui attendu.")
	}

	expectedIDs := []int{1, 3}
	for i, emprunter := range emprunter {
		if emprunter.ID_EMPRUNT != expectedIDs[i] {
			t.Errorf("La valeur de l'ID_EMPRUNT ne correspond pas à celle attendue.")
		}
	}
}

func TestPostEmprunter(t *testing.T) {
	testDB, err := InitTestDB()

	//create TYPE_OUVRAGE
	_, err = testDB.Exec("INSERT INTO TYPE_OUVRAGE (ID_TYPE) VALUES ('type1'), ('type2'), ('type3')")
	if err != nil {
		t.Errorf("Erreur lors de l'insertion de données de test : %v", err)
	}

	//create EDITEUR
	_, err = testDB.Exec("INSERT INTO EDITEUR (ID_EDITEUR, NOM_EDITEUR, TEL_EDITEUR, MAIL_EDITEUR, ADRESSE_EDITEUR) VALUES (1, 'nom1', 'tel1', 'mail1', 'adresse1'), (2, 'nom2', 'tel2', 'mail2', 'adresse2'), (3, 'nom3', 'tel3', 'mail3', 'adresse3')")
	if err != nil {
		t.Errorf("Erreur lors de l'insertion de données de test : %v", err)
	}

	//create OUVRAGE
	_, err = testDB.Exec("INSERT INTO OUVRAGE (ID_OUVRAGE, TITRE_OUVRAGE, PARUTION_OUVRAGE, ID_TYPE, ID_EDITEUR) VALUES (1, 'titre1', '2023-02-01', 'type1', 1), (2, 'titre2', '2023-02-01', 'type2', 2), (3, 'titre3', '2023-02-01', 'type3', 3)")
	if err != nil {
		t.Errorf("Erreur lors de l'insertion de données de test : %v", err)
	}

	//create CATEGORY_ABONNE
	_, err = testDB.Exec("INSERT INTO CATEGORIE_ABONNE (ID_CATEGORIE) VALUES ('1'), ('2'), ('3')")
	if err != nil {
		t.Errorf("Erreur lors de l'insertion de données de test : %v", err)
	}

	//create ABONNE
	_, err = testDB.Exec("INSERT INTO ABONNE (ID_ABONNE, NOM_ABONNE, PRENOM_ABONNE, TEL_ABONNE, MAIL_ABONNE, ADRESSE_ABONNE, ID_CATEGORIE) VALUES (1, 'nom1', 'prenom1', 'tel1', 'mail1', 'adresse1', '1'), (2, 'nom2', 'prenom2', 'tel2', 'mail2', 'adresse2', '2'), (3, 'nom3', 'prenom3', 'tel3', 'mail3', 'adresse3', '3')")
	if err != nil {
		t.Errorf("Erreur lors de l'insertion de données de test : %v", err)
	}

	emprunter := EMPRUNTER{
		ID_OUVRAGE:  2,
		ID_ABONNE:   3,
		DATE_SORTIE: "2023-02-01",
	}

	err = PostEmprunter(testDB, emprunter)
	if err != nil {
		t.Fatalf("Erreur dans PostEmprunter : %v", err)
	}

	emprunter, err = GetEmprunterById(testDB, 4)
	if err != nil {
		t.Fatalf("Erreur dans GetEmprunterById : %v", err)
	}

}

func TestPutEmprunter(t *testing.T) {
	testDB, err := InitTestDB()

	err = CreateTestEmprunter(testDB)
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	emprunter := EMPRUNTER{
		ID_EMPRUNT:  2,
		ID_OUVRAGE:  2,
		ID_ABONNE:   2,
		DATE_SORTIE: "2023-02-02",
	}

	err = PutEmprunter(testDB, emprunter)
	if err != nil {
		t.Fatalf("Erreur dans PutEmprunter : %v", err)
	}

	emprunter, err = GetEmprunterById(testDB, 2)
	if err != nil {
		t.Fatalf("Erreur dans GetEmprunterById : %v", err)
	}

	if emprunter.ID_OUVRAGE != 2 {
		t.Errorf("La valeur de l'ID_OUVRAGE ne correspond pas à celle attendue.")
	}
}

func TestRetourEmprunter(t *testing.T) {
	testDB, err := InitTestDB()

	err = CreateTestEmprunter(testDB)
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	err = RetourEmprunter(testDB, 2)
	if err != nil {
		t.Fatalf("Erreur dans RetourEmprunter : %v", err)
	}

}
