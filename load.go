package main

import (
	"encoding/csv"
	"os"

	"fmt"
	"strconv"
	//"github.com/kr/pretty"
)

func readFile(filePath string, delim rune) (records [][]string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = delim
	r.Comment = '#'

	records, err = r.ReadAll()
	if err != nil {
		return
	}

	return
}

func load(filePath string, delim rune, keys []string) (t table, err error) {
	t.rows = rows{}

	t.sourceFileName = filePath

	records, err := readFile(filePath, delim)
	if err != nil {
		return
	}

	for i, record := range records {
		if i == 0 { // skip the header row
			continue
		}

		r := row{}
		for j, _ := range record {
			//if records[0][j] in keys{// handle keys}

			r[records[0][j]] = records[i][j]
		}

		t.rows[strconv.Itoa(i)] = r //TODO: key needs to be change from index to actual key that user passes to the program

	}

	/*for k, r := range t.rows {
		fmt.Println()
		fmt.Println(k)
		for c, v := range r {
			fmt.Println(c, v)
		}
	}*/

	return
}

func (t table) print() {
	fmt.Println("Showing data from ", t.sourceFileName)
	for k, r := range t.rows {
		fmt.Print("\t[", k, "]")
		for c, v := range r {
			fmt.Print("\t", c, ":", v)
		}
		println()
	}

}
