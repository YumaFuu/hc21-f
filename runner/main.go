package runner

import (
	"database/sql"
	"hc21f/job"
	"hc21f/pkg/twitter"
)

type runner struct {
	twitter twitter.Twitter
	db      *sql.DB
	job     job.Job
}

func New(t twitter.Twitter, db *sql.DB, job string) error {

	return job.New(t, db, job)
}
