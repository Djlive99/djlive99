package embed

import "embed"

// AcmeSh script
//go:embed acme.sh
var AcmeSh string

// APIDocFiles contain all the files used for swagger schema generation
//go:embed api_docs
var APIDocFiles embed.FS

// Assets are frontend assets served from within this app
//go:embed assets
var Assets embed.FS

// MigrationFiles are database migrations
//go:embed migrations/*.sql
var MigrationFiles embed.FS

// NginxFiles hold nginx config templates
//go:embed nginx
var NginxFiles embed.FS
