package lib

import (
	"encoding/csv"
	"os"
)

func CSVReader(file string) ([][]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return [][]string{}, err
	}
	r := csv.NewReader(f)
	c, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return c, nil
}
