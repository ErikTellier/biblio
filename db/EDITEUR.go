package db

import "database/sql"

func GetAllEditeur(db *sql.DB) ([]EDITEUR, error) {
	rows, err := db.Query("SELECT * FROM EDITEUR")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var editeurs []EDITEUR
	for rows.Next() {
		var editeur EDITEUR
		err := rows.Scan(&editeur.ID_EDITEUR, &editeur.NOM_EDITEUR, &editeur.TEL_EDITEUR, &editeur.MAIL_EDITEUR, &editeur.ADRESSE_EDITEUR)
		if err != nil {
			return nil, err
		}
		editeurs = append(editeurs, editeur)
	}
	return editeurs, nil
}

func DeleteEditeur(db *sql.DB, id int) error {
	// Exécutez la requête SQL pour récupérer tous les abonnés
	query := "DELETE FROM EDITEUR WHERE ID_EDITEUR = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func PostEditeur(db *sql.DB, editeur EDITEUR) error {
	// Exécutez la requête SQL pour récupérer tous les abonnés
	query := "INSERT INTO EDITEUR (NOM_EDITEUR, TEL_EDITEUR, MAIL_EDITEUR, ADRESSE_EDITEUR) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, editeur.NOM_EDITEUR, editeur.TEL_EDITEUR, editeur.MAIL_EDITEUR, editeur.ADRESSE_EDITEUR)
	if err != nil {
		return err
	}
	return nil
}
