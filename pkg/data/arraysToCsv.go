package data

import (
	"bufio"
	"encoding/csv"
	"os"

	"github.com/CodingCat12/pacgraph/pkg/config"
)

func ConvertArrays(packages []Package) error {
	type attributeData struct {
		field   func(pkg Package) []string
		file    string
		name    string
		records [][]string
	}

	attributes := []attributeData{
		{func(pkg Package) []string { return pkg.Groups }, config.AdjustedConfig.Paths.GroupsFile, "groups", nil},
		{func(pkg Package) []string { return pkg.Licenses }, config.AdjustedConfig.Paths.LicensesFile, "licenses", nil},
		{func(pkg Package) []string { return pkg.Conflicts }, config.AdjustedConfig.Paths.ConflictsFile, "conflicts", nil},
		{func(pkg Package) []string { return pkg.Provides }, config.AdjustedConfig.Paths.ProvidesFile, "provides", nil},
		{func(pkg Package) []string { return pkg.Replaces }, config.AdjustedConfig.Paths.ReplacesFile, "replaces", nil},
		{func(pkg Package) []string { return pkg.Depends }, config.AdjustedConfig.Paths.DependsFile, "depends", nil},
		{func(pkg Package) []string { return pkg.Optdepends }, config.AdjustedConfig.Paths.OptDependsFile, "optdepends", nil},
		{func(pkg Package) []string { return pkg.Makedepends }, config.AdjustedConfig.Paths.MakeDependsFile, "makedepends", nil},
		{func(pkg Package) []string { return pkg.Checkdepends }, config.AdjustedConfig.Paths.CheckDependsFile, "checkdepends", nil},
	}

	for i := range attributes {
		file := attributes[i].file
		header := []string{"pkg", attributes[i].name}
		err := writeHeader(header, file)
		if err != nil {
			return err
		}

		for j, pkg := range packages {
			for _, value := range attributes[i].field(pkg) {
				attributes[i].records = append(attributes[i].records, []string{pkg.Pkgname, value})
			}

			if ((j + 1) % config.AdjustedConfig.BatchSize) == 0 {
				err := writeToCsv(attributes[i].records, file)
				if err != nil {
					return err
				}
				attributes[i].records = nil
			}
		}

		err = writeToCsv(attributes[i].records, file)
		if err != nil {
			return err
		}
	}

	return nil
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
