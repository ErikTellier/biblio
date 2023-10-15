package db

import "database/sql"

// Create a new CATEGORIE_ABONNE record.
func CreateCategorieAbonne(db *sql.DB, c CategorieAbonne) error {
	_, err := db.Exec("INSERT INTO CATEGORIE_ABONNE (ID_CATEGORIE) VALUES (?)", c.IdCategorie)
	return err
}

// Read all CATEGORIE_ABONNE records and return slices of CategorieAbonne.
func ReadCategorieAbonne(db *sql.DB) ([]CategorieAbonne, error) {
	rows, err := db.Query("SELECT ID_CATEGORIE FROM CATEGORIE_ABONNE")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []CategorieAbonne

	for rows.Next() {
		var c CategorieAbonne
		if err := rows.Scan(&c.IdCategorie); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}

// Get a CATEGORIE_ABONNE by its ID.
func GetCategorieAbonne(db *sql.DB, idCategorie string) (CategorieAbonne, error) {
	var c CategorieAbonne
	err := db.QueryRow("SELECT ID_CATEGORIE FROM CATEGORIE_ABONNE WHERE ID_CATEGORIE = ?", idCategorie).Scan(&c.IdCategorie)
	return c, err
}

// Update an existing CATEGORIE_ABONNE record.
func UpdateCategorieAbonne(db *sql.DB, c CategorieAbonne, newIdCategorie string) error {
	_, err := db.Exec("UPDATE CATEGORIE_ABONNE SET ID_CATEGORIE = ? WHERE ID_CATEGORIE = ?", newIdCategorie, c.IdCategorie)
	return err
}

// Delete a CATEGORIE_ABONNE record by ID.
func DeleteCategorieAbonne(db *sql.DB, idCategorie string) error {
	_, err := db.Exec("DELETE FROM CATEGORIE_ABONNE WHERE ID_CATEGORIE = ?", idCategorie)
	return err
}
