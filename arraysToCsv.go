package main

import (
	"bufio"
	"encoding/csv"
	"os"
	"path/filepath"
)

func convertArrays(packages []Package) {
	type attributeData struct {
		field   func(pkg Package) []string
		file    string
		name    string
		records [][]string
	}

	attributes := []attributeData{
		{func(pkg Package) []string { return pkg.Groups }, "groups.csv", "groups", nil},
		{func(pkg Package) []string { return pkg.Licenses }, "licenses.csv", "licenses", nil},
		{func(pkg Package) []string { return pkg.Conflicts }, "conflicts.csv", "conflicts", nil},
		{func(pkg Package) []string { return pkg.Provides }, "provides.csv", "provides", nil},
		{func(pkg Package) []string { return pkg.Replaces }, "replaces.csv", "replaces", nil},
		{func(pkg Package) []string { return pkg.Depends }, "depends.csv", "depends", nil},
		{func(pkg Package) []string { return pkg.Optdepends }, "optdepends.csv", "optdepends", nil},
		{func(pkg Package) []string { return pkg.Makedepends }, "makedepends.csv", "makedepends", nil},
		{func(pkg Package) []string { return pkg.Checkdepends }, "checkdepends.csv", "checkdepends", nil},
	}

	for _, attr := range attributes {
		file := filepath.Join(csvDir, attr.file)
		header := []string{"pkg", attr.name}
		writeHeader(header, file)
	}

	for i, pkg := range packages {
		for j, attr := range attributes {
			for _, value := range attributes[j].field(pkg) {
				attributes[j].records = append(attributes[j].records, []string{pkg.Pkgname, value})
			}

			if (i % batchSize) == 0 {
				file := filepath.Join(csvDir, attr.file)
				writeToCsv(attr.records, file)
				attributes[j].records = nil
			}
		}
	}

	for _, attr := range attributes {
		file := filepath.Join(csvDir, attr.file)
		writeToCsv(attr.records, file)
	}
}

func writeToCsv(data [][]string, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(bufio.NewWriter(file))
	csvWriter.UseCRLF = true

	csvWriter.WriteAll(data)
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		return err
	}

	return nil
}
