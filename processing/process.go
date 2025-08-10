package processing

import (
	"alcoscanxlsx/processing/ucexcel"
	"fmt"
	"path/filepath"
)

const defaultNameOutput = "ПалетыКороба"

func (k *Processing) Proccess(name string) error {
	if name == "" {
		name = defaultNameOutput
	}
	name = filepath.Join(k.outDir, name)
	excel := ucexcel.New(name)
	if err := excel.Open(); err != nil {
		return fmt.Errorf("%w", err)
	}
	if err := excel.UtilizationPalet(k.Palet, k.Korob, k.PaletSort); err != nil {
		return fmt.Errorf("%w", err)
	}
	if err := excel.SaveSimple(); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
