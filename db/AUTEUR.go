package db

import "database/sql"

// Create a new AUTEUR record.
func CreateAuteur(db *sql.DB, a Auteur) error {
	_, err := db.Exec("INSERT INTO AUTEUR (NOM_AUTEUR, PRENOM_AUTEUR, TEL_AUTEUR, MAIL_AUTEUR, ADRESSE_AUTEUR) VALUES (?, ?, ?, ?, ?)",
		a.nomAuteur, a.prenomAuteur, a.telAuteur, a.mailAuteur, a.adresseAuteur)
	return err
}

// Read all AUTEUR records and return slices of Auteur.
func ReadAuteur(db *sql.DB) ([]Auteur, error) {
	rows, err := db.Query("SELECT ID_AUTEUR, NOM_AUTEUR, PRENOM_AUTEUR, TEL_AUTEUR, MAIL_AUTEUR, ADRESSE_AUTEUR FROM AUTEUR")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var auteurs []Auteur

	for rows.Next() {
		var a Auteur
		if err := rows.Scan(&a.idAuteur, &a.nomAuteur, &a.prenomAuteur, &a.telAuteur, &a.mailAuteur, &a.adresseAuteur); err != nil {
			return nil, err
		}
		auteurs = append(auteurs, a)
	}

	return auteurs, nil
}

// Get an AUTEUR by its ID.
func GetAuteur(db *sql.DB, idAuteur int) (Auteur, error) {
	var a Auteur
	err := db.QueryRow("SELECT NOM_AUTEUR, PRENOM_AUTEUR, TEL_AUTEUR, MAIL_AUTEUR, ADRESSE_AUTEUR FROM AUTEUR WHERE ID_AUTEUR = ?", idAuteur).
		Scan(&a.nomAuteur, &a.prenomAuteur, &a.telAuteur, &a.mailAuteur, &a.adresseAuteur)
	a.idAuteur = idAuteur
	return a, err
}

// Update an existing AUTEUR record.
func UpdateAuteur(db *sql.DB, a Auteur) error {
	_, err := db.Exec("UPDATE AUTEUR SET NOM_AUTEUR = ?, PRENOM_AUTEUR = ?, TEL_AUTEUR = ?, MAIL_AUTEUR = ?, ADRESSE_AUTEUR = ? WHERE ID_AUTEUR = ?",
		a.nomAuteur, a.prenomAuteur, a.telAuteur, a.mailAuteur, a.adresseAuteur, a.idAuteur)
	return err
}

// Delete an AUTEUR record by ID.
func DeleteAuteur(db *sql.DB, idAuteur int) error {
	_, err := db.Exec("DELETE FROM AUTEUR WHERE ID_AUTEUR = ?", idAuteur)
	return err
}
