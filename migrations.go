package magisterlinguae

import "embed"

//go:embed sql/schema/*.sql
var MigrationsFS embed.FS
