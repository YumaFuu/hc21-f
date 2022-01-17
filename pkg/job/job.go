package job

import (
	"errors"
	"fmt"
	"hc21f/pkg/twitter"
)

var jobTypes = []string{
	"members",
	"friends",
	"followers",
	"result",
}

type jobType string

func NewJobType(s string) (jobType, error) {
	var b = false
	for _, v := range jobTypes {
		b = (v == s)
		if b {
			break
		}
	}

	if !b {
		return "", errors.New(fmt.Sprintf("Invalid JobType: %s", s))
	}

	return jobType(s), nil
}

type Job struct {
	twitter twitter.Twitter
	jobType jobType
}

func New(t twitter.Twitter, s string) (Job, error) {
	jt, err := NewJobType(s)
	if err != nil {
		return Job{}, err
	}
	j := Job{t, jt}

	return j, err
}

func (j Job) Do() error {
	var err error
	switch j.jobType {
	case "members":
		err = j.FillMemberID()
	case "friends":
		err = j.SearchFriends()
	case "followers":
		err = j.SearchFollowers()
	case "result":
		err = j.GetResult()
	}

	return err
}
