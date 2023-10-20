package db

import "database/sql"

func GetAllAuteur(db *sql.DB) ([]AUTEUR, error) {
	// Exécutez la requête SQL pour récupérer tous les abonnés
	query := "SELECT * FROM AUTEUR"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var auteurs []AUTEUR

	// Parcourez les lignes résultantes et mappez-les sur la structure ABONNE
	for rows.Next() {
		var auteur AUTEUR
		err := rows.Scan(&auteur.ID_AUTEUR, &auteur.NOM_AUTEUR, &auteur.PRENOM_AUTEUR, &auteur.TEL_AUTEUR, &auteur.MAIL_AUTEUR, &auteur.ADRESSE_AUTEUR)
		if err != nil {
			return nil, err
		}
		auteurs = append(auteurs, auteur)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return auteurs, nil
}

func GetAuteurById(db *sql.DB, id int) (AUTEUR, error) {
	// Exécutez la requête SQL pour récupérer tous les abonnés
	query := "SELECT * FROM AUTEUR WHERE ID_AUTEUR = ?"
	rows, err := db.Query(query, id)
	if err != nil {
		return AUTEUR{}, err
	}
	defer rows.Close()

	var auteur AUTEUR

	// Parcourez les lignes résultantes et mappez-les sur la structure ABONNE
	for rows.Next() {
		err := rows.Scan(&auteur.ID_AUTEUR, &auteur.NOM_AUTEUR, &auteur.PRENOM_AUTEUR, &auteur.TEL_AUTEUR, &auteur.MAIL_AUTEUR, &auteur.ADRESSE_AUTEUR)
		if err != nil {
			return AUTEUR{}, err
		}
	}

	if err := rows.Err(); err != nil {
		return AUTEUR{}, err
	}

	return auteur, nil
}

func DeleteAuteur(db *sql.DB, id int) error {
	// Exécutez la requête SQL pour récupérer tous les abonnés
	query := "DELETE FROM AUTEUR WHERE ID_AUTEUR = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func PostAuteur(db *sql.DB, auteur AUTEUR) error {
	// Exécutez la requête SQL pour récupérer tous les abonnés
	query := "INSERT INTO AUTEUR (NOM_AUTEUR, PRENOM_AUTEUR, TEL_AUTEUR, MAIL_AUTEUR, ADRESSE_AUTEUR) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(query, auteur.NOM_AUTEUR, auteur.PRENOM_AUTEUR, auteur.TEL_AUTEUR, auteur.MAIL_AUTEUR, auteur.ADRESSE_AUTEUR)
	if err != nil {
		return err
	}
	return nil
}

func PutAuteur(db *sql.DB, auteur AUTEUR) error {
	// Exécutez la requête SQL pour récupérer tous les abonnés
	query := "UPDATE AUTEUR SET NOM_AUTEUR = ?, PRENOM_AUTEUR = ?, TEL_AUTEUR = ?, MAIL_AUTEUR = ?, ADRESSE_AUTEUR = ? WHERE ID_AUTEUR = ?"
	_, err := db.Exec(query, auteur.NOM_AUTEUR, auteur.PRENOM_AUTEUR, auteur.TEL_AUTEUR, auteur.MAIL_AUTEUR, auteur.ADRESSE_AUTEUR, auteur.ID_AUTEUR)
	if err != nil {
		return err
	}
	return nil
}
