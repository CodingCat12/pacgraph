package main

import (
	"bufio"
	"encoding/csv"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func convertArrays(jsonFiles []fs.DirEntry) {
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
	for _, file := range jsonFiles {
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

		for _, value := range row.Maintainers {
			maintainers = append(maintainers, []string{row.Pkgname, value})
		}
		for _, value := range row.Groups {
			groups = append(groups, []string{row.Pkgname, value})
		}
		for _, value := range row.Licenses {
			licenses = append(licenses, []string{row.Pkgname, value})
		}
		for _, value := range row.Conflicts {
			conflicts = append(conflicts, []string{row.Pkgname, value})
		}
		for _, value := range row.Provides {
			provides = append(provides, []string{row.Pkgname, value})
		}
		for _, value := range row.Replaces {
			replaces = append(replaces, []string{row.Pkgname, value})
		}
		for _, value := range row.Depends {
			depends = append(depends, []string{row.Pkgname, value})
		}
		for _, value := range row.Optdepends {
			optdepends = append(optdepends, []string{row.Pkgname, value})
		}
		for _, value := range row.Makedepends {
			makedepends = append(makedepends, []string{row.Pkgname, value})
		}
		for _, value := range row.Checkdepends {
			checkdepends = append(checkdepends, []string{row.Pkgname, value})
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
