package processing

import (
	"alcoscanxlsx/domain/models/application"
	"alcoscanxlsx/reductor"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func (k *Processing) ReadXlsx() (m *application.Application, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic ReadXlsx %v", r)
		}
	}()

	// очищаем ошибки и предупреждения
	k.Reset()
	model, err := k.GetModel()
	if err != nil {
		return model, fmt.Errorf("processing write %w", err)
	}
	f, err := excelize.OpenFile(model.File)
	if err != nil {
		return m, fmt.Errorf("open xlsx error %w", err)
	}
	defer func() {
		// Close the spreadsheet.
		if errr := f.Close(); errr != nil {
			err = errr
		}
	}()

	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return model, fmt.Errorf("len sheet wrong")
	}
	currentSheet := sheets[0]

	// Получить все строки в Sheet1
	rows, err := f.GetRows(currentSheet)
	if err != nil {
		return model, fmt.Errorf("%w", err)
	}

	prevZakaz := application.Zakaz{}
	for rowNumber, row := range rows {
		if len(row) == 0 {
			continue
		}
		if rowNumber == 0 {
			continue
		}
		if k.isRecord(row) {
			// берем запись
			zakaz := parseZakaz(row)
			if zakaz.Magazin == "" {
				zakaz.Magazin = prevZakaz.Magazin
				zakaz.ID = prevZakaz.ID
				zakaz.Gtin = prevZakaz.Gtin
			}
			if _, ok := model.Reestr[zakaz.Magazin]; !ok {
				model.Reestr[zakaz.Magazin] = make([]*application.Zakaz, 0)
				model.Magazins = append(model.Magazins, zakaz.Magazin)
			}
			model.Reestr[zakaz.Magazin] = append(model.Reestr[zakaz.Magazin], zakaz)
			prevZakaz = *zakaz
		}
	}
	reductor.Instance().SetModel(model, false)
	return model, nil
}

// полная строка 5 ячеек
func (k *Processing) isRecord(row []string) bool {
	return len(row) > 8 && row[8] != ""
}

func parseZakaz(row []string) *application.Zakaz {
	if len(row) <= 8 || row[8] == "" {
		return &application.Zakaz{} // or return error
	}
	return &application.Zakaz{
		Magazin:  row[2],
		ID:       row[1],
		Gtin:     "0" + row[7],
		Quantity: row[8],
	}
}
