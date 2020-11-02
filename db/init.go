package db

import "database/sql"

func init() {
	db, err := sql.Open("sqlite3", "./mtg_collection.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(schema1)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(schema2)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(schema3)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(schema4)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(schema5)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(schema6)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(schema7)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(schema8)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(schema9)
	if err != nil {
		panic(err)
	}

}
