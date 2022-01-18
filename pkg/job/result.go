package job

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var (
	ResultCsvHeader    = []string{"Count", "ID", "Name", "URL"}
	ResultCsvFilePath  = "./data/row_result.csv"
	MinimunFollowCount = 5
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

	_, err = os.OpenFile(
		ResultCsvFilePath,
		os.O_APPEND|os.O_WRONLY,
		os.ModeAppend,
	)
	if err != nil {
		return err
	}

	var list [][]string
	list = append(list, ResultCsvHeader)
	for k, v := range result {
		if v > MinimunFollowCount {
			b, err := ioutil.ReadFile(ResultCsvFilePath)
			if err != nil {
				return err
			}

			var id, name, url string
			hasID := strings.Contains(string(b), k)
			if hasID {
				fmt.Println(k, "hasID")
				r, err := findRecord(k)
				if err != nil {
					return err
				}
				if len(r) < 2 {
					continue
				}

				id = r[1]
				name = r[2]
				url = r[3]
			} else {
				u, err := job.twitter.GetUserByID(k)
				if err != nil {
					return err
				}
				id = u.ID
				name = u.Name
				url = fmt.Sprintf("https://twitter.com/%s", u.Username)
			}

			row := []string{
				fmt.Sprintf("%d", v),
				id,
				name,
				url,
			}

			list = append(list, row)
			if err != nil {
				return err
			}
		}
	}

	f, err = os.Create(ResultCsvFilePath)
	if err != nil {
		return err
	}

	w := csv.NewWriter(f)
	w.WriteAll(list)

	w.Flush()
	return nil
}

func findRecord(id string) ([]string, error) {
	f, err := os.Open(ResultCsvFilePath)
	if err != nil {
		return []string{}, err
	}

	r := csv.NewReader(f)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return []string{}, err
		}

		if record[1] == id {
			return record, nil
		}
	}

	return []string{}, nil
}
