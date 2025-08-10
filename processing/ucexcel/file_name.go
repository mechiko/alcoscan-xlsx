package ucexcel

import (
	"alcoscanxlsx/utility"
	"fmt"
	"path/filepath"
	"time"
	_ "time/tzdata"
)

const ext = "xlsx"

func (ue *ucexcel) ExcelFileName(prefix string) string {
	// ds := time.Now()
	prefix = filepath.Base(prefix)
	randName := utility.String(10)
	// startDate := fmt.Sprintf("%02d.%02d.%4d", ds.Local().Day(), ds.Local().Month(), ds.Local().Year())
	// endDate := fmt.Sprintf("%02d.%02d.%4d", ds.Local().Day(), ds.Local().Month(), ds.Local().Year())
	// guid := utility.String(8)
	name := prefix + "_" + randName
	name += "." + ext
	return filepath.Join("output", name)
}

func (ue *ucexcel) ExcelFileNameDownload(prefix string) string {
	ds := time.Now()
	startDate := fmt.Sprintf("%02d.%02d.%4d", ds.Local().Day(), ds.Local().Month(), ds.Local().Year())
	endDate := fmt.Sprintf("%02d.%02d.%4d", ds.Local().Day(), ds.Local().Month(), ds.Local().Year())
	name := prefix + "_" + startDate + "_" + endDate
	name += "." + ext
	return name
}

func (ue *ucexcel) ExcelFileNameSimple(prefix string) string {
	ds := time.Now()
	startDate := fmt.Sprintf("%02d.%02d.%4d", ds.Local().Day(), ds.Local().Month(), ds.Local().Year())
	name := prefix + "_" + startDate
	name += "." + ext
	return filepath.Join(name)
}
