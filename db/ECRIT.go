package db

import "database/sql"

func GetAllEcrit(db *sql.DB) ([]ECRIT, error) {
	rows, err := db.Query("SELECT * FROM ECRIT")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ecrits []ECRIT
	for rows.Next() {
		var ecrit ECRIT
		err := rows.Scan(&ecrit.ID_ECRIT, &ecrit.ID_OUVRAGE, &ecrit.ID_AUTEUR)
		if err != nil {
			return nil, err
		}
		ecrits = append(ecrits, ecrit)
	}
	return ecrits, nil
}

func DeleteEcrit(db *sql.DB, id int) error {
	// Exécutez la requête SQL pour récupérer tous les abonnés
	query := "DELETE FROM ECRIT WHERE ID_ECRIT = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func PostEcrit(db *sql.DB, ecrit ECRIT) error {
	// Exécutez la requête SQL pour récupérer tous les abonnés
	query := "INSERT INTO ECRIT (ID_OUVRAGE, ID_AUTEUR) VALUES (?, ?)"
	_, err := db.Exec(query, ecrit.ID_OUVRAGE, ecrit.ID_AUTEUR)
	if err != nil {
		return err
	}
	return nil
}
