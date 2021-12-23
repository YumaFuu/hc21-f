package job

import (
	"database/sql"
	"errors"
	"fmt"
	"hc21f/pkg/twitter"
)

var jobTypes = []string{
	"memberid",
	"followings",
	"followers",
}

type jobType = string

func NewJobType(s string) (jobType, error) {
	var b = false
	for _, v := range jobTypes {
		b = (v == s)
		if b {
			break
		}
	}

	if !b {
		return "", errors.New(fmt.Sprintf("Invalid JobType %s", s))
	}

	return jobType(s), nil
}

type Job struct {
	twitter twitter.Twitter
	db      *sql.DB
	jobType jobType
}

func New(t twitter.Twitter, db *sql.DB, s string) (Job, error) {
	jt, err := NewJobType(s)
	if err != nil {
		return Job{}, err
	}
	j := Job{t, db, jt}

	return j, err
}

func (j Job) Do() error {
	var err error
	switch j.jobType {
	case "memberid":
		err = j.FillMemberID()
	case "followings":
		err = j.SearchFollowings()
	case "followers":
		err = j.SearchFollowers()
	}

	return err
}
