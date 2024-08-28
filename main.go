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
	Output  string   `yaml:"output"`
}

type Column struct {
	Column string `yaml:"column"`
	Name   string `yaml:"name"`
	Index  int    `yaml:"index"`
	Suffix string `yaml:"suffix"`
	Prefix string `yaml:"prefix"`
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
	headers := csvData[0]
	for i, row := range csvData {
		if i == 0 {
			continue
		}
		shipment := Shipment{}
		for _, col := range config.Columns {
			var value string
			if col.Index > 0 && col.Index < len(row) {
				value = row[col.Index-1]
			} else if col.Name != "" {
				for idx, header := range headers {
					if header == col.Name {
						value = row[idx]
						break
					}
				}
			}
			if col.Suffix != "" {
				value += col.Suffix
			}
			if col.Prefix != "" {
				value = col.Prefix + value
			}
			shipment[col.Column] = value
		}
		shipmentList = append(shipmentList, shipment)
	}

	if len(config.Columns) > 0 {
		sort.Slice(shipmentList, func(i, j int) bool {
			return shipmentList[i][config.Columns[0].Column] < shipmentList[j][config.Columns[0].Column]
		})
	}

	var outputFilename string
	if config.Output != "" {
		outputFilename = config.Output
	} else {
		outputFilename = "result.csv"
	}
	resultFile, err := os.Create(outputFilename)
	if err != nil {
		panic(err)
	}
	defer resultFile.Close()
	resultFile.WriteString("\xEF\xBB\xBF")

	writer := csv.NewWriter(resultFile)
	writer.Comma = ';'
	defer writer.Flush()

	resultHeaders := []string{}
	for _, col := range config.Columns {
		resultHeaders = append(resultHeaders, col.Column)
	}
	writer.Write(resultHeaders)

	for _, shipment := range shipmentList {
		row := []string{}
		for _, col := range config.Columns {
			row = append(row, shipment[col.Column])
		}
		writer.Write(row)
	}
}
