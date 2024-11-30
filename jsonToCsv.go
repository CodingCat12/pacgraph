package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"time"
)

var jsonDir string = filepath.Join("packages", "json")
var csvDir string = filepath.Join("packages", "csv")

var startTime time.Time = time.Now()

func main() {
	argParser()

	os.Truncate(pkgFile, 0)
	writeCsvRow(Header, pkgFile)

	jsonFiles, err := os.ReadDir(jsonDir)
	if err != nil {
		log.Fatalf("error opening directory: %v\n", jsonDir)
	}

	var csvData []Package
	for i, file := range jsonFiles {
		if file.IsDir() {
			continue
		}

		fullpath := filepath.Join(jsonDir, file.Name())
		data, err := os.ReadFile(fullpath)
		if err != nil {
			log.Printf("error reading file: %v", err)
			continue
		}

		row := getCsvData(data)
		csvData = append(csvData, row)

		if (i % batchSize) == 0 {
			writePackages(csvData, pkgFile)
			csvData = nil
		}
	}

	if len(csvData) > 0 {
		writePackages(csvData, pkgFile)
	}

	if debugMode {
		logSpecs()
	}
}

func getCsvData(data []byte) Package {
	var unmarshaled Package

	err := json.Unmarshal(data, &unmarshaled)
	if err != nil {
		log.Printf("error reading JSON: %v\n", err)
		return Package{}
	}

	return unmarshaled
}

func writePackages(packages []Package, filePath string) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("error opening file: %v\n", err)
		return
	}
	defer file.Close()

	var result [][]string

	for row, pkg := range packages {
		result = append(result, []string{})
		val := reflect.ValueOf(pkg)
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fieldType := field.Kind()
			if fieldType == reflect.Slice {
				continue
			}

			result[row] = append(result[row], fmt.Sprintf("%v", field.Interface()))
		}
	}

	csvWriter := csv.NewWriter(bufio.NewWriter(file))
	csvWriter.UseCRLF = true

	csvWriter.WriteAll(result)
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		log.Printf("error writing CSV data: %v\n", err)
	}
}

func writeCsvRow(row []string, filePath string) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("error opening file: %v\n", err)
		return
	}
	defer file.Close()

	csvWriter := csv.NewWriter(bufio.NewWriter(file))
	csvWriter.UseCRLF = true

	csvWriter.Write(row)
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		log.Printf("error writing CSV data: %v\n", err)
	}
}