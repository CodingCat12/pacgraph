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

	for _, pkg := range packages {
		for i := range attributes {
			for _, value := range attributes[i].field(pkg) {
				attributes[i].records = append(attributes[i].records, []string{pkg.Pkgname, value})
			}
		}
	}

	for _, attr := range attributes {
		file := filepath.Join(csvDir, attr.file)
		header := []string{"pkg", attr.name}
		writeToCsv(attr.records, header, file)
	}
}

func writeToCsv(data [][]string, header []string, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(bufio.NewWriter(file))
	csvWriter.UseCRLF = true

	csvWriter.Write(header)
	csvWriter.WriteAll(data)
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		return err
	}

	return nil
}
