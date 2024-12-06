package main

import (
	"bufio"
	"encoding/csv"
	"os"
)

func convertValues(packages []Package) {
	writeHeader(pkgHeader[:], adjustedConfig.Paths.PackageFile)

	var result [][]string
	for i, pkg := range packages {

		result = append(result, []string{
			pkg.Pkgname,
			pkg.Pkgbase,
			toString(pkg.Repo),
			toString(pkg.Arch),
			pkg.Pkgver,
			pkg.Pkgdesc,
			pkg.URL,
			pkg.Filename,
			toString(pkg.CompressedSize),
			toString(pkg.InstalledSize),
			pkg.BuildDate,
			pkg.Packager})

		if ((i + 1) % adjustedConfig.BatchSize) == 0 {
			err := writeToCsv(result, adjustedConfig.Paths.PackageFile)
			if err != nil {
				logger.Fatalf("error writing packages to csv: %v", err)
			}

			result = nil
		}
	}

	if len(result) > 0 {
		writeToCsv(result, adjustedConfig.Paths.PackageFile)
	}
}

func writeHeader(header []string, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(bufio.NewWriter(file))
	csvWriter.UseCRLF = true

	csvWriter.Write(header)
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		return err
	}

	return nil
}
