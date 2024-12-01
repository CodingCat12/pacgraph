package main

type Package struct {
	Pkgname        string   `json:"pkgname"`
	Pkgbase        string   `json:"pkgbase"`
	Repo           Repo     `json:"repo"`
	Arch           Arch     `json:"arch"`
	Pkgver         string   `json:"pkgver"`
	Pkgrel         string   `json:"pkgrel"`
	Epoch          int      `json:"epoch"`
	Pkgdesc        string   `json:"pkgdesc"`
	URL            string   `json:"url"`
	Filename       string   `json:"filename"`
	CompressedSize uint64   `json:"compressed_size"`
	InstalledSize  uint64   `json:"installed_size"`
	BuildDate      string   `json:"build_date"`
	LastUpdate     string   `json:"last_update"`
	FlagDate       *string  `json:"flag_date"`
	Maintainers    []string `json:"maintainers"`
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
	Core     Repo = "core"
	Extra    Repo = "extra"
	Multilib Repo = "multilib"
)

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
	"pkgrel",
	"epoch",
	"pkgdesc",
	"URL",
	"filename",
	"compressedSize",
	"installedSize",
	"buildDate",
	"lastUpdate",
	"flagDate",
	"packager",
}
