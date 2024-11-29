package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"
)

const batchSize = 2000

const jsonDir string = "packages/json"
const csvDir string = "packages/csv"

var pkgFile string = filepath.Join(csvDir, "packages2.csv")

func main() {
	startTime := time.Now()
	os.Truncate(pkgFile, 0)

	jsonFiles, err := os.ReadDir(jsonDir)
	if err != nil {
		log.Fatalf("error opening directory: %v\n", jsonDir)
	}

	var csvData [][]string
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

		if i == 0 {
			csvData = append(csvData, getCsvHeaders(data))
		}

		row := getCsvData(data)
		csvData = append(csvData, row)

		if (i % batchSize) == 0 {
			writeToCsv(csvData, pkgFile)
			csvData = nil
		}
	}

	if len(csvData) > 0 {
		writeToCsv(csvData, pkgFile)
	}

	fmt.Printf("Operation took %v\n", time.Since(startTime))
}

func getCsvData(jsonString []byte) []string {
	data := jsonString
	var unmarshaled map[string]interface{}

	err := json.Unmarshal(data, &unmarshaled)
	if err != nil {
		log.Printf("error reading JSON: %v\n", err)
		return nil
	}

	var values []string
	var keys []string

	for key := range unmarshaled {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		if _, ok := unmarshaled[key].([]interface{}); ok {
			continue
		}
		values = append(values, fmt.Sprintf("%v", unmarshaled[key]))
	}

	return values
}

func getCsvHeaders(jsonString []byte) []string {
	data := jsonString
	var unmarshaled map[string]interface{}

	err := json.Unmarshal(data, &unmarshaled)
	if err != nil {
		log.Printf("error reading JSON: %v\n", err)
		return nil
	}

	var keys []string

	for key, value := range unmarshaled {
		if _, isArray := value.([]interface{}); !isArray {
			keys = append(keys, key)
		}
	}

	sort.Strings(keys)

	return keys
}

func writeToCsv(values [][]string, filePath string) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("error opening file: %v\n", err)
		return
	}
	defer file.Close()

	csvWriter := csv.NewWriter(bufio.NewWriter(file))
	csvWriter.UseCRLF = true

	csvWriter.WriteAll(values)
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		log.Printf("error writing CSV data: %v\n", err)
	}
}
