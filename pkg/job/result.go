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

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		fmt.Println(record)
	}
	return nil
}
