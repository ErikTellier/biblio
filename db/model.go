package db

import "database/sql"

type CATEGORIE_ABONNE struct {
	ID_CATEGORIE string //PK
}

type ABONNE struct {
	ID_ABONNE        int //PK : should not be set when creating a new ABONNE
	NOM_ABONNE       string
	PRENOM_ABONNE    string
	TEL_ABONNE       string
	MAIL_ABONNE      string
	ADRESSE_ABONNE   string
	STATUS_ABONNE    bool   //should not be set when creating a new ABONNE
	CONFIANCE_ABONNE int    //should not be set when creating a new ABONNE
	ID_CATEGORIE     string //FK -> CATEGORIE_ABONNE : ID_CATEGORIE
}

type TYPE_OUVRAGE struct {
	ID_TYPE string //PK
}

type AUTEUR struct {
	ID_AUTEUR      int //PK : should not be set when creating a new AUTEUR
	NOM_AUTEUR     string
	PRENOM_AUTEUR  string
	TEL_AUTEUR     string
	MAIL_AUTEUR    string
	ADRESSE_AUTEUR string
}

type EDITEUR struct {
	ID_EDITEUR      int //PK : should not be set when creating a new EDITEUR
	NOM_EDITEUR     string
	TEL_EDITEUR     string
	MAIL_EDITEUR    string
	ADRESSE_EDITEUR string
}

type OUVRAGE struct {
	ID_OUVRAGE       int //PK : should not be set when creating a new OUVRAGE
	TITRE_OUVRAGE    string
	PARUTION_OUVRAGE string
	ID_TYPE          string //FK -> TYPE_OUVRAGE : ID_TYPE_OUVRAGE
	ID_EDITEUR       int    //FK -> EDITEUR : ID_EDITEUR
}

type ECRIT struct {
	ID_ECRIT   int //PK : should not be set when creating a new ECRIT
	ID_OUVRAGE int //FK -> OUVRAGE : ID_OUVRAGE
	ID_AUTEUR  int //FK -> AUTEUR : ID_AUTEUR
}

type EMPRUNTER struct {
	ID_EMPRUNT  int            //PK : should not be set when creating a new EMPRUNTER
	ID_OUVRAGE  int            //FK -> OUVRAGE : ID_OUVRAGE
	ID_ABONNE   int            //FK -> ABONNE : ID_ABONNE
	DATE_SORTIE string         //should not be set when creating a new EMPRUNTER
	DATE_RETOUR sql.NullString //should not be set when creating a new EMPRUNTER
}
