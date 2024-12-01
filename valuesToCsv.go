package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/goccy/go-json"
)

func convertValues(jsonFiles []fs.DirEntry) {
	err := writeHeaders(pkgFile)
	if err != nil {
		logger.Errorf("error writing headers: %v", err)
	}

	var csvData []Package
	for i, file := range jsonFiles {
		if file.IsDir() {
			continue
		}

		fullpath := filepath.Join(jsonDir, file.Name())
		data, err := os.ReadFile(fullpath)
		if err != nil {
			logger.Errorf("error reading file: %v", err)
			continue
		}

		pkg, err := jsonToPackage(data)
		if err != nil {
			logger.Errorf("error reading JSON: %v\n", err)
			continue
		}

		csvData = append(csvData, *pkg)

		if (i % batchSize) == 0 {
			err := writePackages(csvData, pkgFile)
			if err != nil {
				logger.Errorf("error writing packages to csv: %v", err)
			}

			csvData = nil
		}
	}

	if len(csvData) > 0 {
		writePackages(csvData, pkgFile)
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

func writePackages(packages []Package, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	var result [][]string

	for _, pkg := range packages {
		result = append(result, []string{
			pkg.Pkgname,
			pkg.Pkgbase,
			fmt.Sprintf("%v", pkg.Repo),
			fmt.Sprintf("%v", pkg.Arch),
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
		return err
	}

	return nil
}

func writeHeaders(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(bufio.NewWriter(file))
	csvWriter.UseCRLF = true

	csvWriter.Write(header[:])
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		return err
	}

	return nil
}
