package data

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/Jguer/go-alpm/v2"
)

func GetData() ([]Package, error) {
	handle, err := alpm.Initialize("/", "/var/lib/pacman")
	if err != nil {
		return []Package{}, err
	}
	defer handle.Release()

	var result []Package
	for _, repo := range repos {
		db, err := handle.RegisterSyncDB(repo, 0)
		if err != nil {
			return []Package{}, err
		}

		for _, pkg := range db.PkgCache().Slice() {
			packager, err := parsePerson(pkg.Packager())
			if err != nil {
				return []Package{}, err
			}

			result = append(result, (Package{
				Pkgname:        pkg.Name(),
				Pkgbase:        pkg.Base(),
				Repo:           Repo(pkg.DB().Name()),
				Arch:           Arch(pkg.Architecture()),
				Pkgver:         pkg.Version(),
				Pkgdesc:        pkg.Description(),
				URL:            pkg.URL(),
				Filename:       pkg.FileName(),
				CompressedSize: pkg.Size(),
				InstalledSize:  pkg.ISize(),
				BuildDate:      pkg.BuildDate().UTC().Format(time.RFC3339),
				Packager:       packager,
				Groups:         pkg.Groups().Slice(),
				Licenses:       pkg.Licenses().Slice(),
				Conflicts:      dependsToStrings(pkg.Conflicts().Slice()),
				Provides:       dependsToStrings(pkg.Provides().Slice()),
				Replaces:       dependsToStrings(pkg.Replaces().Slice()),
				Depends:        dependsToStrings(pkg.Depends().Slice()),
				Optdepends:     dependsToStrings(pkg.OptionalDepends().Slice()),
				Makedepends:    dependsToStrings(pkg.MakeDepends().Slice()),
				Checkdepends:   dependsToStrings(pkg.CheckDepends().Slice()),
			}))
		}
	}

	return result, nil
}

func dependsToStrings(dependencies []alpm.Depend) []string {
	var names []string
	for _, dep := range dependencies {
		names = append(names, dep.Name)
	}
	return names
}

func parsePerson(input string) (Person, error) {
	re := regexp.MustCompile(`^\s*([^\<]+)\s*<(.+?)>\s*$`)
	matches := re.FindStringSubmatch(input)

	if len(matches) != 3 {
		return Person{}, fmt.Errorf("invalid format for person: %s", input)
	}

	name := strings.TrimSpace(matches[1])

	return Person{
		Name:  name,
		Email: matches[2],
	}, nil
}
