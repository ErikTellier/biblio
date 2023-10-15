package db

import "database/sql"

func CreateCategorieAbonne(db *sql.DB, idCategorie string) error {
	_, err := db.Exec("INSERT INTO CATEGORIE_ABONNE (ID_CATEGORIE) VALUES (?)", idCategorie)
	return err
}

func ReadCategorieAbonne(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT ID_CATEGORIE FROM CATEGORIE_ABONNE")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []string

	for rows.Next() {
		var category string
		if err := rows.Scan(&category); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func GetCategorieAbonne(db *sql.DB, idCategorie string) (string, error) {
	var categorie string
	err := db.QueryRow("SELECT ID_CATEGORIE FROM CATEGORIE_ABONNE WHERE ID_CATEGORIE = ?", idCategorie).Scan(&categorie)
	return categorie, err
}

func UpdateCategorieAbonne(db *sql.DB, idCategorie string, newIdCategorie string) error {
	_, err := db.Exec("UPDATE CATEGORIE_ABONNE SET ID_CATEGORIE = ? WHERE ID_CATEGORIE = ?", newIdCategorie, idCategorie)
	return err
}

func DeleteCategorieAbonne(db *sql.DB, idCategorie string) error {
	_, err := db.Exec("DELETE FROM CATEGORIE_ABONNE WHERE ID_CATEGORIE = ?", idCategorie)
	return err
}
