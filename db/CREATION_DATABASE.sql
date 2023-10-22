#------------------------------------------------------------
#        Script MySQL.
#------------------------------------------------------------

DROP DATABASE biblio_test;
CREATE DATABASE biblio_test;
USE biblio_test;


#------------------------------------------------------------
# Table: CATEGORIE_ABONNE
#------------------------------------------------------------

CREATE TABLE CATEGORIE_ABONNE(
        ID_CATEGORIE Varchar (50) NOT NULL
	,CONSTRAINT CATEGORIE_ABONNE_PK PRIMARY KEY (ID_CATEGORIE)
)ENGINE=InnoDB;


#------------------------------------------------------------
# Table: ABONNE
#------------------------------------------------------------

CREATE TABLE ABONNE(
        ID_ABONNE        Int  Auto_increment  NOT NULL ,
        NOM_ABONNE       Varchar (50) NOT NULL ,
        PRENOM_ABONNE    Varchar (50) NOT NULL ,
        TEL_ABONNE       Varchar (50) NOT NULL ,
        MAIL_ABONNE      Varchar (50) NOT NULL ,
        ADRESSE_ABONNE   Varchar (50) NOT NULL ,
        STATUS_ABONNE    Bool NOT NULL ,
        CONFIANCE_ABONNE Int NOT NULL ,
        ID_CATEGORIE     Varchar (50) NOT NULL
	,CONSTRAINT ABONNE_PK PRIMARY KEY (ID_ABONNE)

	,CONSTRAINT ABONNE_CATEGORIE_ABONNE_FK FOREIGN KEY (ID_CATEGORIE) REFERENCES CATEGORIE_ABONNE(ID_CATEGORIE)
)ENGINE=InnoDB;


#------------------------------------------------------------
# Table: TYPE_OUVRAGE
#------------------------------------------------------------

CREATE TABLE TYPE_OUVRAGE(
        ID_TYPE Varchar (50) NOT NULL
	,CONSTRAINT TYPE_OUVRAGE_PK PRIMARY KEY (ID_TYPE)
)ENGINE=InnoDB;


#------------------------------------------------------------
# Table: AUTEUR
#------------------------------------------------------------

CREATE TABLE AUTEUR(
        ID_AUTEUR      Int  Auto_increment  NOT NULL ,
        NOM_AUTEUR     Varchar (50) NOT NULL ,
        PRENOM_AUTEUR  Varchar (50) NOT NULL ,
        TEL_AUTEUR     Varchar (50) NOT NULL ,
        MAIL_AUTEUR    Varchar (50) NOT NULL ,
        ADRESSE_AUTEUR Varchar (50) NOT NULL
	,CONSTRAINT AUTEUR_PK PRIMARY KEY (ID_AUTEUR)
)ENGINE=InnoDB;


#------------------------------------------------------------
# Table: EDITEUR
#------------------------------------------------------------

CREATE TABLE EDITEUR(
        ID_EDITEUR      Int  Auto_increment  NOT NULL ,
        NOM_EDITEUR     Varchar (50) NOT NULL ,
        TEL_EDITEUR     Varchar (50) NOT NULL ,
        MAIL_EDITEUR    Varchar (50) NOT NULL ,
        ADRESSE_EDITEUR Varchar (50) NOT NULL
	,CONSTRAINT EDITEUR_PK PRIMARY KEY (ID_EDITEUR)
)ENGINE=InnoDB;


#------------------------------------------------------------
# Table: OUVRAGE
#------------------------------------------------------------

CREATE TABLE OUVRAGE(
        ID_OUVRAGE       Int  Auto_increment  NOT NULL ,
        TITRE_OUVRAGE    Varchar (50) NOT NULL ,
        PARUTION_OUVRAGE Date NOT NULL ,
        ID_TYPE          Varchar (50) NOT NULL ,
        ID_EDITEUR       Int NOT NULL
	,CONSTRAINT OUVRAGE_PK PRIMARY KEY (ID_OUVRAGE)

	,CONSTRAINT OUVRAGE_TYPE_OUVRAGE_FK FOREIGN KEY (ID_TYPE) REFERENCES TYPE_OUVRAGE(ID_TYPE)
	,CONSTRAINT OUVRAGE_EDITEUR0_FK FOREIGN KEY (ID_EDITEUR) REFERENCES EDITEUR(ID_EDITEUR) ON DELETE CASCADE
)ENGINE=InnoDB;


#------------------------------------------------------------
# Table: ECRIT
#------------------------------------------------------------

CREATE TABLE ECRIT(
        ID_ECRIT  Int  Auto_increment  NOT NULL ,
        ID_OUVRAGE Int NOT NULL ,
        ID_AUTEUR  Int NOT NULL
	,CONSTRAINT ECRIT_PK PRIMARY KEY (ID_ECRIT)

	,CONSTRAINT ECRIT_OUVRAGE_FK FOREIGN KEY (ID_OUVRAGE) REFERENCES OUVRAGE(ID_OUVRAGE) ON DELETE CASCADE
	,CONSTRAINT ECRIT_AUTEUR0_FK FOREIGN KEY (ID_AUTEUR) REFERENCES AUTEUR(ID_AUTEUR) ON DELETE CASCADE
)ENGINE=InnoDB;


#------------------------------------------------------------
# Table: EMPRUNTER
#------------------------------------------------------------

CREATE TABLE EMPRUNTER(
        ID_EMPRUNT  Int  Auto_increment  NOT NULL ,
        ID_ABONNE   Int NOT NULL ,
        ID_OUVRAGE  Int NOT NULL ,
        DATE_SORTIE Date NOT NULL ,
        DATE_RETOUR Date
	,CONSTRAINT EMPRUNTER_PK PRIMARY KEY (ID_EMPRUNT)

	,CONSTRAINT EMPRUNTER_ABONNE_FK FOREIGN KEY (ID_ABONNE) REFERENCES ABONNE(ID_ABONNE) ON DELETE CASCADE
	,CONSTRAINT EMPRUNTER_OUVRAGE0_FK FOREIGN KEY (ID_OUVRAGE) REFERENCES OUVRAGE(ID_OUVRAGE) ON DELETE CASCADE
)ENGINE=InnoDB;



#------------------------------------------------------------
# Procedure: LISTE_ABONNES_PAR_CONFIANCE
#------------------------------------------------------------

DELIMITER //

CREATE PROCEDURE LISTE_ABONNES_PAR_CONFIANCE(
    IN niveau_confiance INT
)
BEGIN
    -- Sélectionner les abonnés en fonction du niveau de confiance
    SELECT ID_ABONNE, NOM_ABONNE, PRENOM_ABONNE
    FROM ABONNE
    WHERE CONFIANCE_ABONNE = niveau_confiance;
END //

DELIMITER ;

#------------------------------------------------------------
# Procedure: LIVRES_DISPONIBLE
#------------------------------------------------------------

DELIMITER //

CREATE PROCEDURE LIVRES_DISPONIBLE(
    IN date_instant DATE
)
BEGIN
    -- Utiliser la date actuelle si date_instant est vide
    IF date_instant IS NULL THEN
        SET date_instant = CURDATE();
    END IF;

    -- Sélectionner les livres disponibles à la date spécifiée
    SELECT o.ID_OUVRAGE, o.TITRE_OUVRAGE
    FROM OUVRAGE o
    LEFT JOIN (
        SELECT DISTINCT ID_OUVRAGE
        FROM EMPRUNTER
        WHERE DATE_SORTIE <= date_instant
          AND DATE_RETOUR >= date_instant
    ) AS LivresEmpruntes ON o.ID_OUVRAGE = LivresEmpruntes.ID_OUVRAGE
    WHERE LivresEmpruntes.ID_OUVRAGE IS NULL;
END //

DELIMITER ;

#------------------------------------------------------------
# Procedure: LIVRES_PLUS_EMPRUNTE
#------------------------------------------------------------

DELIMITER //

CREATE PROCEDURE LIVRES_PLUS_EMPRUNTE(
    IN annee INT
)
BEGIN
    -- Utiliser la date de début et de fin de l'année spécifiée
    SET @date_debut = STR_TO_DATE(CONCAT(annee, '-01-01'), '%Y-%m-%d');
    SET @date_fin = STR_TO_DATE(CONCAT(annee, '-12-31'), '%Y-%m-%d');

    -- Sélectionner les livres les plus empruntés pour l'année spécifiée
    SELECT o.ID_OUVRAGE, o.TITRE_OUVRAGE, COUNT(e.ID_EMPRUNT) AS NombreEmprunts
    FROM OUVRAGE o
    JOIN EMPRUNTER e ON o.ID_OUVRAGE = e.ID_OUVRAGE
    WHERE e.DATE_SORTIE BETWEEN @date_debut AND @date_fin
    GROUP BY o.ID_OUVRAGE
    ORDER BY NombreEmprunts DESC;
END //

DELIMITER ;

#------------------------------------------------------------
# Procedure: SELECT_RETARD
#------------------------------------------------------------

DELIMITER //

CREATE PROCEDURE SELECT_RETARD()
BEGIN
    SELECT e.ID_ABONNE, a.NOM_ABONNE, a.PRENOM_ABONNE, e.ID_OUVRAGE, o.TITRE_OUVRAGE, e.DATE_RETOUR
    FROM EMPRUNTER e
    INNER JOIN ABONNE a ON e.ID_ABONNE = a.ID_ABONNE
    INNER JOIN OUVRAGE o ON e.ID_OUVRAGE = o.ID_OUVRAGE
    WHERE e.DATE_RETOUR IS NULL
          AND e.DATE_SORTIE < DATE_SUB(CURDATE(), INTERVAL 15 DAY);
END //

DELIMITER ;

#------------------------------------------------------------
# Procedure: INSERER_ABONNE_ALEATOIRE
#------------------------------------------------------------

DELIMITER //

CREATE PROCEDURE INSERER_ABONNE_ALEATOIRE(
    IN NomAbonne VARCHAR(50),
    IN PrenomAbonne VARCHAR(50),
    IN TelAbonne VARCHAR(50),
    IN MailAbonne VARCHAR(50),
    IN AdresseAbonne VARCHAR(50)
)
BEGIN
    -- Sélectionner un ID de catégorie aléatoire parmi les catégories existantes
    DECLARE RandomCategoryID VARCHAR(50);

    SELECT ID_CATEGORIE INTO RandomCategoryID
    FROM CATEGORIE_ABONNE
    ORDER BY RAND() LIMIT 1;

    -- Insérer l'abonné avec l'ID de catégorie aléatoire
    INSERT INTO ABONNE (NOM_ABONNE, PRENOM_ABONNE, TEL_ABONNE, MAIL_ABONNE, ADRESSE_ABONNE, STATUS_ABONNE, CONFIANCE_ABONNE, ID_CATEGORIE)
    VALUES (NomAbonne, PrenomAbonne, TelAbonne, MailAbonne, AdresseAbonne, TRUE, 0, RandomCategoryID);
END //

DELIMITER ;

#------------------------------------------------------------
# Procedure: INSERER_OUVRAGE_ALEATOIRE
#------------------------------------------------------------

DELIMITER //

CREATE PROCEDURE INSERER_OUVRAGE_ALEATOIRE(
    IN TitreOuvrage VARCHAR(50)
)
BEGIN
    -- Variables pour les valeurs aléatoires
    DECLARE RandomEditeurID INT;
    DECLARE RandomTypeID VARCHAR(50);
    DECLARE RandomDate DATE;

    -- Sélectionner un éditeur aléatoire
    SELECT ID_EDITEUR INTO RandomEditeurID
    FROM EDITEUR
    ORDER BY RAND() LIMIT 1;

    -- Sélectionner un type d'ouvrage aléatoire
    SELECT ID_TYPE INTO RandomTypeID
    FROM TYPE_OUVRAGE
    ORDER BY RAND() LIMIT 1;

    -- Générer une date aléatoire dans une plage
    SET RandomDate = DATE_ADD('2020-01-01', INTERVAL FLOOR(RAND() * 1096) DAY);

    -- Insérer l'ouvrage avec le titre et les valeurs aléatoires
    INSERT INTO OUVRAGE (TITRE_OUVRAGE, PARUTION_OUVRAGE, ID_TYPE, ID_EDITEUR)
    VALUES (TitreOuvrage, RandomDate, RandomTypeID, RandomEditeurID);
END //

DELIMITER ;

#------------------------------------------------------------
# Déclencheur (Trigger): ENREGISTRER_DATE_SORTIE
#------------------------------------------------------------

DELIMITER //

CREATE TRIGGER ENREGISTRER_DATE_SORTIE BEFORE INSERT ON EMPRUNTER
FOR EACH ROW
BEGIN
    SET NEW.DATE_SORTIE = CURDATE();
END //

DELIMITER ;

#------------------------------------------------------------
# Déclencheur (Trigger): RENDU_RETARD
#------------------------------------------------------------

DELIMITER //

CREATE TRIGGER RENDU_RETARD AFTER UPDATE ON EMPRUNTER
FOR EACH ROW
BEGIN
    DECLARE date_emprunt DATE;
    DECLARE date_retour_prevue DATE;
    DECLARE jours_retard INT;
    
    -- Obtenir la date d'emprunt et la date de retour prévue
    SELECT DATE_SORTIE, DATE_RETOUR
    INTO date_emprunt, date_retour_prevue
    FROM EMPRUNTER
    WHERE ID_EMPRUNT = NEW.ID_EMPRUNT;
    
    -- Calculer la différence en jours entre la date de retour prévue et la date de retour effective
    SET jours_retard = DATEDIFF(date_retour_prevue, NEW.DATE_RETOUR);
    
    -- Augmenter le niveau de confiance de l'abonné si le retard est de 15 jours ou plus
    IF jours_retard >= 15 THEN
        UPDATE ABONNE
        SET CONFIANCE_ABONNE = CONFIANCE_ABONNE + 1
        WHERE ID_ABONNE = NEW.ID_ABONNE;
    END IF;
END //

DELIMITER ;

#------------------------------------------------------------
# Déclencheur (Trigger): RENDU_RETARD_DEL
#------------------------------------------------------------

DELIMITER //

CREATE TRIGGER RENDU_RETARD_DEL AFTER DELETE ON EMPRUNTER
FOR EACH ROW
BEGIN
    DECLARE date_emprunt DATE;
    DECLARE date_retour_prevue DATE;
    DECLARE jours_retard INT;
    
    -- Obtenir la date d'emprunt et la date de retour prévue du livre supprimé
    SELECT DATE_SORTIE, DATE_RETOUR
    INTO date_emprunt, date_retour_prevue
    FROM EMPRUNTER
    WHERE ID_EMPRUNT = OLD.ID_EMPRUNT;
    
    -- Calculer la différence en jours entre la date de retour prévue et la date de retour effective
    SET jours_retard = DATEDIFF(date_retour_prevue, OLD.DATE_RETOUR);
    
    -- Si le retard est de 15 jours ou plus, réduire le niveau de confiance de 1 (mais pas en dessous de 0)
    IF jours_retard >= 15 THEN
        UPDATE ABONNE
        SET CONFIANCE_ABONNE = GREATEST(CONFIANCE_ABONNE - 1, 0)
        WHERE ID_ABONNE = OLD.ID_ABONNE;
    END IF;
END //

DELIMITER ;

#------------------------------------------------------------
# Déclencheur (Trigger): DEFAULT_STATUS
#------------------------------------------------------------

DELIMITER //

CREATE TRIGGER DEFAULT_STATUS BEFORE INSERT ON ABONNE
FOR EACH ROW
BEGIN
    -- Définir les valeurs par défaut
    SET NEW.STATUS_ABONNE = TRUE;
    SET NEW.CONFIANCE_ABONNE = 0;
END //

DELIMITER ;

#------------------------------------------------------------
# Déclencheur (Trigger): CHECK_CONFIANCE
#------------------------------------------------------------

DELIMITER //

CREATE TRIGGER CHECK_CONFIANCE BEFORE UPDATE ON ABONNE
FOR EACH ROW
BEGIN
    -- Vérifier si le niveau de confiance atteint 3
    IF NEW.CONFIANCE_ABONNE >= 3 THEN
        SET NEW.STATUS_ABONNE = FALSE;
    END IF;
END //

DELIMITER ;

#------------------------------------------------------------
# Déclencheur (Trigger): VERIFIER_DISPONIBILITE_OUVRAGE
#------------------------------------------------------------

DELIMITER //

CREATE TRIGGER VERIFIER_DISPONIBILITE_OUVRAGE
BEFORE INSERT ON EMPRUNTER
FOR EACH ROW
BEGIN
    -- Vérifier si l'ouvrage est déjà emprunté
    IF EXISTS (
        SELECT 1
        FROM EMPRUNTER
        WHERE ID_OUVRAGE = NEW.ID_OUVRAGE
          AND DATE_RETOUR IS NULL
    ) THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Cet ouvrage est déjà emprunté.';
    END IF;
END //

DELIMITER ;

#------------------------------------------------------------
# Déclencheur (Trigger): VERIFIER_EMPRUNT_ABONNE
#------------------------------------------------------------

DELIMITER //

CREATE TRIGGER VERIFIER_EMPRUNT_ABONNE
BEFORE INSERT ON EMPRUNTER
FOR EACH ROW
BEGIN
    -- Vérifier si l'abonné a déjà un emprunt en cours
    IF EXISTS (
        SELECT 1
        FROM EMPRUNTER
        WHERE ID_ABONNE = NEW.ID_ABONNE
          AND DATE_RETOUR IS NULL
    ) THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Cet abonné a déjà un emprunt en cours.';
    END IF;

    -- Vérifier si le statut de l'abonné est "false"
    IF EXISTS (
        SELECT 1
        FROM ABONNE
        WHERE ID_ABONNE = NEW.ID_ABONNE
          AND STATUS_ABONNE = FALSE
    ) THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Cet abonné a un statut inactif et ne peut pas emprunter.';
    END IF;
END //

DELIMITER ;
