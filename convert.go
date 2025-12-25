package csvtojson

import (
	"encoding/csv"
	"encoding/json"
	"io"
)

func Convert(fr io.Reader) ([]byte, error) {
	r := csv.NewReader(fr)
	header, err := r.Read()
	if err != nil {
		return nil, err
	}
	var result []map[string]string
	record, err := r.ReadAll()

	if err != nil {
		return nil, err
	}
	for _, row := range record {
		item := make(map[string]string)
		for col, value := range row {
			item[header[col]] = value
		}
		result = append(result, item)
	}
	b, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		return nil, err
	}
	return b, nil
}
