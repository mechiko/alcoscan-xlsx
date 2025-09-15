package ucexcel

import (
	"github.com/xuri/excelize/v2"
)

func (ue *ucexcel) Open() (*excelize.File, error) {
	// создаем чистый файл без стилей
	ue.file = excelize.NewFile()
	return ue.file, nil
}
