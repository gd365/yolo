package service

import (
	"embed"
	"io/fs"
)

//go:embed all:templates
var templateFS embed.FS

func GetTemplateFS() fs.FS {
	sub, _ := fs.Sub(templateFS, "templates")
	return sub
}
