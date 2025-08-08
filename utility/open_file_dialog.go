package utility

import (
	"github.com/sqweek/dialog"
)

func DialogOpenFileXlsx() string {
	result, err := dialog.File().Filter("Excel", "xlsx").Filter("all", "*").Load()
	if err != nil {
		return err.Error()
	}
	return result
}

func DialogOpenFileCsv() string {
	result, err := dialog.File().Filter("csv", "csv").Filter("txt", "txt").Filter("all", "*").Load()
	if err != nil {
		return err.Error()
	}
	return result
}

func DialogOpenFileDb() string {
	result, err := dialog.File().Filter("db", "db").Filter("all", "*").Load()
	if err != nil {
		return err.Error()
	}
	return result
}

func DialogOpenFileTxt() string {
	result, err := dialog.File().Filter("txt", "txt").Filter("all", "*").Load()
	if err != nil {
		return err.Error()
	}
	return result
}

func DialogSaveFile() string {
	result, err := dialog.File().Filter("Excel", "xlsx").Filter("all", "*").Save()
	if err != nil {
		return err.Error()
	}
	return result
}

func MessageBox(title, msg string) {
	dialog.Message("%s", msg).Title(title).Info()
}
