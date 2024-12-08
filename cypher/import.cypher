CREATE INDEX IF NOT EXISTS FOR (p:Package) ON (p.name);
CREATE INDEX IF NOT EXISTS FOR (p:Package) ON (p.base);
CREATE INDEX IF NOT EXISTS FOR (p:Person) ON (p.name);
CREATE INDEX IF NOT EXISTS FOR (r:Repo) ON (r.name);
CREATE INDEX IF NOT EXISTS FOR (a:Arch) ON (a.name);
CREATE INDEX IF NOT EXISTS FOR (g:Group) ON (g.name);
CREATE INDEX IF NOT EXISTS FOR (l:License) ON (l.name);
CREATE INDEX IF NOT EXISTS FOR (v:VirtualPackage) ON (v.name);

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

MERGE (pe:Person {name: row.packager})
MERGE (r:Repo {name: row.repo})
MERGE (a:Arch {name: row.arch})
MERGE (p)-[:PACKAGED_BY]->(pe)
MERGE (p)-[:IS_IN_REPO]->(r)
MERGE (p)-[:SUPPORTS_ARCHITECTURE]->(a)

RETURN *;

LOAD CSV WITH HEADERS FROM 'https://raw.githubusercontent.com/CodingCat12/pacgraph/refs/heads/main/packages/depends.csv' AS row
MATCH (p:Package {name: row.pkg})
MATCH (d:Package {name: row.depends})
MERGE (p)-[:DEPENDS_ON]->(d)
RETURN *;

LOAD CSV WITH HEADERS FROM 'https://raw.githubusercontent.com/CodingCat12/pacgraph/refs/heads/main/packages/optdepends.csv' AS row
MATCH (p:Package {name: row.pkg})
MATCH (d:Package {name: row.optdepends})
MERGE (p)-[:OPTDEPENDS_ON]->(d)
RETURN *;

LOAD CSV WITH HEADERS FROM 'https://raw.githubusercontent.com/CodingCat12/pacgraph/refs/heads/main/packages/makedepends.csv' AS row
MATCH (p:Package {name: row.pkg})
MATCH (d:Package {name: row.makedepends})
MERGE (p)-[:MAKEDEPENDS_ON]->(d)
RETURN *;

LOAD CSV WITH HEADERS FROM 'https://raw.githubusercontent.com/CodingCat12/pacgraph/refs/heads/main/packages/checkdepends.csv' AS row
MATCH (p:Package {name: row.pkg})
MATCH (d:Package {name: row.checkdepends})
MERGE (p)-[:CHECKDEPENDS_ON]->(d)
RETURN *;

LOAD CSV WITH HEADERS FROM 'https://raw.githubusercontent.com/CodingCat12/pacgraph/refs/heads/main/packages/conflicts.csv' AS row
MATCH (p:Package {name: row.pkg})
MATCH (d:Package {name: row.conflicts})
MERGE (p)-[:CONFLICTS_WITH]->(d)
RETURN *;

LOAD CSV WITH HEADERS FROM 'https://raw.githubusercontent.com/CodingCat12/pacgraph/refs/heads/main/packages/groups.csv' AS row
MATCH (p:Package {name: row.pkg})
MERGE (d:Group {name: row.groups})
MERGE (p)-[:IS_IN_GROUP]->(d)
RETURN *;

LOAD CSV WITH HEADERS FROM 'https://raw.githubusercontent.com/CodingCat12/pacgraph/refs/heads/main/packages/licenses.csv' AS row
MATCH (p:Package {name: row.pkg})
MERGE (d:License {name: row.licenses})
MERGE (p)-[:USES_LICENSE]->(d)
RETURN *;

LOAD CSV WITH HEADERS FROM 'https://raw.githubusercontent.com/CodingCat12/pacgraph/refs/heads/main/packages/replaces.csv' AS row
MATCH (p:Package {name: row.pkg})
MATCH (d:Package {name: row.replaces})
MERGE (p)-[:REPLACES]->(d)
RETURN *;

LOAD CSV WITH HEADERS FROM 'https://raw.githubusercontent.com/CodingCat12/pacgraph/refs/heads/main/packages/provides.csv' AS row
MATCH (p:Package {name: row.pkg})
MERGE (d:VirtualPackage {name: row.provides})
MERGE (p)-[:PROVIDES]->(d)
RETURN *;