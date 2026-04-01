package tpls

import "embed"

//go:embed *.tmpl
var templateFS embed.FS

func ReadFile(name string) ([]byte, error) {
	return templateFS.ReadFile(name)
}
