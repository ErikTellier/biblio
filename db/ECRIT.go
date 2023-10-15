package db

import "database/sql"

// Create a new ECRIT record.
func CreateEcrit(db *sql.DB, e Ecrit) error {
	_, err := db.Exec("INSERT INTO ECRIT (ID_OUVRAGE, ID_AUTEUR) VALUES (?, ?)", e.IdOuvrage, e.IdAuteur)
	return err
}

// Read all ECRIT records and return slices of Ecrit.
func ReadEcrit(db *sql.DB) ([]Ecrit, error) {
	rows, err := db.Query("SELECT ID_ECRIT, ID_OUVRAGE, ID_AUTEUR FROM ECRIT")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ecrits []Ecrit

	for rows.Next() {
		var e Ecrit
		if err := rows.Scan(&e.IdEcrit, &e.IdOuvrage, &e.IdAuteur); err != nil {
			return nil, err
		}
		ecrits = append(ecrits, e)
	}

	return ecrits, nil
}

// Get an ECRIT by its ID.
func GetEcrit(db *sql.DB, idEcrit int) (Ecrit, error) {
	var e Ecrit
	err := db.QueryRow("SELECT ID_OUVRAGE, ID_AUTEUR FROM ECRIT WHERE ID_ECRIT = ?", idEcrit).Scan(&e.IdOuvrage, &e.IdAuteur)
	e.IdEcrit = idEcrit
	return e, err
}

// Update an existing ECRIT record.
func UpdateEcrit(db *sql.DB, e Ecrit) error {
	_, err := db.Exec("UPDATE ECRIT SET ID_OUVRAGE = ?, ID_AUTEUR = ? WHERE ID_ECRIT = ?", e.IdOuvrage, e.IdAuteur, e.IdEcrit)
	return err
}

// Delete an ECRIT record by ID.
func DeleteEcrit(db *sql.DB, idEcrit int) error {
	_, err := db.Exec("DELETE FROM ECRIT WHERE ID_ECRIT = ?", idEcrit)
	return err
}
