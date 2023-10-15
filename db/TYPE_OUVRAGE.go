package db

import (
	"database/sql"
)

// Create a new TYPE_OUVRAGE record.
func CreateTypeOuvrage(db *sql.DB, t TypeOuvrage) error {
	_, err := db.Exec("INSERT INTO TYPE_OUVRAGE (ID_TYPE) VALUES (?)", t.IdType)
	return err
}

// Read all TYPE_OUVRAGE records and return slices of TypeOuvrage.
func ReadTypeOuvrage(db *sql.DB) ([]TypeOuvrage, error) {
	rows, err := db.Query("SELECT ID_TYPE FROM TYPE_OUVRAGE")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var types []TypeOuvrage

	for rows.Next() {
		var t TypeOuvrage
		if err := rows.Scan(&t.IdType); err != nil {
			return nil, err
		}
		types = append(types, t)
	}

	return types, nil
}

// Get a TYPE_OUVRAGE by its ID.
func GetTypeOuvrage(db *sql.DB, idType string) (TypeOuvrage, error) {
	var t TypeOuvrage
	err := db.QueryRow("SELECT ID_TYPE FROM TYPE_OUVRAGE WHERE ID_TYPE = ?", idType).Scan(&t.IdType)
	return t, err
}

// Update an existing TYPE_OUVRAGE record.
func UpdateTypeOuvrage(db *sql.DB, t TypeOuvrage, newIdType string) error {
	_, err := db.Exec("UPDATE TYPE_OUVRAGE SET ID_TYPE = ? WHERE ID_TYPE = ?", newIdType, t.IdType)
	return err
}

// Delete a TYPE_OUVRAGE record by ID.
func DeleteTypeOuvrage(db *sql.DB, idType string) error {
	_, err := db.Exec("DELETE FROM TYPE_OUVRAGE WHERE ID_TYPE = ?", idType)
	return err
}
