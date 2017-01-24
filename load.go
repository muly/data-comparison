package main

import (
	"encoding/csv"
	"os"
	//"fmt"
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

func (certs *CertRecords) Load(filePath string, delim rune) error {

	records, err := readFile(filePath, delim)
	if err != nil {
		return err
	}

	var CertID, CertificationType, OriginalGrantDate, CertEndDate, CertStatus int
	for i, record := range records {
		if i == 0 {
			for j, c := range record {
				//find the position of each field
				switch c {
				case "CertID":
					CertID = j
				case "CertificationType":
					CertificationType = j
				case "OriginalGrantDate":
					OriginalGrantDate = j
				case "CertEndDate":
					CertEndDate = j
				case "CertStatus":
					CertStatus = j
				}
			}
			continue

		}

		k := CertKey{}
		d := CertData{}

		k.CertID = record[CertID]

		d.CertEndDate = record[CertEndDate]
		d.CertStatus = record[CertStatus]
		d.CertType = record[CertificationType]
		d.OriginalGrantDate = record[OriginalGrantDate]

		c := CertRecord{}

		//fmt.Printf("%# v\n", pretty.Formatter(d))

		c[k] = d

		*certs = append(*certs, c)
	}

	/*	for i, d := range *certs {
		fmt.Println()
		fmt.Println(i)
		for k, v := range d {
			fmt.Println(k, v)
		}
	} */

	return nil
}
