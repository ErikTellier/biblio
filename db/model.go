package db

type Editeur struct {
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

type Auteur struct {
	idAuteur      int
	nomAuteur     string
	prenomAuteur  string
	telAuteur     string
	mailAuteur    string
	adresseAuteur string
}

type Ouvrage struct {
	idOuvrage       int
	titreOuvrage    string
	parutionOuvrage string
	idType          string
	idEditeur       int
}

type Abonne struct {
	idAbonne        int
	nomAbonne       string
	prenomAbonne    string
	telAbonne       string
	mailAbonne      string
	adresseAbonne   string
	statusAbonne    bool
	confianceAbonne int
	idCategorie     string
}
