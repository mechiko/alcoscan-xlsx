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
	for pp := range k.Palet {
		k.PaletSort = append(k.PaletSort, pp)
	}
	slices.Sort(k.PaletSort)
	return nil
}

func (k *Processing) paletTxtFile(file string) error {
	arrPalet, err := readStringCsv(file)
	if err != nil {
		return err
	}
	for _, row := range arrPalet {
		plt := row[0]
		krb := row[1]
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
	for i, row := range arrKorob {
		cis, err := utility.ParseCisInfo(row[1])
		if err != nil {
			return fmt.Errorf("файл [%.20s] номер строки %d [%s] %w", filepath.Base(file), i+1, row[1], err)
		}
		krb := row[0]
		cisKorob, _ := utility.ParseCisInfo(row[0])
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
