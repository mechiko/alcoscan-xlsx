package gui

import (
	"alcoscanxlsx/utility"
	"fmt"
)

const defaultNameOutput = "Палеты_Короба.xlsx"

// кнопка Пуск
// запускать в отдельном поток от tk9
func (a *GuiApp) generate() {
	defer func() {
		a.stateIsProcess <- false
	}()
	a.stateIsProcess <- true

	a.SendLog("обрабатываем файлы...")
	err := a.processing.Scan()
	if err != nil {
		a.SendError(fmt.Sprintf("ошибка загрузки файлов: %s", err.Error()))
		a.stateStart <- struct{}{}
		return
	}
	a.SendLog(fmt.Sprintf("загружены %d короба(ов)", len(a.processing.Korob)))
	a.SendLog(fmt.Sprintf("загружены %d палет", len(a.processing.Palet)))

	if ff, err := utility.DialogSaveFile(utility.Excel, defaultNameOutput, "."); err != nil {
		a.SendError(err.Error())
		a.stateStart <- struct{}{}
		return
	} else {
		err = a.processing.Proccess(ff)
		if err != nil {
			a.SendError(err.Error())
			a.stateStart <- struct{}{}
			return
		}
		if err := OpenDir(ff); err != nil {
			a.SendError(err.Error())
		}
	}
	// по завершению обработки в БД кнопка Пуск запрещена
	a.stateFinish <- struct{}{}
}
