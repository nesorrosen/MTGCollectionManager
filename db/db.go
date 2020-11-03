package db

import (
	"database/sql"

	// SQLite3 Driver
	_ "github.com/mattn/go-sqlite3"
)

var (
	conn *sql.DB
)

var schema1 = `
CREATE TABLE IF NOT EXISTS cards (
        name VARCHAR(25) NOT NULL,
				amount INT(8) DEFAULT 1,
				comment VARCHAR(500)
);`

var schema2 = `
CREATE TABLE IF NOT EXISTS colors (
        color VARCHAR(25) NOT NULL
);`

var schema3 = `
CREATE TABLE IF NOT EXISTS types (
        type VARCHAR(25) NOT NULL
);`

var schema4 = `
CREATE TABLE IF NOT EXISTS subtypes (
        subtype VARCHAR(25) NOT NULL
);`

var schema5 = `
CREATE TABLE IF NOT EXISTS supertypes (
        supertype VARCHAR(25) NOT NULL
);`

var schema6 = `
CREATE TABLE IF NOT EXISTS card_colors (
				card_id INT,
        color_id NOT NULL,
				PRIMARY KEY(card_id, color_id),
				FOREIGN KEY(card_id) REFERENCES cards(rowid)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
				FOREIGN KEY(color_id) REFERENCES colors(rowid)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);`

var schema7 = `
CREATE TABLE IF NOT EXISTS card_types (
				card_id INT,
        type_id NOT NULL,
				PRIMARY KEY(card_id, type_id),
				FOREIGN KEY(card_id) REFERENCES cards(rowid)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
				FOREIGN KEY(type_id) REFERENCES types(rowid)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);`

var schema8 = `
CREATE TABLE IF NOT EXISTS card_subtypes (
				card_id INT,
        subtype_id NOT NULL,
				PRIMARY KEY(card_id, subtype_id),
				FOREIGN KEY(card_id) REFERENCES cards(rowid)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
				FOREIGN KEY(subtype_id) REFERENCES subtypes(rowid)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);`

var schema9 = `
CREATE TABLE IF NOT EXISTS card_supertypes (
				card_id INT,
        supertype_id NOT NULL,
				PRIMARY KEY(card_id, supertype_id),
				FOREIGN KEY(card_id) REFERENCES cards(rowid)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
				FOREIGN KEY(supertype_id) REFERENCES supertype(rowid)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);`

// Card is the card table
type Card struct {
	ID      int64  `json:"rowid"`
	Name    string `json:"name"`
	Amount  int    `json:"amount"`
	Comment string `json:"comment"`
}

// Color is the color table
type Color struct {
	ID    int64  `json:"rowid"`
	Color string `json:"color"`
}

// Type is the type table
type Type struct {
	ID   int64  `json:"rowid"`
	Type string `json:"type"`
}

// Subtype is the subtype table
type Subtype struct {
	ID      int64  `json:"rowid"`
	Subtype string `json:"subtype"`
}

// Supertype is the supertype table
type Supertype struct {
	ID        int64  `json:"rowid"`
	Supertype string `json:"supertype"`
}

// CardColor is the card_color table
type CardColor struct {
	CardID  int64 `json:"card_id"`
	ColorID int64 `json:"color_id"`
}

// CardType is the card_type table
type CardType struct {
	CardID int64 `json:"card_id"`
	TypeID int64 `json:"type_id"`
}

// CardSubtype is the card_subtype table
type CardSubtype struct {
	CardID    int64 `json:"card_id"`
	SubtypeID int64 `json:"subtype_id"`
}

// CardSupertype is the card_supertype table
type CardSupertype struct {
	CardID      int64 `json:"card_id"`
	SupertypeID int64 `json:"supertype_id"`
}
