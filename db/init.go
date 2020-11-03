package db

import "database/sql"

func init() {
	var err error
	conn, err = sql.Open("sqlite3", "./mtg_collection.db")
	if err != nil {
		panic(err)
	}

	_, err = conn.Exec(schema1)
	if err != nil {
		panic(err)
	}

	_, err = conn.Exec(schema2)
	if err != nil {
		panic(err)
	}

	_, err = conn.Exec(schema3)
	if err != nil {
		panic(err)
	}

	_, err = conn.Exec(schema4)
	if err != nil {
		panic(err)
	}
	_, err = conn.Exec(schema5)
	if err != nil {
		panic(err)
	}

	_, err = conn.Exec(schema6)
	if err != nil {
		panic(err)
	}

	_, err = conn.Exec(schema7)
	if err != nil {
		panic(err)
	}

	_, err = conn.Exec(schema8)
	if err != nil {
		panic(err)
	}

	_, err = conn.Exec(schema9)
	if err != nil {
		panic(err)
	}

}
