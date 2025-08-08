package utility

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

func ReadCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	// csvReader.Comma = '\t'
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func ReadTextStringArray(filePath string) []string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	arr := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return arr
}

func ReadTextStringArrayReader(file io.Reader) []string {
	arr := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return arr
}
