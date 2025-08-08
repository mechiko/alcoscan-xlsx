package utility

import (
	"fmt"
	"time"
	_ "time/tzdata"
)

// TimeFileName generates a timestamped filename by combining the given name
// with the current date and time in the format: name_YYYY.MM.DD_HHMMSS
// The name parameter should not contain filesystem-unsafe characters.
func TimeFileName(name string) string {
	n := time.Now().Local()
	return fmt.Sprintf("%s_%4d.%02d.%02d_%02d%02d%02d", name, n.Year(), n.Month(), n.Day(), n.Hour(), n.Minute(), n.Second())
}
