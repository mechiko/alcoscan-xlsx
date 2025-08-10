package gui

// кнопка Пуск
// запускать в отдельном поток от tk9
func (a *GuiApp) generate() {
	defer func() {
		a.stateIsProcess <- false
		// по завершению обработки в БД кнопка Пуск запрещена
		a.stateFinish <- struct{}{}
	}()
	a.stateIsProcess <- true
	a.SendLog("")
}
