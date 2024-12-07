LOAD CSV WITH HEADERS FROM 'https://raw.githubusercontent.com/CodingCat12/pacgraph/refs/heads/main/packages/depends.csv' AS row
MERGE (p:Package {name: row.pkg})
MERGE (d:Package {name: row.depends})

MERGE (p)-[:DEPENDS_ON]->(d)

RETURN *