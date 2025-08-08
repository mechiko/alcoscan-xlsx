package utility

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

// Exists reports whether the named file or directory exists.
func PathOrFileExists(name string) (ret bool) {
	defer func() {
		if r := recover(); r != nil {
			ret = false
		}
	}()
	if _, err := os.Stat(name); err != nil {
		return false
	}
	return true
}

func AbsPathCreate(path string) error {
	if filepath.IsAbs(path) {
		if !PathOrFileExists(path) {
			if err := os.Mkdir(path, os.ModePerm); err != nil {
				os.Exit(1)
			}
		}
	}
	return fmt.Errorf("path not absolute")
}

func PathCreate(path string) error {
	if path != "" {
		if !PathOrFileExists(path) {
			if err := os.MkdirAll(path, os.ModePerm); err != nil {
				os.Exit(1)
			}
		}
	}
	return nil
}

func HomePathCreate(path string) error {
	home := UserHomeDir()
	if path != "" {
		fullPath := filepath.Join(home, path)
		if !PathOrFileExists(fullPath) {
			if err := os.MkdirAll(fullPath, os.ModePerm); err != nil {
				fmt.Printf("ошибка создания %s %s\n", fullPath, err.Error())
				os.Exit(1)
			}
		}
	} else {
		fmt.Println("пустой путь")
		os.Exit(1)
	}
	return nil
}

func UserHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		// Fallback to HOME environment variable
		return os.Getenv("HOME")
	}
	return home
}

func PathOrFileExistsMust(name string) (ret bool) {
	if _, err := os.Stat(name); err != nil {
		return false
	}
	return true
}

// Remove all non-ASCII characters
// Create string from string s, keeping only ASCII characters
func RemoveAllNonPrintable(s string) string {
	return strings.Map(func(r rune) rune {
		if r > unicode.MaxASCII {
			return -1
		}
		return r
	}, s)
}

// Remove all non-ASCII characters
// Create string from string s, keeping only ASCII characters
func RemoveAllNonNumber(s string) string {
	return strings.Map(func(r rune) rune {
		if r < '0' || r > '9' {
			return -1
		}
		return r
	}, s)
}

// BenchmarkRange-4    20000000    82.0 ns/op
// func isASCII(s string) bool {
//     for _, c := range s {
//         if c > unicode.MaxASCII {
//             return false
//         }
//     }
//     return true
// }

// BenchmarkIndex-4    30000000    55.4 ns/op
func IsASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

// строка содержит только цифры
func IsNumber(s string) bool {
	isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
	return strings.IndexFunc(s, isNotDigit) == -1
}

// строка содержит только цифры
func IsNumber2(s string) bool {
	b := true
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return b
}

// EX: re, err := regexp.Compile(`^0[0-9]{11}\.db$`)
//
//	files, err := FilteredSearchOfDirectoryTree(re, ""); err != nil {
//
// "^[a-zA-Z0-9].*\\.db$"
// `^[a-zA-Z0-9].*\.db$`
// глубина поиска 0 только в указанном каталоге
func FilteredSearchOfDirectoryTree(re *regexp.Regexp, dir string) ([]string, error) {
	absDir := dir
	if !filepath.IsAbs(absDir) {
		absDir, _ = filepath.Abs(absDir)
	}
	files := []string{}
	base := filepath.Base(absDir)
	walk := func(path string, d fs.DirEntry, err error) error {
		// каталоги не равные base пропускаем
		if d.IsDir() && (d.Name() != base) {
			return fs.SkipDir
		}
		if d.IsDir() {
			return nil
		}
		if !re.MatchString(d.Name()) {
			return nil
		}
		files = append(files, path)
		return nil
	}
	filepath.WalkDir(absDir, walk)
	return files, nil
}

func IndexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
