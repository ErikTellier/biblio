package db

import "database/sql"

func GetAllCategorieAbonnes(db *sql.DB) ([]CATEGORIE_ABONNE, error) {
	// Exécutez la requête SQL pour récupérer tous les abonnés
	query := "SELECT * FROM CATEGORIE_ABONNE"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categorie_abonnes []CATEGORIE_ABONNE

	// Parcourez les lignes résultantes et mappez-les sur la structure ABONNE
	for rows.Next() {
		var categorie_abonne CATEGORIE_ABONNE
		err := rows.Scan(&categorie_abonne.ID_CATEGORIE)
		if err != nil {
			return nil, err
		}
		categorie_abonnes = append(categorie_abonnes, categorie_abonne)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categorie_abonnes, nil
}

func DeleteCategorieAbonne(db *sql.DB, id string) error {
	// Exécutez la requête SQL pour récupérer tous les abonnés
	query := "DELETE FROM CATEGORIE_ABONNE WHERE ID_CATEGORIE = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func PostCategorieAbonne(db *sql.DB, abonne CATEGORIE_ABONNE) error {
	// Exécutez la requête SQL pour récupérer tous les abonnés
	query := "INSERT INTO CATEGORIE_ABONNE (ID_CATEGORIE) VALUES (?)"
	_, err := db.Exec(query, abonne.ID_CATEGORIE)
	if err != nil {
		return err
	}
	return nil
}
