package db

import "database/sql"

// Create a new ABONNE record.
func CreateAbonne(db *sql.DB, a Abonne) error {
	_, err := db.Exec("INSERT INTO ABONNE (NOM_ABONNE, PRENOM_ABONNE, TEL_ABONNE, MAIL_ABONNE, ADRESSE_ABONNE, STATUS_ABONNE, CONFIANCE_ABONNE, ID_CATEGORIE) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		a.nomAbonne, a.prenomAbonne, a.telAbonne, a.mailAbonne, a.adresseAbonne, a.statusAbonne, a.confianceAbonne, a.idCategorie)
	return err
}

// Read all ABONNE records and return slices of Abonne.
func ReadAbonne(db *sql.DB) ([]Abonne, error) {
	rows, err := db.Query("SELECT ID_ABONNE, NOM_ABONNE, PRENOM_ABONNE, TEL_ABONNE, MAIL_ABONNE, ADRESSE_ABONNE, STATUS_ABONNE, CONFIANCE_ABONNE, ID_CATEGORIE FROM ABONNE")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var abonnes []Abonne

	for rows.Next() {
		var a Abonne
		if err := rows.Scan(&a.idAbonne, &a.nomAbonne, &a.prenomAbonne, &a.telAbonne, &a.mailAbonne, &a.adresseAbonne, &a.statusAbonne, &a.confianceAbonne, &a.idCategorie); err != nil {
			return nil, err
		}
		abonnes = append(abonnes, a)
	}

	return abonnes, nil
}

// Get an ABONNE by its ID.
func GetAbonne(db *sql.DB, idAbonne int) (Abonne, error) {
	var a Abonne
	err := db.QueryRow("SELECT NOM_ABONNE, PRENOM_ABONNE, TEL_ABONNE, MAIL_ABONNE, ADRESSE_ABONNE, STATUS_ABONNE, CONFIANCE_ABONNE, ID_CATEGORIE FROM ABONNE WHERE ID_ABONNE = ?", idAbonne).
		Scan(&a.nomAbonne, &a.prenomAbonne, &a.telAbonne, &a.mailAbonne, &a.adresseAbonne, &a.statusAbonne, &a.confianceAbonne, &a.idCategorie)
	a.idAbonne = idAbonne
	return a, err
}

// Update an existing ABONNE record.
func UpdateAbonne(db *sql.DB, a Abonne) error {
	_, err := db.Exec("UPDATE ABONNE SET NOM_ABONNE = ?, PRENOM_ABONNE = ?, TEL_ABONNE = ?, MAIL_ABONNE = ?, ADRESSE_ABONNE = ?, STATUS_ABONNE = ?, CONFIANCE_ABONNE = ?, ID_CATEGORIE = ? WHERE ID_ABONNE = ?",
		a.nomAbonne, a.prenomAbonne, a.telAbonne, a.mailAbonne, a.adresseAbonne, a.statusAbonne, a.confianceAbonne, a.idCategorie, a.idAbonne)
	return err
}

// Delete an ABONNE record by ID.
func DeleteAbonne(db *sql.DB, idAbonne int) error {
	_, err := db.Exec("DELETE FROM ABONNE WHERE ID_ABONNE = ?", idAbonne)
	return err
}
