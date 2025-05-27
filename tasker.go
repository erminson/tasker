package tasker

import "embed"

const ApplicationName = "tasker"

//go:embed migrations/*.sql
var Migrations embed.FS
