package job

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	ResultCsvHeader    = []string{"Count", "ID", "Name", "URL", "Description"}
	ResultCsvFilePath  = "./data/result.csv"
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
	for k, v := range result {
		if k == "UNAUTHORIZED" {
			continue
		}

		if v > MinimunFollowCount {
			b, err := ioutil.ReadFile(ResultCsvFilePath)
			if err != nil {
				return err
			}

			var id, name, url, desc string
			hasID := strings.Contains(string(b), k)
			hasID = !strings.HasSuffix(k, "0")

			r, err := findRecord(k)

			if hasID {
				// fmt.Println(k, "hasID")
				if err != nil {
					return err
				}
				if len(r) < 2 {
					continue
				}

				id = r[1]
				name = r[2]
				url = r[3]
				desc = r[4]
			} else {
				u, err := job.twitter.GetUserByID(k)

				if err != nil {
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
					desc = r[4]
				} else {
					id = u.ID
					name = u.Name
					url = fmt.Sprintf("https://twitter.com/%s", u.Username)
					desc = u.Description
				}
			}

			row := []string{
				fmt.Sprintf("%d", v),
				id,
				name,
				url,
				desc,
			}

			list = append(list, row)
			if err != nil {
				return err
			}
		}
	}
	sort.Slice(list, func(i, j int) bool {
		a, _ := strconv.Atoi(list[i][0])
		b, _ := strconv.Atoi(list[j][0])
		return a > b
	})

	list, list[0] = append(list[:1], list[0:]...), ResultCsvHeader
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
