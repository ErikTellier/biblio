package db

import "database/sql"

// Create a new OUVRAGE record.
func CreateOuvrage(db *sql.DB, o Ouvrage) error {
	_, err := db.Exec("INSERT INTO OUVRAGE (TITRE_OUVRAGE, PARUTION_OUVRAGE, ID_TYPE, ID_EDITEUR) VALUES (?, ?, ?, ?)",
		o.titreOuvrage, o.parutionOuvrage, o.idType, o.idEditeur)
	return err
}

// Read all OUVRAGE records and return slices of Ouvrage.
func ReadOuvrage(db *sql.DB) ([]Ouvrage, error) {
	rows, err := db.Query("SELECT ID_OUVRAGE, TITRE_OUVRAGE, PARUTION_OUVRAGE, ID_TYPE, ID_EDITEUR FROM OUVRAGE")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ouvrages []Ouvrage

	for rows.Next() {
		var o Ouvrage
		if err := rows.Scan(&o.idOuvrage, &o.titreOuvrage, &o.parutionOuvrage, &o.idType, &o.idEditeur); err != nil {
			return nil, err
		}
		ouvrages = append(ouvrages, o)
	}

	return ouvrages, nil
}

// Get an OUVRAGE by its ID.
func GetOuvrage(db *sql.DB, idOuvrage int) (Ouvrage, error) {
	var o Ouvrage
	err := db.QueryRow("SELECT TITRE_OUVRAGE, PARUTION_OUVRAGE, ID_TYPE, ID_EDITEUR FROM OUVRAGE WHERE ID_OUVRAGE = ?", idOuvrage).
		Scan(&o.titreOuvrage, &o.parutionOuvrage, &o.idType, &o.idEditeur)
	o.idOuvrage = idOuvrage
	return o, err
}

// Update an existing OUVRAGE record.
func UpdateOuvrage(db *sql.DB, o Ouvrage) error {
	_, err := db.Exec("UPDATE OUVRAGE SET TITRE_OUVRAGE = ?, PARUTION_OUVRAGE = ?, ID_TYPE = ?, ID_EDITEUR = ? WHERE ID_OUVRAGE = ?",
		o.titreOuvrage, o.parutionOuvrage, o.idType, o.idEditeur, o.idOuvrage)
	return err
}

// Delete an OUVRAGE record by ID.
func DeleteOuvrage(db *sql.DB, idOuvrage int) error {
	_, err := db.Exec("DELETE FROM OUVRAGE WHERE ID_OUVRAGE = ?", idOuvrage)
	return err
}
