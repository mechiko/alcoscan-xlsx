package gui

import (
	"alcoscanxlsx/reductor"
	"fmt"
	"path/filepath"
)

// должна выполнятся как gorutine
func (a *GuiApp) openXlsx(ff string) {
	// очистка лога на экране
	a.logClear <- struct{}{}
	a.stateIsProcess <- true
	defer func() {
		a.stateIsProcess <- false
	}()
	file := filepath.Base(ff)
	a.stateSelectedFileXlsx <- file
	model, err := GetModel()
	if err != nil {
		a.Logger().Errorf("gui openXlsx %w", err)
		a.SendError(fmt.Sprintf("gui openXlsx %w", err))
		a.stateStart <- struct{}{}
		return
	}
	// сброс модели
	model.Reset()
	model.File = ff
	err = reductor.Instance().SetModel(model, false)
	if err != nil {
		a.Logger().Errorf("gui openXlsx %w", err)
	}
	a.SendLog("открываем файл данных")
	model, err = a.processing.ReadXlsx()
	if err != nil {
		a.SendError(fmt.Sprintf("ошибка загрузки заказа: %s", err.Error()))
		a.stateStart <- struct{}{}
		return
	}
	a.SendLog("проверяем заказы...")
	a.processing.Scan()
	if len(a.processing.Errors()) > 0 {
		// ошибки проверки данных таблицы выводим в лог
		for _, e := range a.processing.Errors() {
			a.SendError(e)
		}
		a.stateStart <- struct{}{}
		return
	}
	a.SendLog(fmt.Sprintf("список кодов заказа загружен %d магазинов", len(model.Magazins)))
	for _, magazin := range model.Magazins {
		a.SendLog(fmt.Sprintf("магазин %s заказов %d", magazin, len(model.Reestr[magazin])))
	}
	// успешное открытие файла
	a.stateStart <- struct{}{}
}
