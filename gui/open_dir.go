package gui

import (
	"alcoscanxlsx/reductor"
	"fmt"
	"path/filepath"
	"regexp"
)

var reInDir = regexp.MustCompile(`.*\.csv$`)

// должна выполнятся как gorutine
func (a *GuiApp) openInDir(ff string) {
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
	model.InDir = ff
	err = reductor.Instance().SetModel(model, false)
	if err != nil {
		a.Logger().Errorf("gui openXlsx %w", err)
	}
	a.SendLog("считываем файлы палет и коробов")
	err = a.processing.ReadDir(reInDir, model.InDir)
	if err != nil {
		a.SendError(fmt.Sprintf("ошибка загрузки файлов: %s", err.Error()))
		a.stateStart <- struct{}{}
		return
	}
	a.SendLog("обрабатываем файлы...")
	err = a.processing.Scan()
	if err != nil {
		a.SendError(fmt.Sprintf("ошибка загрузки файлов: %s", err.Error()))
		if len(a.processing.Errors()) > 0 {
			// ошибки проверки данных таблицы выводим в лог
			for _, e := range a.processing.Errors() {
				a.SendError(e)
			}
			a.stateStart <- struct{}{}
			return
		}
	}
	a.SendLog("обработаны файлы")
	// успешное открытие файла
	a.stateStart <- struct{}{}
}
