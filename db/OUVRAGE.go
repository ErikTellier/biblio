package db

import "database/sql"

func GetAllOuvrage(db *sql.DB) ([]OUVRAGE, error) {
	rows, err := db.Query("SELECT * FROM OUVRAGE")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ouvrages []OUVRAGE
	for rows.Next() {
		var ouvrage OUVRAGE
		err := rows.Scan(&ouvrage.ID_OUVRAGE, &ouvrage.TITRE_OUVRAGE, &ouvrage.PARUTION_OUVRAGE, &ouvrage.ID_TYPE, &ouvrage.ID_EDITEUR)
		if err != nil {
			return nil, err
		}
		ouvrages = append(ouvrages, ouvrage)
	}
	return ouvrages, nil
}

func GetOuvrageById(db *sql.DB, id int) (OUVRAGE, error) {
	rows, err := db.Query("SELECT * FROM OUVRAGE WHERE ID_OUVRAGE = ?", id)
	if err != nil {
		return OUVRAGE{}, err
	}
	defer rows.Close()

	var ouvrage OUVRAGE
	for rows.Next() {
		err := rows.Scan(&ouvrage.ID_OUVRAGE, &ouvrage.TITRE_OUVRAGE, &ouvrage.PARUTION_OUVRAGE, &ouvrage.ID_TYPE, &ouvrage.ID_EDITEUR)
		if err != nil {
			return OUVRAGE{}, err
		}
	}
	return ouvrage, nil
}

func DeleteOuvrage(db *sql.DB, id int) error {
	// Exécutez la requête SQL pour récupérer tous les abonnés
	query := "DELETE FROM OUVRAGE WHERE ID_OUVRAGE = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func PostOuvrage(db *sql.DB, ouvrage OUVRAGE) error {
	// Exécutez la requête SQL pour récupérer tous les abonnés
	query := "INSERT INTO OUVRAGE ( TITRE_OUVRAGE, PARUTION_OUVRAGE, ID_TYPE, ID_EDITEUR) VALUES ( ?, ?, ?, ?)"
	_, err := db.Exec(query, ouvrage.TITRE_OUVRAGE, ouvrage.PARUTION_OUVRAGE, ouvrage.ID_TYPE, ouvrage.ID_EDITEUR)
	if err != nil {
		return err
	}
	return nil
}

func PutOuvrage(db *sql.DB, ouvrage OUVRAGE) error {
	// Exécutez la requête SQL pour récupérer tous les abonnés
	query := "UPDATE OUVRAGE SET TITRE_OUVRAGE = ?, PARUTION_OUVRAGE = ?, ID_TYPE = ?, ID_EDITEUR = ? WHERE ID_OUVRAGE = ?"
	_, err := db.Exec(query, ouvrage.TITRE_OUVRAGE, ouvrage.PARUTION_OUVRAGE, ouvrage.ID_TYPE, ouvrage.ID_EDITEUR, ouvrage.ID_OUVRAGE)
	if err != nil {
		return err
	}
	return nil
}
