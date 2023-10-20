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
