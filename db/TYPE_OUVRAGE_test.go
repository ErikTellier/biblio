package db

import (
	"database/sql"
	"testing"
)

func CreateTestTypeOuvrage(testDB *sql.DB) error {
	_, err := testDB.Exec("INSERT INTO TYPE_OUVRAGE (ID_TYPE) VALUES ('1'), ('2'), ('3')")
	if err != nil {
		return err
	}

	return nil
}

func TestGetAllTypeOuvrage(t *testing.T) {
	testDB, err := InitTestDB()

	err = CreateTestTypeOuvrage(testDB)
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	typeOuvrages, err := GetAllTypeOuvrage(testDB)
	if err != nil {
		t.Fatalf("Erreur dans GetAllTypeOuvrage : %v", err)
	}

	if len(typeOuvrages) != 3 {
		t.Errorf("Le nombre d'enregistrements renvoyés n'est pas conforme.")
	}

	expectedIDs := []string{"1", "2", "3"}
	for i, ouvrage := range typeOuvrages {
		if ouvrage.ID_TYPE != expectedIDs[i] {
			t.Errorf("La valeur de l'ID_TYPE ne correspond pas à celle attendue.")
		}
	}

}

func TestGetTypeOuvrageById(t *testing.T) {
	testDB, err := InitTestDB()

	err = CreateTestTypeOuvrage(testDB)
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	typeOuvrage, err := GetTypeOuvrageById(testDB, "1")
	if err != nil {
		t.Fatalf("Erreur dans GetTypeOuvrageById : %v", err)
	}

	if typeOuvrage.ID_TYPE != "1" {
		t.Errorf("La valeur de l'ID_TYPE ne correspond pas à celle attendue.")
	}
}

func TestDeleteTypeOuvrage(t *testing.T) {
	testDB, err := InitTestDB()

	err = CreateTestTypeOuvrage(testDB)
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	err = DeleteTypeOuvrage(testDB, "1")
	if err != nil {
		t.Fatalf("Erreur dans DeleteTypeOuvrage : %v", err)
	}

	typeOuvrages, err := GetAllTypeOuvrage(testDB)
	if err != nil {
		t.Fatalf("Erreur dans GetAllTypeOuvrage : %v", err)
	}

	if len(typeOuvrages) != 2 {
		t.Errorf("Le nombre d'enregistrements renvoyés n'est pas conforme.")
	}

	expectedIDs := []string{"2", "3"}
	for i, ouvrage := range typeOuvrages {
		if ouvrage.ID_TYPE != expectedIDs[i] {
			t.Errorf("La valeur de l'ID_TYPE ne correspond pas à celle attendue.")
		}
	}
}

func TestPostTypeOuvrage(t *testing.T) {
	testDB, err := InitTestDB()

	typeOuvrage := TYPE_OUVRAGE{
		ID_TYPE: "1",
	}

	err = PostTypeOuvrage(testDB, typeOuvrage)
	if err != nil {
		t.Fatalf("Erreur dans PostTypeOuvrage : %v", err)
	}

	typeOuvrages, err := GetAllTypeOuvrage(testDB)
	if err != nil {
		t.Fatalf("Erreur dans GetAllTypeOuvrage : %v", err)
	}

	if len(typeOuvrages) != 1 {
		t.Errorf("Le nombre d'enregistrements renvoyés n'est pas conforme.")
	}

	if typeOuvrages[0].ID_TYPE != "1" {
		t.Errorf("La valeur de l'ID_TYPE ne correspond pas à celle attendue.")
	}
}
