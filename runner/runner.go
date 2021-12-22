package runner

import (
	"database/sql"
	"hc21f/pkg/database"
	"hc21f/pkg/twitter"
	"log"
)

type runner struct {
	twitter twitter.Twitter
	db      *sql.DB
	job     job.Type
}

func valid() bool {

}

func New(t twitter.Twitter, db *sql.DB, job string) error {
	if err := twitter.Init(); err != nil {
		log.Fatal(err)
	}

	if err := database.Init(); err != nil {
		log.Fatal(err)
	}

	Run(t, db)

	return nil
}
