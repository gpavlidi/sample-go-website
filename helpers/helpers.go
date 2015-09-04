package helpers

import (
	"path/filepath"
)

func GetSharedTemplates() []string {
	layouts, _ := filepath.Glob("views/layouts/*.html")
	shared, _ := filepath.Glob("views/shared/*.html")

	return append(layouts, shared...)
}
