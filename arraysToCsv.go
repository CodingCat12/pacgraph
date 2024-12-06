package main

import (
	"bufio"
	"encoding/csv"
	"os"
)

func convertArrays(packages []Package) {
	type attributeData struct {
		field   func(pkg Package) []string
		file    string
		name    string
		records [][]string
	}

	attributes := []attributeData{
		{func(pkg Package) []string { return pkg.Groups }, adjustedConfig.Paths.GroupsFile, "groups", nil},
		{func(pkg Package) []string { return pkg.Licenses }, adjustedConfig.Paths.LicensesFile, "licenses", nil},
		{func(pkg Package) []string { return pkg.Conflicts }, adjustedConfig.Paths.ConflictsFile, "conflicts", nil},
		{func(pkg Package) []string { return pkg.Provides }, adjustedConfig.Paths.ProvidesFile, "provides", nil},
		{func(pkg Package) []string { return pkg.Replaces }, adjustedConfig.Paths.ReplacesFile, "replaces", nil},
		{func(pkg Package) []string { return pkg.Depends }, adjustedConfig.Paths.DependsFile, "depends", nil},
		{func(pkg Package) []string { return pkg.Optdepends }, adjustedConfig.Paths.OptDependsFile, "optdepends", nil},
		{func(pkg Package) []string { return pkg.Makedepends }, adjustedConfig.Paths.MakeDependsFile, "makedepends", nil},
		{func(pkg Package) []string { return pkg.Checkdepends }, adjustedConfig.Paths.CheckDependsFile, "checkdepends", nil},
	}

	for i := range attributes {
		file := attributes[i].file
		header := []string{"pkg", attributes[i].name}
		writeHeader(header, file)

		for j, pkg := range packages {
			for _, value := range attributes[i].field(pkg) {
				attributes[i].records = append(attributes[i].records, []string{pkg.Pkgname, value})
			}

			if ((j + 1) % adjustedConfig.BatchSize) == 0 {
				writeToCsv(attributes[i].records, file)
				attributes[i].records = nil
			}
		}

		writeToCsv(attributes[i].records, file)
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
