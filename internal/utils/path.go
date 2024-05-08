package utils

import (
	"kedr/internal/config"
	"os"
	"path/filepath"
)

// GetTemplatePath путь к шаблону
func GetTemplatePath(template string) string {
	rootPath, _ := os.Getwd()
	return filepath.Join(rootPath, config.DirWeb, config.DirTemplates, template)
}
