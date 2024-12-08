CREATE INDEX IF NOT EXISTS FOR (p:Package) ON (p.name);
CREATE INDEX IF NOT EXISTS FOR (p:Package) ON (p.base);
CREATE INDEX IF NOT EXISTS FOR (p:Package) ON (p.version);
CREATE INDEX IF NOT EXISTS FOR (p:Package) ON (p.description);
CREATE INDEX IF NOT EXISTS FOR (p:Package) ON (p.URL);
CREATE INDEX IF NOT EXISTS FOR (p:Package) ON (p.filename);
CREATE INDEX IF NOT EXISTS FOR (p:Package) ON (p.compressedSize);
CREATE INDEX IF NOT EXISTS FOR (p:Package) ON (p.installedSize);
CREATE INDEX IF NOT EXISTS FOR (p:Package) ON (p.buildDate);
CREATE INDEX IF NOT EXISTS FOR (p:Package) ON (p.packager);
CREATE INDEX IF NOT EXISTS FOR (r:Repo) ON (r.name);
CREATE INDEX IF NOT EXISTS FOR (a:Arch) ON (a.name);

LOAD CSV WITH HEADERS FROM 'https://raw.githubusercontent.com/CodingCat12/pacgraph/refs/heads/main/packages/packages.csv' AS row
MERGE (p:Package {
  name: row.pkgname,
  base: row.pkgbase,
  version: row.pkgver,
  description: row.pkgdesc,
  URL: row.URL,
  filename: row.filename,
  compressedSize: toInteger(row.compressedSize),
  installedSize: toInteger(row.installedSize),
  buildDate: row.buildDate,
  packager: row.packager
})

MERGE (r:Repo {name: row.repo})
MERGE (a:Arch {name: row.arch})

MERGE (p)-[:IS_IN_REPO]->(r)
MERGE (p)-[:SUPPORTS_ARCHITECTURE]->(a)

RETURN *;

LOAD CSV WITH HEADERS FROM 'https://raw.githubusercontent.com/CodingCat12/pacgraph/refs/heads/main/packages/depends.csv' AS row
MATCH (p:Package {name: row.pkg})
MATCH (d:Package {name: row.depends})

MERGE (p)-[:DEPENDS_ON]->(d)

RETURN *;