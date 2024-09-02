package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/sqweek/dialog"
)

func main() {
	config, err := GetConfig()
	if err != nil {
		ShowErrorDialog(fmt.Sprintf("Failed to get config: %v", err))
		log.Fatalf("Failed to get config: %v", err)
	}

	csvFilename, err := dialog.File().Filter("CSV", "csv").Title("CSV Datei auswählen").Load()
	if err != nil {
		ShowErrorDialog(fmt.Sprintf("Failed to load CSV file: %v", err))
		log.Fatalf("Failed to load CSV file: %v", err)
	}

	csvData, err := LoadCSVData(csvFilename, config)
	if err != nil {
		ShowErrorDialog(fmt.Sprintf("Failed to load CSV data: %v", err))
		log.Fatalf("Failed to load CSV data: %v", err)
	}

	newDataList, err := ProcessCSVData(csvData, config)
	if err != nil {
		ShowErrorDialog(fmt.Sprintf("Failed to process CSV data: %v", err))
		log.Fatalf("Failed to process CSV data: %v", err)
	}

	err = writeResultFile(newDataList, config)
	if err != nil {
		ShowErrorDialog(fmt.Sprintf("Failed to write result file: %v", err))
		log.Fatalf("Failed to write result file: %v", err)
	}
}

func writeResultFile(newDataList []NewData, config Config) error {
	var outputFilename string
	if config.Output != "" {
		outputFilename = config.Output
	} else {
		outputFilename = "result.csv"
	}
	resultFile, err := os.Create(outputFilename)
	if err != nil {
		return err
	}
	defer resultFile.Close()
	resultFile.WriteString("\xEF\xBB\xBF")

	writer := csv.NewWriter(resultFile)
	if config.Delimiter != "" {
		writer.Comma = rune(config.Delimiter[0])
	} else {
		writer.Comma = ';'
	}
	defer writer.Flush()

	resultHeaders := []string{}
	for _, col := range config.Columns {
		resultHeaders = append(resultHeaders, col.Column)
	}
	err = writer.Write(resultHeaders)
	if err != nil {
		return err
	}

	for _, newData := range newDataList {
		row := []string{}
		for _, col := range config.Columns {
			row = append(row, newData[col.Column])
		}
		err = writer.Write(row)
		if err != nil {
			return err
		}
	}

	return nil
}
