package job

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

var (
	ResultCsvHeader   = []string{"ID"}
	ResultCsvFilePath = "./data/result.csv"
)

func (job *Job) GetResult() error {
	f, err := os.Open(FriendsCsvFilePath)
	if err != nil {
		return err
	}

	r := csv.NewReader(f)

	result := make(map[string]int)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		id := record[1]
		result[id] += 1
	}

	file, err := os.OpenFile(
		ResultCsvFilePath,
		os.O_APPEND|os.O_WRONLY,
		os.ModeAppend,
	)
	for k, v := range result {
		if v > 10 {
			u, err := job.twitter.GetUserByID(k)
			if err != nil {
				return err
			}

			str := fmt.Sprintf("%d,%s,%s,https://twitter.com/%s\n", v, u.ID, u.Name, u.Username)
			_, err = file.Write([]byte(str))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
