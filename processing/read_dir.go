package processing

import (
	"alcoscanxlsx/utility"
	"fmt"
	"regexp"
)

func (k *Processing) ReadDir(re *regexp.Regexp, inDir string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic ReadXlsx %v", r)
		}
	}()

	// очищаем ошибки и предупреждения
	k.Reset()

	if k.files, err = utility.FilteredSearchOfDirectoryTree(re, inDir); err != nil {
		fmt.Print(err.Error())
		return err
	}
	return nil
}
