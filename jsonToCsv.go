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

	for row, pkg := range packages {
		result = append(result, []string{})
		result[row] = append(result[row], pkg.Pkgname)
		result[row] = append(result[row], pkg.Pkgbase)
		result[row] = append(result[row], pkg.Repo)
		result[row] = append(result[row], pkg.Arch)
		result[row] = append(result[row], pkg.Pkgver)
		result[row] = append(result[row], pkg.Pkgrel)
		result[row] = append(result[row], fmt.Sprintf("%v", pkg.Epoch))
		result[row] = append(result[row], pkg.Pkgdesc)
		result[row] = append(result[row], pkg.URL)
		result[row] = append(result[row], pkg.Filename)
		result[row] = append(result[row], fmt.Sprintf("%v", pkg.CompressedSize))
		result[row] = append(result[row], fmt.Sprintf("%v", pkg.InstalledSize))
		result[row] = append(result[row], pkg.BuildDate)
		result[row] = append(result[row], pkg.LastUpdate)
		result[row] = append(result[row], fmt.Sprintf("%v", pkg.FlagDate))
		result[row] = append(result[row], pkg.Packager)
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
