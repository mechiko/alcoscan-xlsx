package processing

import (
	"alcoscanxlsx/utility"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func (k *Processing) Scan() error {
	for _, file := range k.Files {
		if strings.Contains(file, "_АРМ") {
			if err := k.korobTxtFile(file); err != nil {
				return err
			}
		}
		if strings.Contains(file, "_Паллет") {
			if err := k.paletTxtFile(file); err != nil {
				return err
			}
		}
	}
	// найдем потерянные короба
	keysKorob := make([]string, 0, len(k.Korob))
	for k2 := range k.Korob {
		keysKorob = append(keysKorob, k2)
	}
	k.AllKorobaInPalet = make(map[string]string)
	for pal, value := range k.Palet {
		for _, kp := range value {
			_, ok := k.AllKorobaInPalet[kp]
			if ok {
				k.Logger().Errorf("in palet %s double korob kp", pal, kp)
			} else {
				k.AllKorobaInPalet[kp] = pal
			}
		}
	}
	for _, korob := range keysKorob {
		_, ok := k.AllKorobaInPalet[korob]
		if !ok {
			if _, ok := k.Palet[""]; !ok {
				k.Palet[""] = make([]string, 0)
			}
			k.Palet[""] = append(k.Palet[""], korob)
		}
	}
	// найдем короба без марок
	for krb, plt := range k.AllKorobaInPalet {
		if _, ok := k.Korob[krb]; !ok {
			return fmt.Errorf("палета %s содержит коробку %s по ней нет данных о КМ (файл АРМ_*)", plt, krb)
		}
	}
	for pp := range k.Palet {
		k.PaletSort = append(k.PaletSort, pp)
	}
	slices.Sort(k.PaletSort)
	return nil
}

func (k *Processing) paletTxtFile(file string) error {
	arrPalet, err := readStringCsv(file)
	mkrb := map[string]bool{}
	if err != nil {
		return err
	}
	for _, row := range arrPalet {
		plt := row[0]
		krb := row[1]
		if _, ok := mkrb[krb]; ok {
			k.Logger().Errorf("double korob %s in palet %s", krb, plt)
		}
		mkrb[krb] = true
		if cis, err := utility.ParseCisInfo(krb); err == nil {
			krb = cis.Cis
		} else {
			k.Logger().Errorf("parse palet %s korob cis [%s] error %s", plt, krb, err.Error())
		}
		if _, ok := k.Palet[plt]; !ok {
			k.Palet[plt] = make([]string, 0)
		}
		k.Palet[plt] = append(k.Palet[plt], krb)
	}
	return nil
}

func (k *Processing) korobTxtFile(file string) error {
	arrKorob, err := readStringCsv(file)
	if err != nil {
		return err
	}
	k.Logger().Infof("считано %d из файла %s", len(arrKorob), file)
	for i, row := range arrKorob {
		cis, err := utility.ParseCisInfo(row[1])
		if err != nil {
			return fmt.Errorf("файл [%.20s] номер строки %d [%s] %w", filepath.Base(file), i+1, row[1], err)
		}
		krb := row[0]
		cisKorob, err := utility.ParseCisInfo(row[0])
		if err != nil {
			return fmt.Errorf("файл [%.20s] номер строки %d [%s] %w", filepath.Base(file), i+1, row[0], err)
		}
		if cisKorob != nil {
			krb = cisKorob.Cis
		}
		if _, ok := k.Korob[krb]; !ok {
			k.Korob[krb] = make([]*utility.CisInfo, 0)
		}
		k.Korob[krb] = append(k.Korob[krb], cis)
	}
	return nil
}

func readStringCsv(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	arr := make([][]string, 0)
	scanner := bufio.NewScanner(f)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		// txt := utility.RemoveNonAscii(scanner.Text())
		txt := scanner.Text()
		row := strings.Split(txt, "\t")
		if len(row) != 2 {
			return nil, fmt.Errorf("в строке [%s] должно быть два значения разделенных табуляцией", txt)
		}
		arr = append(arr, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return arr, nil
}
