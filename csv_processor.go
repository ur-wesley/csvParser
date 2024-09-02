package main

import (
	"encoding/csv"
	"errors"
	"os"
	"sort"
	"strings"
)

type NewData map[string]string

type CsvData struct {
	Header []string
	Data   [][]string
}

func LoadCSVData(filename string, config Config) (CsvData, error) {
	csvData := CsvData{}
	file, err := os.Open(filename)
	if err != nil {
		return CsvData{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		return CsvData{}, err
	}

	if config.IgnoreHeader {
		csvData.Header = []string{}
		csvData.Data = data
	} else {
		csvData.Header = data[0]
		csvData.Data = data[1:]
	}

	return csvData, nil
}

func ProcessCSVData(csvData CsvData, config Config) ([]NewData, error) {
	newDataList := []NewData{}
	for i, row := range csvData.Data {
		if config.IgnoreHeader && i == 0 {
			continue
		}
		newData := NewData{}
		for _, col := range config.Columns {
			var value string
			if col.Index > 0 && col.Index < len(row) {
				value = row[col.Index-1]
			} else if !config.IgnoreHeader && col.Name != "" {
				for idx, header := range csvData.Header {
					if header == col.Name {
						value = row[idx]
						break
					}
				}
			} else {
				return nil, errors.New("cannot use config.name with config.ignore_header == true")
			}

			if len(col.Replace) > 0 {
				for old, new := range col.Replace {
					if strings.Contains(value, old) {
						value = strings.ReplaceAll(value, old, new)
					}
				}
			}

			if col.Suffix != "" {
				value += col.Suffix
			}

			if col.Prefix != "" {
				value = col.Prefix + value
			}
			newData[col.Column] = value
		}
		newDataList = append(newDataList, newData)
	}

	if len(config.Columns) > 0 {
		sort.Slice(newDataList, func(i, j int) bool {
			return newDataList[i][config.Columns[0].Column] < newDataList[j][config.Columns[0].Column]
		})
	}

	return newDataList, nil
}
