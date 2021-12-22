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
      members
    (
      id INTEGER PRIMARY KEY,
      uid VARCHAR(255),
      username VARCHAR(255) NOT NULL
    );

    CREATE TABLE IF NOT EXISTS
      users
    (
      id INTEGER PRIMARY KEY,
      uid VARCHAR(255),
      following_member_id INTEGER NOT NULL,
      username VARCHAR(255) NOT NULL
    );

    CREATE TABLE IF NOT EXISTS
      user_followings
    (
      id INTEGER PRIMARY KEY,
      uid VARCHAR(255),
      user_id INTEGER NOT NULL,
      username VARCHAR(255) NOT NULL,
      followed_count INTEGER NOT NULL DEFAULT 0
    );
    `,
	)
	if err != nil {
		return err
	}

	return nil
}
