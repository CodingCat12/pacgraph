package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/goccy/go-json"
)

var jsonDir string = filepath.Join("packages", "json")
var csvDir string = filepath.Join("packages", "csv")

var startTime time.Time = time.Now()

func main() {
	argParser()

	writeHeaders(pkgFile)

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

		row, err := jsonToPackage(data)
		if err != nil {
			log.Printf("error reading JSON: %v\n", err)
			continue
		}

		csvData = append(csvData, *row)

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

func jsonToPackage(data []byte) (*Package, error) {
	var unmarshaled Package

	err := json.Unmarshal(data, &unmarshaled)
	if err != nil {
		return nil, err
	}

	return &unmarshaled, nil
}

func writePackages(packages []Package, filePath string) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("error opening file: %v\n", err)
		return
	}
	defer file.Close()

	var result [][]string

	for _, pkg := range packages {
		result = append(result, []string{
			pkg.Pkgname,
			pkg.Pkgbase,
			pkg.Repo,
			pkg.Arch,
			pkg.Pkgver,
			pkg.Pkgrel,
			fmt.Sprintf("%v", pkg.Epoch),
			pkg.Pkgdesc,
			pkg.URL,
			pkg.Filename,
			fmt.Sprintf("%v", pkg.CompressedSize),
			fmt.Sprintf("%v", pkg.InstalledSize),
			pkg.BuildDate,
			pkg.LastUpdate,
			fmt.Sprintf("%v", pkg.FlagDate),
			pkg.Packager})
	}

	csvWriter := csv.NewWriter(bufio.NewWriter(file))
	csvWriter.UseCRLF = true

	csvWriter.WriteAll(result)
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		log.Printf("error writing CSV data: %v\n", err)
	}
}

func writeHeaders(filePath string) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Printf("error opening file: %v\n", err)
		return
	}
	defer file.Close()

	csvWriter := csv.NewWriter(bufio.NewWriter(file))
	csvWriter.UseCRLF = true

	csvWriter.Write(Header[:])
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		log.Printf("error writing CSV data: %v\n", err)
	}
}
