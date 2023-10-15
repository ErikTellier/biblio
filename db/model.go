package db

type editeur struct {
	idEditeur      int
	nomEditeur     string
	telEditeur     string
	mailEditeur    string
	adresseEditeur string
}

type TypeOuvrage struct {
	IdType string
}

type CategorieAbonne struct {
	IdCategorie string
}
