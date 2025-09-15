package utility

import (
	"errors"
	"fmt"
	"os"

	"github.com/sqweek/dialog"
)

type FileType string

const (
	Excel FileType = "xlsx"
	Txt   FileType = "txt"
	Csv   FileType = "csv"
	DB    FileType = "db"
	All   FileType = "all"
)

// func isValidFileType(s FileType) bool {
// 	switch s {
// 	case Excel, Csv, Txt:
// 		return true
// 	default:
// 		return false
// 	}
// }

func DialogOpenFile(fileType []FileType, name string, wd string) (string, error) {
	if wd == "." {
		dir, err := os.Getwd()
		if err != nil {
			return "", err
		}
		wd = dir
	}
	dlg := dialog.File().SetStartDir(wd).SetStartFile(name)
	for _, t := range fileType {
		switch t {
		case Excel:
			dlg.Filter("Excel", "xlsx")
		case Csv:
			dlg.Filter("CSV", "csv")
		case Txt:
			dlg.Filter("Txt", "txt")
		case All:
			dlg.Filter("All", "*")
		}
	}
	result, err := dlg.Load()
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	if result == "Cancelled" {
		return "", fmt.Errorf("диалог выбора отменен")
	}
	return result, nil
}

func DialogSaveFile(fileType FileType, name string, wd string) (string, error) {
	if wd == "." {
		dir, err := os.Getwd()
		if err != nil {
			return "", err
		}
		wd = dir
	}
	dlg := dialog.File().SetStartDir(wd).SetStartFile(name)
	switch fileType {
	case Excel:
		dlg.Filter("Excel", "xlsx")
	case Csv:
		dlg.Filter("CSV", "csv")
	case Txt:
		dlg.Filter("Txt", "txt")
	}
	result, err := dlg.Save()
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	if result == "Cancelled" {
		return "", fmt.Errorf("диалог выбора отменен")
	}
	return result, nil
}

func MessageBox(title, msg string) {
	if msg == "" {
		return
	}
	dialog.Message("%s", msg).Title(title).Info()
}

func DialogSelectDir(wd string) (string, error) {
	if wd == "." {
		dir, err := os.Getwd()
		if err != nil {
			return "", err
		}
		wd = dir
	}
	directory, err := dialog.Directory().SetStartDir(wd).Title("Выберите папку с файлами данных").Browse()
	if err != nil {
		return "", err
	}
	if directory == "Cancelled" {
		return "", errors.New("прерван диалог выбора каталога")
	}
	return directory, nil
}
