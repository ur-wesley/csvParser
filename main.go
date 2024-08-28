package main

import (
	"encoding/csv"
	"os"
	"sort"

	"github.com/sqweek/dialog"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Columns []Column `yaml:"columns"`
}

type Column struct {
	Index  int    `yaml:"index"`
	Label  string `yaml:"label"`
	Suffix string `yaml:"suffix"`
}

type Shipment map[string]string

func main() {
	configFile, err := os.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		panic(err)
	}

	filename, err := dialog.File().Filter("csv").Load()
	if err != nil {
		panic(err)
	}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	csvData, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	shipmentList := []Shipment{}
	for i, row := range csvData {
		if i == 0 {
			continue
		}
		shipment := Shipment{}
		for _, col := range config.Columns {
			if len(row) > col.Index {
				if col.Suffix != "" {
					shipment[col.Label] = row[col.Index] + col.Suffix
				} else {
					shipment[col.Label] = row[col.Index]
				}
			}
		}
		shipmentList = append(shipmentList, shipment)
	}

	if len(config.Columns) > 0 {
		sort.Slice(shipmentList, func(i, j int) bool {
			return shipmentList[i][config.Columns[0].Label] < shipmentList[j][config.Columns[0].Label]
		})
	}

	resultFile, err := os.Create("result.csv")
	if err != nil {
		panic(err)
	}
	defer resultFile.Close()

	writer := csv.NewWriter(resultFile)
	writer.Comma = ';'
	defer writer.Flush()

	headers := []string{}
	for _, col := range config.Columns {
		headers = append(headers, col.Label)
	}
	writer.Write(headers)

	for _, shipment := range shipmentList {
		row := []string{}
		for _, col := range config.Columns {
			row = append(row, shipment[col.Label])
		}
		writer.Write(row)
	}
}
