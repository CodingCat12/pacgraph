package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func convertValues(packages []Package) {
	var csvData []Package
	for i, pkg := range packages {
		csvData = append(csvData, pkg)

		if (i % batchSize) == 0 {
			err := writePackages(csvData, pkgFile)
			if err != nil {
				logger.Fatalf("error writing packages to csv: %v", err)
			}

			csvData = nil
		}
	}

	if len(csvData) > 0 {
		writePackages(csvData, pkgFile)
	}
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
			pkg.Pkgdesc,
			pkg.URL,
			pkg.Filename,
			fmt.Sprintf("%v", pkg.CompressedSize),
			fmt.Sprintf("%v", pkg.InstalledSize),
			pkg.BuildDate,
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
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
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
