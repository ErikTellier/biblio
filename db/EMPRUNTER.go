package db

import "database/sql"

func GetAllEmprunter(db *sql.DB) ([]EMPRUNTER, error) {
	rows, err := db.Query("SELECT * FROM EMPRUNTER")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var emprunters []EMPRUNTER
	for rows.Next() {
		var emprunter EMPRUNTER
		err := rows.Scan(&emprunter.ID_EMPRUNT, &emprunter.ID_OUVRAGE, &emprunter.ID_ABONNE, &emprunter.DATE_SORTIE, &emprunter.DATE_RETOUR)
		if err != nil {
			return nil, err
		}
		emprunters = append(emprunters, emprunter)
	}
	return emprunters, nil
}

func GetEmprunterById(db *sql.DB, id int) (EMPRUNTER, error) {
	rows, err := db.Query("SELECT * FROM EMPRUNTER WHERE ID_EMPRUNT = ?", id)
	if err != nil {
		return EMPRUNTER{}, err
	}
	defer rows.Close()

	var emprunter EMPRUNTER
	for rows.Next() {
		err := rows.Scan(&emprunter.ID_EMPRUNT, &emprunter.ID_OUVRAGE, &emprunter.ID_ABONNE, &emprunter.DATE_SORTIE, &emprunter.DATE_RETOUR)
		if err != nil {
			return EMPRUNTER{}, err
		}
	}
	return emprunter, nil
}

func DeleteEmprunter(db *sql.DB, id int) error {
	// Exécutez la requête SQL pour récupérer tous les abonnés
	query := "DELETE FROM EMPRUNTER WHERE ID_EMPRUNT = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func PostEmprunter(db *sql.DB, emprunter EMPRUNTER) error {
	// Exécutez la requête SQL pour récupérer tous les abonnés
	query := "INSERT INTO EMPRUNTER (ID_OUVRAGE, ID_ABONNE) VALUES (?, ?)"
	_, err := db.Exec(query, emprunter.ID_OUVRAGE, emprunter.ID_ABONNE)
	if err != nil {
		return err
	}
	return nil
}

func PutEmprunter(db *sql.DB, emprunter EMPRUNTER) error {
	// Exécutez la requête SQL pour récupérer tous les abonnés
	query := "UPDATE EMPRUNTER SET ID_OUVRAGE = ?, ID_ABONNE = ?, DATE_SORTIE = ?, DATE_RETOUR = ? WHERE ID_EMPRUNT = ?"
	_, err := db.Exec(query, emprunter.ID_OUVRAGE, emprunter.ID_ABONNE, emprunter.DATE_SORTIE, emprunter.DATE_RETOUR, emprunter.ID_EMPRUNT)
	if err != nil {
		return err
	}
	return nil
}

func RetourEmprunter(db *sql.DB, id int) error {
	// Exécutez la requête SQL pour récupérer tous les abonnés
	query := "UPDATE EMPRUNTER SET DATE_RETOUR = CURDATE() WHERE ID_EMPRUNT = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
