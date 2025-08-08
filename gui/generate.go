package gui

import (
	"fmt"
)

// кнопка Пуск
// запускать в отдельном поток от tk9
func (a *GuiApp) generate() {
	defer func() {
		a.stateIsProcess <- false
		// по завершению обработки в БД кнопка Пуск запрещена
		a.stateFinish <- struct{}{}
	}()
	a.stateIsProcess <- true
	model, err := GetModel()
	if err != nil {
		a.Logger().Errorf("gui generate %w", err)
		a.SendError(fmt.Sprintf("gui generate %w", err))
		return
	}
	if model.Magazin == "" {
		a.SendError("магазин не указан")
		return
	}
	orders, ok := model.Reestr[model.Magazin]
	if !ok {
		a.SendError(fmt.Sprintf("для магазина %s нет заказов", model.Magazin))
		return
	}

	// запись заказа в БД
	// if err := a.processing.Write(); err != nil {
	// 	a.SendError(err.Error())
	// 	return
	// }
	// после успешной записи
	a.SendLog(fmt.Sprintf("%d заказов на маркировку по магазину %s записаны в БД", len(orders), model.Magazin))
}
