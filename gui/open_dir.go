package gui

import (
	"alcoscanxlsx/reductor"
	"fmt"
	"path/filepath"
	"regexp"
)

var reInDir = regexp.MustCompile(`.*\.csv$`)

// должна выполнятся как gorutine
func (a *GuiApp) openInDir(inDir string) {
	// очистка лога на экране
	a.logClear <- struct{}{}
	a.stateIsProcess <- true
	defer func() {
		a.stateIsProcess <- false
	}()
	model, err := GetModel()
	if err != nil {
		a.Logger().Errorf("gui openXlsx %w", err)
		a.SendError(fmt.Sprintf("gui openXlsx %w", err))
		a.stateStart <- struct{}{}
		return
	}
	// сброс модели
	model.Reset()
	model.InDir = inDir
	err = reductor.Instance().SetModel(model, false)
	if err != nil {
		a.Logger().Errorf("gui openXlsx %w", err)
		a.SendError(fmt.Sprintf("ошибка загрузки файлов: %s", err.Error()))
		a.stateStart <- struct{}{}
		return
	}
	a.SendLog("считываем файлы палет и коробов")
	err = a.processing.ReadDir(reInDir, model.InDir)
	if err != nil {
		a.SendError(fmt.Sprintf("ошибка загрузки файлов: %s", err.Error()))
		a.stateStart <- struct{}{}
		return
	}
	if len(a.processing.Files) < 2 {
		// должны быть хотя бы два файла, короба и палеты
		a.SendError(fmt.Sprintf("ошибка мало файлов, найдено только %d", len(a.processing.Files)))
		a.stateStart <- struct{}{}
		return
	}
	a.SendLog(fmt.Sprintf("найдены %d файла(ы):", len(a.processing.Files)))
	for _, file := range a.processing.Files {
		baseFile := filepath.Base(file)
		a.SendLog(fmt.Sprintf("- %.50s...%s", baseFile, filepath.Ext(baseFile)))
		// a.SendLog(fmt.Sprintf("- %s", baseFile))
	}
	// устанавливаем состояни для пуск
	a.stateSelectedInDir <- filepath.Base(inDir)
}
