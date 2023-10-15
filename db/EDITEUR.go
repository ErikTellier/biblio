package db

import (
	"database/sql"
)

// CreateEditeur inserts a new editeur record into the database.
func CreateEditeur(db *sql.DB, e Editeur) error {
	_, err := db.Exec("INSERT INTO EDITEUR (NOM_EDITEUR, TEL_EDITEUR, MAIL_EDITEUR, ADRESSE_EDITEUR) VALUES (?, ?, ?, ?)",
		e.nomEditeur, e.telEditeur, e.mailEditeur, e.adresseEditeur)
	return err
}

// Read all EDITEUR records and return slices of editeurs.
func ReadEditeur(db *sql.DB) ([]Editeur, error) {
	rows, err := db.Query("SELECT ID_EDITEUR, NOM_EDITEUR, TEL_EDITEUR, MAIL_EDITEUR, ADRESSE_EDITEUR FROM EDITEUR")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var editeurs []Editeur

	for rows.Next() {
		var e Editeur
		if err := rows.Scan(&e.idEditeur, &e.nomEditeur, &e.telEditeur, &e.mailEditeur, &e.adresseEditeur); err != nil {
			return nil, err
		}
		editeurs = append(editeurs, e)
	}

	return editeurs, nil
}

// GetEditeur retrieves an editeur by its ID.
func GetEditeur(db *sql.DB, idEditeur int) (Editeur, error) {
	var e Editeur
	err := db.QueryRow("SELECT ID_EDITEUR, NOM_EDITEUR, TEL_EDITEUR, MAIL_EDITEUR, ADRESSE_EDITEUR FROM EDITEUR WHERE ID_EDITEUR = ?", idEditeur).
		Scan(&e.idEditeur, &e.nomEditeur, &e.telEditeur, &e.mailEditeur, &e.adresseEditeur)
	return e, err
}

// Update an existing EDITEUR record.
func UpdateEditeur(db *sql.DB, e Editeur) error {
	_, err := db.Exec("UPDATE EDITEUR SET NOM_EDITEUR = ?, TEL_EDITEUR = ?, MAIL_EDITEUR = ?, ADRESSE_EDITEUR = ? WHERE ID_EDITEUR = ?",
		e.nomEditeur, e.telEditeur, e.mailEditeur, e.adresseEditeur, e.idEditeur)
	return err
}

// Delete an EDITEUR record by ID.
func DeleteEditeur(db *sql.DB, idEditeur int) error {
	_, err := db.Exec("DELETE FROM EDITEUR WHERE ID_EDITEUR = ?", idEditeur)
	return err
}
