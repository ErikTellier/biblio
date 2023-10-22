package db

import (
	"database/sql"
	"testing"
)

func CreateTestAbonne(testDB *sql.DB) error {
	_, err := testDB.Exec("INSERT INTO CATEGORIE_ABONNE (ID_CATEGORIE) VALUES ('1'), ('2'), ('3')")
	if err != nil {
		return err
	}

	_, err = testDB.Exec("INSERT INTO ABONNE (ID_ABONNE, NOM_ABONNE, PRENOM_ABONNE, TEL_ABONNE, MAIL_ABONNE, ADRESSE_ABONNE, ID_CATEGORIE) VALUES (1, 'nom1', 'prenom1', 'tel1', 'mail1', 'adresse1', '1'), (2, 'nom2', 'prenom2', 'tel2', 'mail2', 'adresse2', '2'), (3, 'nom3', 'prenom3', 'tel3', 'mail3', 'adresse3', '3')")
	if err != nil {
		return err
	}

	return nil
}

func TestGetAllAbonnes(t *testing.T) {
	testDB, err := InitTestDB()

	err = CreateTestAbonne(testDB)
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	abonnes, err := GetAllAbonnes(testDB)
	if err != nil {
		t.Fatalf("Erreur dans GetAllAbonnes : %v", err)
	}

	if len(abonnes) != 3 {
		t.Errorf("Le nombre d'abonnés ne correspond pas à celui attendu.")
	}

	expectedIDs := []int{1, 2, 3}
	for i, abonne := range abonnes {
		if abonne.ID_ABONNE != expectedIDs[i] {
			t.Errorf("La valeur de l'ID_ABONNE ne correspond pas à celle attendue.")
		}
	}
}

func TestGetAbonneById(t *testing.T) {
	testDB, err := InitTestDB()

	err = CreateTestAbonne(testDB)
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	abonne, err := GetAbonneById(testDB, 1)
	if err != nil {
		t.Fatalf("Erreur dans GetAbonneById : %v", err)
	}
	if abonne.ID_ABONNE != 1 {
		t.Errorf("La valeur de l'ID_ABONNE ne correspond pas à celle attendue.")
	}
}

func TestDeleteAbonne(t *testing.T) {
	testDB, err := InitTestDB()

	err = CreateTestAbonne(testDB)
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	err = DeleteAbonne(testDB, 1)
	if err != nil {
		t.Fatalf("Erreur dans DeleteAbonne : %v", err)
	}

	abonnes, err := GetAllAbonnes(testDB)
	if err != nil {
		t.Fatalf("Erreur dans GetAllAbonnes : %v", err)
	}

	if len(abonnes) != 2 {
		t.Errorf("Le nombre d'abonnés ne correspond pas à celui attendu.")
	}

	expectedIDs := []int{2, 3}
	for i, abonne := range abonnes {
		if abonne.ID_ABONNE != expectedIDs[i] {
			t.Errorf("La valeur de l'ID_ABONNE ne correspond pas à celle attendue.")
		}
	}
}

func TestPostAbonne(t *testing.T) {
	testDB, err := InitTestDB()

	_, err = testDB.Exec("INSERT INTO CATEGORIE_ABONNE (ID_CATEGORIE) VALUES ('1'), ('2'), ('3')")
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	err = PostAbonne(testDB, ABONNE{ID_ABONNE: 1, NOM_ABONNE: "nom1", PRENOM_ABONNE: "prenom1", TEL_ABONNE: "tel1", MAIL_ABONNE: "mail1", ADRESSE_ABONNE: "adresse1", ID_CATEGORIE: "1"})
	if err != nil {
		t.Fatalf("Erreur dans PostAbonne : %v", err)
	}

	abonnes, err := GetAllAbonnes(testDB)
	if err != nil {
		t.Fatalf("Erreur dans GetAllAbonnes : %v", err)
	}

	if len(abonnes) != 1 {
		t.Errorf("Le nombre d'abonnés ne correspond pas à celui attendu.")
	}

}

func TestPutAbonne(t *testing.T) {
	testDB, err := InitTestDB()

	err = CreateTestAbonne(testDB)
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	err = PutAbonne(testDB, ABONNE{ID_ABONNE: 1, NOM_ABONNE: "nom1", PRENOM_ABONNE: "prenom1", TEL_ABONNE: "tel1", MAIL_ABONNE: "mail1", ADRESSE_ABONNE: "adresse1", ID_CATEGORIE: "1"})
	if err != nil {
		t.Fatalf("Erreur dans PutAbonne : %v", err)
	}

	abonnes, err := GetAllAbonnes(testDB)
	if err != nil {
		t.Fatalf("Erreur dans GetAllAbonnes : %v", err)
	}

	if len(abonnes) != 3 {
		t.Errorf("Le nombre d'abonnés ne correspond pas à celui attendu.")
	}

	expectedIDs := []int{1, 2, 3}
	for i, abonne := range abonnes {
		if abonne.ID_ABONNE != expectedIDs[i] {
			t.Errorf("La valeur de l'ID_ABONNE ne correspond pas à celle attendue.")
		}
	}
}
