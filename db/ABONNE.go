package db

import "database/sql"

func GetAllAbonnes(db *sql.DB) ([]ABONNE, error) {
	// Exécutez la requête SQL pour récupérer tous les abonnés
	query := "SELECT * FROM ABONNE"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var abonnes []ABONNE

	// Parcourez les lignes résultantes et mappez-les sur la structure ABONNE
	for rows.Next() {
		var abonne ABONNE
		err := rows.Scan(&abonne.ID_ABONNE, &abonne.NOM_ABONNE, &abonne.PRENOM_ABONNE, &abonne.TEL_ABONNE, &abonne.MAIL_ABONNE, &abonne.ADRESSE_ABONNE, &abonne.STATUS_ABONNE, &abonne.CONFIANCE_ABONNE, &abonne.ID_CATEGORIE)
		if err != nil {
			return nil, err
		}
		abonnes = append(abonnes, abonne)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return abonnes, nil
}

func GetAbonneById(db *sql.DB, id int) (ABONNE, error) {
	// Exécutez la requête SQL pour récupérer tous les abonnés
	query := "SELECT * FROM ABONNE WHERE ID_ABONNE = ?"
	rows, err := db.Query(query, id)
	if err != nil {
		return ABONNE{}, err
	}
	defer rows.Close()

	var abonne ABONNE

	// Parcourez les lignes résultantes et mappez-les sur la structure ABONNE
	for rows.Next() {
		err := rows.Scan(&abonne.ID_ABONNE, &abonne.NOM_ABONNE, &abonne.PRENOM_ABONNE, &abonne.TEL_ABONNE, &abonne.MAIL_ABONNE, &abonne.ADRESSE_ABONNE, &abonne.STATUS_ABONNE, &abonne.CONFIANCE_ABONNE, &abonne.ID_CATEGORIE)
		if err != nil {
			return ABONNE{}, err
		}
	}

	if err := rows.Err(); err != nil {
		return ABONNE{}, err
	}

	return abonne, nil
}

func DeleteAbonne(db *sql.DB, id int) error {
	// Exécutez la requête SQL pour récupérer tous les abonnés
	query := "DELETE FROM ABONNE WHERE ID_ABONNE = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func PostAbonne(db *sql.DB, abonne ABONNE) error {
	// Exécutez la requête SQL pour récupérer tous les abonnés
	query := "INSERT INTO ABONNE (NOM_ABONNE, PRENOM_ABONNE, TEL_ABONNE, MAIL_ABONNE, ADRESSE_ABONNE, ID_CATEGORIE) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(query, abonne.NOM_ABONNE, abonne.PRENOM_ABONNE, abonne.TEL_ABONNE, abonne.MAIL_ABONNE, abonne.ADRESSE_ABONNE, abonne.ID_CATEGORIE)
	if err != nil {
		return err
	}
	return nil
}

func PutAbonne(db *sql.DB, abonne ABONNE) error {
	// Exécutez la requête SQL pour récupérer tous les abonnés
	query := "UPDATE ABONNE SET NOM_ABONNE = ?, PRENOM_ABONNE = ?, TEL_ABONNE = ?, MAIL_ABONNE = ?, ADRESSE_ABONNE = ?, ID_CATEGORIE = ? WHERE ID_ABONNE = ?"
	_, err := db.Exec(query, abonne.NOM_ABONNE, abonne.PRENOM_ABONNE, abonne.TEL_ABONNE, abonne.MAIL_ABONNE, abonne.ADRESSE_ABONNE, abonne.ID_CATEGORIE, abonne.ID_ABONNE)
	if err != nil {
		return err
	}
	return nil
}
