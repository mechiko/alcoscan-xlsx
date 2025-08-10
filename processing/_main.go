package main

import (
	"bufio"
	"csv2xls/ucexcel"
	"csv2xls/utility"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
)

const inDir = `IN`
const outDir = `OUT`

var korob = map[string][]*utility.CisInfo{}
var palet = map[string][]string{}
var paletSort []string

func main() {
	var err error
	var files []string

	re, err := regexp.Compile(`.*\.csv$`)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(-1)
	}
	if files, err = utility.FilteredSearchOfDirectoryTree(re, inDir); err != nil {
		fmt.Print(err.Error())
		os.Exit(-1)
	}
	for _, file := range files {
		if strings.Contains(file, "_АРМ") {
			if err := korobTxtFile(file); err != nil {
				log.Fatal(err)
			}
		}
		if strings.Contains(file, "_Паллет") {
			if err := paletTxtFile(file); err != nil {
				log.Fatal(err)
			}
		}
	}
	paletSort = make([]string, 0, len(palet))
	for k := range palet {
		paletSort = append(paletSort, k)
	}
	slices.Sort(paletSort)
	fmt.Println(len(palet))
	if err := proccess(); err != nil {
		fmt.Print(err.Error())
		os.Exit(-1)
	}
}

func proccess() error {
	name := "output"
	name = filepath.Join(outDir, name)
	excel := ucexcel.New(name)
	if err := excel.Open(); err != nil {
		return fmt.Errorf("%w", err)
	}
	if err := excel.UtilizationPalet(palet, korob, paletSort); err != nil {
		return fmt.Errorf("%w", err)
	}
	if err := excel.SaveSimple(); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func paletTxtFile(file string) error {
	arrPalet, err := readStringCsv(file)
	if err != nil {
		return err
	}
	for _, row := range arrPalet {
		plt := row[0]
		krb := row[1]
		if _, ok := palet[plt]; !ok {
			palet[plt] = make([]string, 0)
		}
		palet[plt] = append(palet[plt], krb)
	}
	return nil
}

func korobTxtFile(file string) error {
	arrKorob, err := readStringCsv(file)
	if err != nil {
		return err
	}
	for i, row := range arrKorob {
		cis, err := utility.ParseCisInfo(row[1])
		if err != nil {
			return fmt.Errorf("файл [%s] номер строки %d [%s]", file, i+1, row[1])
		}
		krb := row[0]
		cisKorob, _ := utility.ParseCisInfo(row[0])
		if cisKorob != nil {
			krb = cisKorob.Cis
		}
		if _, ok := korob[krb]; !ok {
			korob[krb] = make([]*utility.CisInfo, 0)
		}
		korob[krb] = append(korob[krb], cis)
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
		txt := utility.RemoveNonAscii(scanner.Text())
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
