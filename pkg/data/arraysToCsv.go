package data

import (
	"bufio"
	"encoding/csv"
	"os"

	"github.com/CodingCat12/pacgraph/pkg/config"
)

type Attribute struct {
	Field func(pkg Package) []string
	File  string
	Name  string
}

func ConvertArrays(packages []Package) error {

	attributes := []Attribute{
		{func(pkg Package) []string { return pkg.Groups }, config.AdjustedConfig.Paths.GroupsFile, "groups"},
		{func(pkg Package) []string { return pkg.Licenses }, config.AdjustedConfig.Paths.LicensesFile, "licenses"},
		{func(pkg Package) []string { return pkg.Conflicts }, config.AdjustedConfig.Paths.ConflictsFile, "conflicts"},
		{func(pkg Package) []string { return pkg.Provides }, config.AdjustedConfig.Paths.ProvidesFile, "provides"},
		{func(pkg Package) []string { return pkg.Replaces }, config.AdjustedConfig.Paths.ReplacesFile, "replaces"},
		{func(pkg Package) []string { return pkg.Depends }, config.AdjustedConfig.Paths.DependsFile, "depends"},
		{func(pkg Package) []string { return pkg.Optdepends }, config.AdjustedConfig.Paths.OptDependsFile, "optdepends"},
		{func(pkg Package) []string { return pkg.Makedepends }, config.AdjustedConfig.Paths.MakeDependsFile, "makedepends"},
		{func(pkg Package) []string { return pkg.Checkdepends }, config.AdjustedConfig.Paths.CheckDependsFile, "checkdepends"},
	}

	for _, attr := range attributes {
		file := attr.File
		header := []string{"pkg", attr.Name}
		var records [][]string

		err := writeHeader(header, file)
		if err != nil {
			return err
		}

		for i, pkg := range packages {
			for _, value := range attr.Field(pkg) {
				records = append(records, []string{pkg.Pkgname, value})
			}

			if ((i + 1) % config.AdjustedConfig.BatchSize) == 0 {
				err := writeToCsv(records, file)
				if err != nil {
					return err
				}
				records = nil
			}
		}

		if len(records) > 0 {
			err = writeToCsv(records, file)
			if err != nil {
				return err
			}
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
