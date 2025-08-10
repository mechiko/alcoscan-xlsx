package ucexcel

import (
	"fmt"
	"os"
	_ "time/tzdata"
)

func (ue *ucexcel) ToFile() error {
	fname := ue.ExcelFileName(ue.name)
	file, err := os.OpenFile(fname, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0660)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer func() error {
		if err := file.Close(); err != nil {
			return fmt.Errorf("%w", err)
		}
		return nil
	}()

	if _, err := ue.file.WriteTo(file); err != nil {
		return fmt.Errorf("%w", err)
	}

	return err
}

func (ue *ucexcel) ToFileSimple() error {
	fname := ue.ExcelFileNameSimple(ue.name)
	file, err := os.OpenFile(fname, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0660)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer func() error {
		if err := file.Close(); err != nil {
			return fmt.Errorf("%w", err)
		}
		return nil
	}()

	if _, err := ue.file.WriteTo(file); err != nil {
		return fmt.Errorf("%w", err)
	}

	return err
}
