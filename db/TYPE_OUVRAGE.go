package db

import "database/sql"

func GetAllTypeOuvrage(db *sql.DB) ([]TYPE_OUVRAGE, error) {
	// Exécutez la requête SQL pour récupérer tous les abonnés
	query := "SELECT * FROM TYPE_OUVRAGE"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var type_ouvrages []TYPE_OUVRAGE

	// Parcourez les lignes résultantes et mappez-les sur la structure ABONNE
	for rows.Next() {
		var type_ouvrage TYPE_OUVRAGE
		err := rows.Scan(&type_ouvrage.ID_TYPE)
		if err != nil {
			return nil, err
		}
		type_ouvrages = append(type_ouvrages, type_ouvrage)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return type_ouvrages, nil
}

func GetTypeOuvrageById(db *sql.DB, id string) (TYPE_OUVRAGE, error) {
	// Exécutez la requête SQL pour récupérer tous les abonnés
	query := "SELECT * FROM TYPE_OUVRAGE WHERE ID_TYPE = ?"
	rows, err := db.Query(query, id)
	if err != nil {
		return TYPE_OUVRAGE{}, err
	}
	defer rows.Close()

	var type_ouvrage TYPE_OUVRAGE

	// Parcourez les lignes résultantes et mappez-les sur la structure ABONNE
	for rows.Next() {
		err := rows.Scan(&type_ouvrage.ID_TYPE)
		if err != nil {
			return TYPE_OUVRAGE{}, err
		}
	}

	if err := rows.Err(); err != nil {
		return TYPE_OUVRAGE{}, err
	}

	return type_ouvrage, nil
}

func DeleteTypeOuvrage(db *sql.DB, id string) error {
	// Exécutez la requête SQL pour récupérer tous les abonnés
	query := "DELETE FROM TYPE_OUVRAGE WHERE ID_TYPE = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func PostTypeOuvrage(db *sql.DB, ouvrage TYPE_OUVRAGE) error {
	// Exécutez la requête SQL pour récupérer tous les abonnés
	query := "INSERT INTO TYPE_OUVRAGE (ID_TYPE) VALUES (?)"
	_, err := db.Exec(query, ouvrage.ID_TYPE)
	if err != nil {
		return err
	}
	return nil
}
