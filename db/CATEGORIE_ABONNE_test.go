package db

import (
	"database/sql"
	"testing"
)

func CreateTestCategoriesAbonnes(testDB *sql.DB) error {
	_, err := testDB.Exec("INSERT INTO CATEGORIE_ABONNE (ID_CATEGORIE) VALUES ('1'), ('2'), ('3')")
	if err != nil {
		return err
	}

	return nil
}

func TestGetAllCategorieAbonnes(t *testing.T) {
	testDB, err := InitTestDB()

	err = CreateTestCategoriesAbonnes(testDB)
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	categorieAbonnes, err := GetAllCategorieAbonnes(testDB)
	if err != nil {
		t.Fatalf("Erreur dans GetAllCategorieAbonnes : %v", err)
	}

	if len(categorieAbonnes) != 3 {
		t.Errorf("Le nombre de catégorie d'abonné retourné est incorrect.")
	}

	expectedID := []string{"1", "2", "3"}
	for i, categorieAbonne := range categorieAbonnes {
		if categorieAbonne.ID_CATEGORIE != expectedID[i] {
			t.Errorf("La valeur de l'ID_CATEGORIE ne correspond pas à celle attendue.")
		}
	}
}
func TestGetCategorieAbonneById(t *testing.T) {
	testDB, err := InitTestDB()

	err = CreateTestCategoriesAbonnes(testDB)
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	categorieAbonne, err := GetCategorieAbonneById(testDB, "1")
	if err != nil {
		t.Fatalf("Erreur dans GetCategorieAbonneById : %v", err)
	}

	if categorieAbonne.ID_CATEGORIE != "1" {
		t.Errorf("La valeur de l'ID_CATEGORIE ne correspond pas à celle attendue.")
	}

}

func TestDeleteCategorieAbonne(t *testing.T) {
	testDB, err := InitTestDB()

	err = CreateTestCategoriesAbonnes(testDB)
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de données de test : %v", err)
	}

	err = DeleteCategorieAbonne(testDB, "1")
	if err != nil {
		t.Fatalf("Erreur dans DeleteCategorieAbonne : %v", err)
	}

	categorieAbonnes, err := GetAllCategorieAbonnes(testDB)
	if err != nil {
		t.Fatalf("Erreur dans GetAllCategorieAbonnes : %v", err)
	}

	if len(categorieAbonnes) != 2 {
		t.Errorf("Le nombre de catégorie d'abonné retourné est incorrect.")
	}

	expectedID := []string{"2", "3"}
	for i, categorieAbonne := range categorieAbonnes {
		if categorieAbonne.ID_CATEGORIE != expectedID[i] {
			t.Errorf("La valeur de l'ID_CATEGORIE ne correspond pas à celle attendue.")
		}
	}
}

func TestPostCategorieAbonne(t *testing.T) {
	testDB, err := InitTestDB()

	categorieAbonne := CATEGORIE_ABONNE{
		ID_CATEGORIE: "1",
	}

	err = PostCategorieAbonne(testDB, categorieAbonne)
	if err != nil {
		t.Fatalf("Erreur dans PostCategorieAbonne : %v", err)
	}

	categorieAbonnes, err := GetAllCategorieAbonnes(testDB)
	if err != nil {
		t.Fatalf("Erreur dans GetAllCategorieAbonnes : %v", err)
	}

	if len(categorieAbonnes) != 1 {
		t.Errorf("Le nombre de catégorie d'abonné retourné est incorrect.")
	}

	if categorieAbonnes[0].ID_CATEGORIE != "1" {
		t.Errorf("La valeur de l'ID_CATEGORIE ne correspond pas à celle attendue.")
	}
}
