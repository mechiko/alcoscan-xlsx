package processing

import (
	"alcoscanxlsx/processing/ucexcel"
	"fmt"
	"path/filepath"
)

func (k *Processing) Proccess(name string) error {
	if name == "" {
		return fmt.Errorf("имя файла не указано")
	}
	if !filepath.IsAbs(name) {
		return fmt.Errorf("имя файла не абсолютное %s", name)
	}
	excel := ucexcel.New(name)
	// так создается файл
	if excelFile, err := excel.Open(); err != nil {
		return fmt.Errorf("%w", err)
	} else {
		_ = excelFile.SetColWidth("Sheet1", "A", "D", 30)
		_ = excelFile.SetColWidth("Sheet1", "B", "B", 40)
	}

	if err := excel.UtilizationPalet(k.Palet, k.Korob, k.PaletSort); err != nil {
		return fmt.Errorf("%w", err)
	}
	if err := excel.ToAbsSimple(); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
