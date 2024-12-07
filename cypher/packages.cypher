CREATE INDEX FOR (p:Package) ON (p.name);
CREATE INDEX FOR (p:Package) ON (p.base);
CREATE INDEX FOR (p:Package) ON (p.version);
CREATE INDEX FOR (p:Package) ON (p.description);
CREATE INDEX FOR (p:Package) ON (p.URL);
CREATE INDEX FOR (p:Package) ON (p.filename);
CREATE INDEX FOR (p:Package) ON (p.compressedSize);
CREATE INDEX FOR (p:Package) ON (p.installedSize);
CREATE INDEX FOR (p:Package) ON (p.buildDate);
CREATE INDEX FOR (p:Package) ON (p.packager);
CREATE INDEX FOR (r:Repo) ON (r.name);
CREATE INDEX FOR (a:Arch) ON (a.name);

LOAD CSV WITH HEADERS FROM 'https://raw.githubusercontent.com/CodingCat12/pacgraph/refs/heads/main/packages/packages.csv' AS pkgrow
MERGE (p:Package {
  name: pkgrow.pkgname,
  base: pkgrow.pkgbase,
  version: pkgrow.pkgver,
  description: pkgrow.pkgdesc,
  URL: pkgrow.URL,
  filename: pkgrow.filename,
  compressedSize: toInteger(pkgrow.compressedSize),
  installedSize: toInteger(pkgrow.installedSize),
  buildDate: pkgrow.buildDate,
  packager: pkgrow.packager
})

MERGE (r:Repo {name: pkgrow.repo})
MERGE (a:Arch {name: pkgrow.arch})

MERGE (p)-[:IS_IN_REPO]->(r)
MERGE (p)-[:SUPPORTS_ARCHITECTURE]->(a)

RETURN *;