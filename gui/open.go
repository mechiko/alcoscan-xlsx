package gui

import (
	"alcoscanxlsx/utility"
)

// открываем в браузере ссылку
func Open(url string) error {
	return utility.OpenHttpLinkInShell(url)
}

// открываем в эксплорере текущую папку программы
func OpenDir(dir string) (err error) {
	return utility.OpenFileInShell(dir)
}
