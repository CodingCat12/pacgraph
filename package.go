package main

type Package struct {
	Pkgname        string   `json:"pkgname"`
	Pkgbase        string   `json:"pkgbase"`
	Repo           Repo     `json:"repo"`
	Arch           Arch     `json:"arch"`
	Pkgver         string   `json:"pkgver"`
	Pkgdesc        string   `json:"pkgdesc"`
	URL            string   `json:"url"`
	Filename       string   `json:"filename"`
	CompressedSize int64    `json:"compressed_size"`
	InstalledSize  int64    `json:"installed_size"`
	BuildDate      string   `json:"build_date"`
	Packager       string   `json:"packager"`
	Groups         []string `json:"groups"`
	Licenses       []string `json:"licenses"`
	Conflicts      []string `json:"conflicts"`
	Provides       []string `json:"provides"`
	Replaces       []string `json:"replaces"`
	Depends        []string `json:"depends"`
	Optdepends     []string `json:"optdepends"`
	Makedepends    []string `json:"makedepends"`
	Checkdepends   []string `json:"checkdepends"`
}

type Repo string
type Arch string

const (
	Core            Repo = "core"
	Extra           Repo = "extra"
	Multilib        Repo = "multilib"
	CoreTesting     Repo = "core-testing"
	ExtraTesting    Repo = "extra-testing"
	MultilibTesting Repo = "multilib-testing"
)

var repos = []Repo{"core", "extra", "multilib"}

const (
	Any    Arch = "any"
	X86_64 Arch = "x86_64"
)

var header = [...]string{
	"pkgname",
	"pkgbase",
	"repo",
	"arch",
	"pkgver",
	"pkgdesc",
	"URL",
	"filename",
	"compressedSize",
	"installedSize",
	"buildDate",
	"packager",
}
