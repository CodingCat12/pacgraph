package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
)

func convertArrays() {
	jsonFiles, err := os.ReadDir(jsonDir)
	if err != nil {
		log.Fatalf("error opening directory: %v\n", jsonDir)
	}

	var result [][]string
	for _, file := range jsonFiles {
		if file.IsDir() {
			continue
		}

		fullpath := filepath.Join(jsonDir, file.Name())
		data, err := os.ReadFile(fullpath)
		if err != nil {
			log.Printf("error reading file: %v", err)
			continue
		}

		row, err := jsonToPackage(data)
		if err != nil {
			log.Printf("error reading JSON: %v\n", err)
			continue
		}

		for _, dependency := range row.Depends {
			result = append(result, []string{row.Pkgname, dependency})
		}
	}

	writeToCsv(result, filepath.Join(csvDir, "depends.csv"))
}

func writeToCsv(packages [][]string, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(bufio.NewWriter(file))
	csvWriter.UseCRLF = true

	csvWriter.WriteAll(packages)
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		return err
	}

	return nil
}
