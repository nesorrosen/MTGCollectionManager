package db

import (
	"database/sql"

	// SQLite3 Driver
	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

var schema1 = `
CREATE TABLE IF NOT EXISTS cards (
        name VARCHAR(25) NOT NULL,
        ON DELETE CASCADE
        ON UPDATE CASCADE
);`

var schema2 = `
CREATE TABLE IF NOT EXISTS colors (
        color VARCHAR(25) NOT NULL,
        ON DELETE CASCADE
        ON UPDATE CASCADE
);`

var schema3 = `
CREATE TABLE IF NOT EXISTS types (
        type VARCHAR(25) NOT NULL,
        ON DELETE CASCADE
        ON UPDATE CASCADE
);`

var schema4 = `
CREATE TABLE IF NOT EXISTS subtypes (
        subtype VARCHAR(25) NOT NULL,
        ON DELETE CASCADE
        ON UPDATE CASCADE
);`

var schema5 = `
CREATE TABLE IF NOT EXISTS supertypes (
        supertype VARCHAR(25) NOT NULL,
        ON DELETE CASCADE
        ON UPDATE CASCADE
);`

var schema6 = `
CREATE TABLE IF NOT EXISTS card_colors (
				card_id INT,
        color_id NOT NULL,
				PRIMARY KEY(card_id, color_id),
				FOREIGN KEY(card_id) REFERENCES cards(rowid),
				FOREIGN KEY(color_id) REFERENCES colors(rowid)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);`

var schema7 = `
CREATE TABLE IF NOT EXISTS card_types (
				card_id INT,
        type_id NOT NULL,
				PRIMARY KEY(card_id, type_id),
				FOREIGN KEY(card_id) REFERENCES cards(rowid),
				FOREIGN KEY(type_id) REFERENCES types(rowid)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);`

var schema8 = `
CREATE TABLE IF NOT EXISTS card_subtypes (
				card_id INT,
        subtype_id NOT NULL,
				PRIMARY KEY(card_id, subtype_id),
				FOREIGN KEY(card_id) REFERENCES cards(rowid),
				FOREIGN KEY(subtype_id) REFERENCES subtypes(rowid)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);`

var schema9 = `
CREATE TABLE IF NOT EXISTS card_supertypes (
				card_id INT,
        supertype_id NOT NULL,
				PRIMARY KEY(card_id, supertype_id),
				FOREIGN KEY(card_id) REFERENCES cards(rowid),
				FOREIGN KEY(supertype_id) REFERENCES supertype(rowid)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);`

// Cards is the cards table
type Cards struct {
	ID   int64  `json:"rowid"`
	Name string `json:"name"`
}

// Colors is the colors table
type Colors struct {
	ID    int64  `json:"rowid"`
	Color string `json:"color"`
}

// Types is the types table
type Types struct {
	ID   int64  `json:"rowid"`
	Type string `json:"type"`
}

// Subtypes is the subtypes table
type Subtypes struct {
	ID      int64  `json:"rowid"`
	Subtype string `json:"subtype"`
}

// Supertypes is the supertypes table
type Supertypes struct {
	ID        int64  `json:"rowid"`
	Supertype string `json:"supertype"`
}

// CardColors is the card_colors table
type CardColors struct {
	CardID  int64 `json:"card_id"`
	ColorID int64 `json:"color_id"`
}

// CardTypes is the card_types table
type CardTypes struct {
	CardID int64 `json:"card_id"`
	TypeID int64 `json:"type_id"`
}

// CardSubtypes is the card_subtypes table
type CardSubtypes struct {
	CardID    int64 `json:"card_id"`
	SubtypeID int64 `json:"subtype_id"`
}

// CardSupertypes is the card_supertypes table
type CardSupertypes struct {
	CardID      int64 `json:"card_id"`
	SupertypeID int64 `json:"supertype_id"`
}
