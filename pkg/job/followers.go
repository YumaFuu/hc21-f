package job

import (
	"fmt"
	"hc21f/pkg/csv"
)

var (
	FollowerCsvHeader   = []string{"ID"}
	FollowerCsvFilePath = "./data/follower.csv"
)

func (job *Job) SearchFollowers() error {
	followers, err := job.twitter.GetFollowers("778927219422797824")
	if err != nil {
		return err
	}

	var rows [][]string
	rows = append(rows, FollowerCsvHeader)

	for _, f := range followers {
		rows = append(rows, []string{fmt.Sprint(f)})
	}

	err = csv.Add(FollowerCsvFilePath, rows)
	if err != nil {
		return err
	}
	return nil
}
