package data

import (
	"bufio"
	"encoding/csv"
	"os"

	"github.com/CodingCat12/pacgraph/pkg/config"
	"github.com/CodingCat12/pacgraph/pkg/helper"
)

func ConvertValues(packages []Package) error {
	writeHeader(pkgHeader[:], config.AdjustedConfig.Paths.PackageFile)

	var result [][]string
	for i, pkg := range packages {

		result = append(result, []string{
			pkg.Pkgname,
			pkg.Pkgbase,
			helper.ToString(pkg.Repo),
			helper.ToString(pkg.Arch),
			pkg.Pkgver,
			pkg.Pkgdesc,
			pkg.URL,
			pkg.Filename,
			helper.ToString(pkg.CompressedSize),
			helper.ToString(pkg.InstalledSize),
			pkg.BuildDate,
			pkg.Packager})

		if ((i + 1) % config.AdjustedConfig.BatchSize) == 0 {
			err := writeToCsv(result, config.AdjustedConfig.Paths.PackageFile)
			if err != nil {
				return err
			}

			result = nil
		}
	}

	if len(result) > 0 {
		writeToCsv(result, config.AdjustedConfig.Paths.PackageFile)
	}

	return nil
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
