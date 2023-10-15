package db

import "database/sql"

// Create a new TYPE_OUVRAGE
func CreateTypeOuvrage(db *sql.DB, idType string) error {
	_, err := db.Exec("INSERT INTO TYPE_OUVRAGE (ID_TYPE) VALUES (?)", idType)
	return err
}

// Read all TYPE_OUVRAGE records
func ReadTypeOuvrage(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT ID_TYPE FROM TYPE_OUVRAGE")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var types []string

	for rows.Next() {
		var type_ouvrage string
		if err := rows.Scan(&type_ouvrage); err != nil {
			return nil, err
		}
		types = append(types, type_ouvrage)
	}

	return types, nil
}

// Update an existing TYPE_OUVRAGE
func UpdateTypeOuvrage(db *sql.DB, idType, newIdType string) error {
	_, err := db.Exec("UPDATE TYPE_OUVRAGE SET ID_TYPE = ? WHERE ID_TYPE = ?", newIdType, idType)
	return err
}

// Delete a TYPE_OUVRAGE
func DeleteTypeOuvrage(db *sql.DB, idType string) error {
	_, err := db.Exec("DELETE FROM TYPE_OUVRAGE WHERE ID_TYPE = ?", idType)
	return err
}
