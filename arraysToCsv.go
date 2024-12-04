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
		records [][]string
	}

	attributes := []attributeData{
		{func(pkg Package) []string { return pkg.Groups }, "groups.csv", nil},
		{func(pkg Package) []string { return pkg.Licenses }, "licenses.csv", nil},
		{func(pkg Package) []string { return pkg.Conflicts }, "conflicts.csv", nil},
		{func(pkg Package) []string { return pkg.Provides }, "provides.csv", nil},
		{func(pkg Package) []string { return pkg.Replaces }, "replaces.csv", nil},
		{func(pkg Package) []string { return pkg.Depends }, "depends.csv", nil},
		{func(pkg Package) []string { return pkg.Optdepends }, "optdepends.csv", nil},
		{func(pkg Package) []string { return pkg.Makedepends }, "makedepends.csv", nil},
		{func(pkg Package) []string { return pkg.Checkdepends }, "checkdepends.csv", nil},
	}

	for _, pkg := range packages {
		for i := range attributes {
			for _, value := range attributes[i].field(pkg) {
				attributes[i].records = append(attributes[i].records, []string{pkg.Pkgname, value})
			}
		}
	}

	for _, attr := range attributes {
		writeToCsv(attr.records, filepath.Join(csvDir, attr.file))
	}
}

func writeToCsv(packages [][]string, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
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
