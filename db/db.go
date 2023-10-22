package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
)

func OpenDBConnection() (*sql.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func openTestDBConnection() (*sql.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_TEST")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func clearData(db *sql.DB) error {
	// Liste des noms de tables à vider
	tables := []string{
		"EMPRUNTER",
		"ECRIT",
		"OUVRAGE",
		"AUTEUR",
		"EDITEUR",
		"ABONNE",
		"CATEGORIE_ABONNE",
		"TYPE_OUVRAGE",
	}

	// Parcours de la liste des tables et suppression des données
	for _, table := range tables {
		_, err := db.Exec(fmt.Sprintf("DELETE FROM %s;", table))
		if err != nil {
			return err
		}
		fmt.Printf("Données de la table %s supprimées avec succès\n", table)
	}

	return nil
}

func InitTestDB() (*sql.DB, error) {
	db, err := openTestDBConnection()
	if err != nil {
		return nil, err
	}

	err = clearData(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}
