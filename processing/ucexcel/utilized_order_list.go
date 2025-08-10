package ucexcel

import (
	"alcoscanxlsx/processing/ucexcel/address"
	"alcoscanxlsx/utility"
	"fmt"
	"slices"
)

func (ue *ucexcel) UtilizationPalet(palet map[string][]string, korob map[string][]*utility.CisInfo, paletSort []string) error {
	ue.sheet = "Sheet1"
	// countRow = 1
	ue.address = address.New(1, 0)
	ue.templateLineHeaderPalet()
	for _, plt := range paletSort {
		// палет
		krbs := palet[plt]
		krbSort := make([]string, len(krbs))
		copy(krbSort, krbs)
		slices.Sort(krbSort)
		for _, krb := range krbSort {
			// короб в палете
			if cises, ok := korob[krb]; ok {
				for _, cis := range cises {
					ue.templateLinePalet(cis, plt, krb)
				}
			}

		}
	}
	return nil
}

func (ue *ucexcel) UtilizationReportList(report [][]string) error {
	ue.sheet = "Sheet1"
	// countRow = 1
	ue.address = address.New(1, 0)
	for _, ss := range report {
		ue.templateReportListLine(ss)
	}

	return nil
}

func (ue *ucexcel) UtilizationTxt(report []string) error {
	ue.sheet = "Sheet1"
	// countRow = 1
	ue.address = address.New(1, 0)
	for _, ss := range report {
		ue.templateLine(ss)
	}

	return nil
}

func (ue *ucexcel) templateReportListLine(s []string) {
	addr := ue.address.Address()
	out := s[0]
	outstr := out[:25]
	if err := ue.file.SetCellStr(ue.sheet, addr, outstr); err != nil {
		fmt.Printf("excel error %s", err.Error())
	}
	ue.address.NextRow()
}

func (ue *ucexcel) templateLine(s string) {
	addr := ue.address.Address()
	outstr := s[:25]
	if err := ue.file.SetCellStr(ue.sheet, addr, outstr); err != nil {
		fmt.Printf("excel error %s", err.Error())
	}
	ue.address.NextRow()
}

func (ue *ucexcel) templateLinePalet(cis *utility.CisInfo, palet string, korob string) {
	if err := ue.file.SetCellStr(ue.sheet, ue.address.Address(), cis.Cis); err != nil {
		fmt.Printf("excel error %s", err.Error())
	}
	if err := ue.file.SetCellStr(ue.sheet, ue.address.NextCol(), cis.Code); err != nil {
		fmt.Printf("excel error %s", err.Error())
	}
	if err := ue.file.SetCellStr(ue.sheet, ue.address.NextCol(), korob); err != nil {
		fmt.Printf("excel error %s", err.Error())
	}
	if err := ue.file.SetCellStr(ue.sheet, ue.address.NextCol(), palet); err != nil {
		fmt.Printf("excel error %s", err.Error())
	}
	ue.address.NextRow()
}

func (ue *ucexcel) templateLineHeaderPalet() {
	if err := ue.file.SetCellStr(ue.sheet, ue.address.Address(), "Серийный номер"); err != nil {
		fmt.Printf("excel error %s", err.Error())
	}
	if err := ue.file.SetCellStr(ue.sheet, ue.address.NextCol(), "Код Маркировки"); err != nil {
		fmt.Printf("excel error %s", err.Error())
	}
	if err := ue.file.SetCellStr(ue.sheet, ue.address.NextCol(), "SSCC код короба"); err != nil {
		fmt.Printf("excel error %s", err.Error())
	}
	if err := ue.file.SetCellStr(ue.sheet, ue.address.NextCol(), "SSCC код паллеты"); err != nil {
		fmt.Printf("excel error %s", err.Error())
	}
	ue.address.NextRow()
}
