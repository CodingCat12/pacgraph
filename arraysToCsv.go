package main

import (
	"bufio"
	"encoding/csv"
	"os"
	"path/filepath"
)

func convertArrays(packages []Package) {
	var maintainers [][]string
	var groups [][]string
	var licenses [][]string
	var conflicts [][]string
	var provides [][]string
	var replaces [][]string
	var depends [][]string
	var optdepends [][]string
	var makedepends [][]string
	var checkdepends [][]string
	for _, pkg := range packages {
		for _, value := range pkg.Groups {
			groups = append(groups, []string{pkg.Pkgname, value})
		}
		for _, value := range pkg.Licenses {
			licenses = append(licenses, []string{pkg.Pkgname, value})
		}
		for _, value := range pkg.Conflicts {
			conflicts = append(conflicts, []string{pkg.Pkgname, value})
		}
		for _, value := range pkg.Provides {
			provides = append(provides, []string{pkg.Pkgname, value})
		}
		for _, value := range pkg.Replaces {
			replaces = append(replaces, []string{pkg.Pkgname, value})
		}
		for _, value := range pkg.Depends {
			depends = append(depends, []string{pkg.Pkgname, value})
		}
		for _, value := range pkg.Optdepends {
			optdepends = append(optdepends, []string{pkg.Pkgname, value})
		}
		for _, value := range pkg.Makedepends {
			makedepends = append(makedepends, []string{pkg.Pkgname, value})
		}
		for _, value := range pkg.Checkdepends {
			checkdepends = append(checkdepends, []string{pkg.Pkgname, value})
		}

	}

	writeToCsv(maintainers, filepath.Join(csvDir, "maintainers.csv"))
	writeToCsv(groups, filepath.Join(csvDir, "groups.csv"))
	writeToCsv(licenses, filepath.Join(csvDir, "licenses.csv"))
	writeToCsv(conflicts, filepath.Join(csvDir, "conflicts.csv"))
	writeToCsv(provides, filepath.Join(csvDir, "provides.csv"))
	writeToCsv(replaces, filepath.Join(csvDir, "replaces.csv"))
	writeToCsv(depends, filepath.Join(csvDir, "depends.csv"))
	writeToCsv(optdepends, filepath.Join(csvDir, "optdepends.csv"))
	writeToCsv(makedepends, filepath.Join(csvDir, "makedepends.csv"))
	writeToCsv(checkdepends, filepath.Join(csvDir, "checkdepends.csv"))
}

func writeToCsv(packages [][]string, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
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
