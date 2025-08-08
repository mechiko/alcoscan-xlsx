package utility

import (
	"path/filepath"
	"regexp"
	"strings"
)

// At top of file, alongside other package-level declarations
var fileNameRegex = regexp.MustCompile(`[^\p{L}\p{N} ]+`)

func ClearForFileName(str string) string {
	str = fileNameRegex.ReplaceAllString(str, "")
	str = strings.TrimSpace(str)
	if str == "" {
		return "unnamed"
	}
	return str
}

func FileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(filepath.Base(fileName), filepath.Ext(fileName))
}
