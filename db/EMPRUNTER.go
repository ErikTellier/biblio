package db

import "database/sql"

// Create a new EMPRUNTER record.
func CreateEmprunter(db *sql.DB, e Emprunter) error {
	_, err := db.Exec("INSERT INTO EMPRUNTER (ID_ABONNE, ID_OUVRAGE, DATE_SORTIE, DATE_RETOUR) VALUES (?, ?, ?, ?)", e.idAbonne, e.idOuvrage, e.dateSortie, e.dateRetour)
	return err
}

// Read all EMPRUNTER records and return slices of Emprunter.
func ReadEmprunter(db *sql.DB) ([]Emprunter, error) {
	rows, err := db.Query("SELECT ID_EMPRUNT, ID_ABONNE, ID_OUVRAGE, DATE_SORTIE, DATE_RETOUR FROM EMPRUNTER")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var emprunts []Emprunter

	for rows.Next() {
		var e Emprunter
		if err := rows.Scan(&e.idEmprunt, &e.idAbonne, &e.idOuvrage, &e.dateSortie, &e.dateRetour); err != nil {
			return nil, err
		}
		emprunts = append(emprunts, e)
	}

	return emprunts, nil
}

// Get an EMPRUNTER by its ID.
func GetEmprunter(db *sql.DB, idEmprunt int) (Emprunter, error) {
	var e Emprunter
	err := db.QueryRow("SELECT ID_ABONNE, ID_OUVRAGE, DATE_SORTIE, DATE_RETOUR FROM EMPRUNTER WHERE ID_EMPRUNT = ?", idEmprunt).
		Scan(&e.idAbonne, &e.idOuvrage, &e.dateSortie, &e.dateRetour)
	e.idEmprunt = idEmprunt
	return e, err
}

// Update an existing EMPRUNTER record.
func UpdateEmprunter(db *sql.DB, e Emprunter) error {
	_, err := db.Exec("UPDATE EMPRUNTER SET ID_ABONNE = ?, ID_OUVRAGE = ?, DATE_SORTIE = ?, DATE_RETOUR = ? WHERE ID_EMPRUNT = ?",
		e.idAbonne, e.idOuvrage, e.dateSortie, e.dateRetour, e.idEmprunt)
	return err
}

// Delete an EMPRUNTER record by ID.
func DeleteEmprunter(db *sql.DB, idEmprunt int) error {
	_, err := db.Exec("DELETE FROM EMPRUNTER WHERE ID_EMPRUNT = ?", idEmprunt)
	return err
}
