package csv

import (
	"encoding/csv"
	"os"
)

func Add(path string, rows [][]string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := csv.NewWriter(f)

	for _, row := range rows {
		writer.Write(row)
	}
	writer.Flush()

	return nil
}
