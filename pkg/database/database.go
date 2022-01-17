package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const DBFile = "./hc21f.db"

var db *sql.DB

func Get() *sql.DB {
	return db
}

func Init() (err error) {
	db, err = sql.Open("sqlite3", DBFile)
	if err = createTables(); err != nil {
		return err
	}
	return err
}

func DB() *sql.DB {
	return db
}

func createTables() error {
	_, err := db.Exec(
		`
    CREATE TABLE IF NOT EXISTS
      friends
    (
      id INTEGER PRIMARY KEY,
      uid INTENTGER NOT NULL,
      user_id INTEGER NOT NULL
    );
    `,
	)
	if err != nil {
		return err
	}

	return nil
}
