package db

import "database/sql"

func CreateEditeur(db *sql.DB, nomEditeur, telEditeur, mailEditeur, adresseEditeur string) error {
	_, err := db.Exec("INSERT INTO EDITEUR (NOM_EDITEUR, TEL_EDITEUR, MAIL_EDITEUR, ADRESSE_EDITEUR) VALUES (?, ?, ?, ?)", nomEditeur, telEditeur, mailEditeur, adresseEditeur)
	return err
}

// Read all EDITEUR records
func ReadEditeur(db *sql.DB) ([]int, []string, []string, []string, []string, error) {
	rows, err := db.Query("SELECT ID_EDITEUR, NOM_EDITEUR, TEL_EDITEUR, MAIL_EDITEUR, ADRESSE_EDITEUR FROM EDITEUR")
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	defer rows.Close()

	var (
		ids      []int
		noms     []string
		tels     []string
		mails    []string
		adresses []string
	)

	for rows.Next() {
		var id int
		var nom, tel, mail, adresse string
		if err := rows.Scan(&id, &nom, &tel, &mail, &adresse); err != nil {
			return nil, nil, nil, nil, nil, err
		}
		ids = append(ids, id)
		noms = append(noms, nom)
		tels = append(tels, tel)
		mails = append(mails, mail)
		adresses = append(adresses, adresse)
	}

	return ids, noms, tels, mails, adresses, nil
}

// get an EDITEUR by its ID
func GetEditeur(db *sql.DB, idEditeur int) (string, string, string, string, error) {
	var nom, tel, mail, adresse string
	err := db.QueryRow("SELECT NOM_EDITEUR, TEL_EDITEUR, MAIL_EDITEUR, ADRESSE_EDITEUR FROM EDITEUR WHERE ID_EDITEUR = ?", idEditeur).Scan(&nom, &tel, &mail, &adresse)
	return nom, tel, mail, adresse, err
}

// Update an existing EDITEUR
func UpdateEditeur(db *sql.DB, idEditeur int, newNom, newTel, newMail, newAdresse string) error {
	_, err := db.Exec("UPDATE EDITEUR SET NOM_EDITEUR = ?, TEL_EDITEUR = ?, MAIL_EDITEUR = ?, ADRESSE_EDITEUR = ? WHERE ID_EDITEUR = ?", newNom, newTel, newMail, newAdresse, idEditeur)
	return err
}

// Delete an EDITEUR
func DeleteEditeur(db *sql.DB, idEditeur int) error {
	_, err := db.Exec("DELETE FROM EDITEUR WHERE ID_EDITEUR = ?", idEditeur)
	return err
}
