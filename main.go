package main

import (
	"encoding/json"
	"fmt"

	"time"
)

func main() {

	startTime := time.Now()

	certFiles := make([]CertFile, 2)

	var err error

	var filePath string
	filePath = `C:\gowstemp\bin\file1.txt`
	certFiles[0].ProcessedFile = filePath
	err = certFiles[0].Load(filePath, '\t')
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(filePath, time.Since(startTime), len(certFiles[0].CertRecords), "rec")
	startTime = time.Now()

	filePath = `C:\gowstemp\bin\file3.txt`
	certFiles[1].ProcessedFile = filePath

	err = certFiles[1].Load(filePath, '\t')
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(filePath, time.Since(startTime), len(certFiles[1].CertRecords), "rec")

	dall := deltasAll{}

	for i := range certFiles {

		if i == len(certFiles)-1 {
			break
		}

		ds := deltas{}
		ds.oldFileName = certFiles[i].ProcessedFile
		ds.newFileName = certFiles[i+1].ProcessedFile

		for j, recs := range certFiles[i].CertRecords {
			for k, _ := range recs {

				if certFiles[i].CertRecords[j][k].CertType !=
					certFiles[i+1].CertRecords[j][k].CertType {
					d := delta{}
					d.key = k
					d.fieldName = "CertType"
					d.oldValue = certFiles[i].CertRecords[j][k].CertType
					d.newValue = certFiles[i+1].CertRecords[j][k].CertType
					ds.delta = append(ds.delta, d)
				}

			}
		}
		dall = append(dall, ds)

	}

	for i, d := range dall {

		fmt.Println(i, d.oldFileName, d.newFileName)
		for j, diff := range d.delta {
			b, err := json.Marshal(diff)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("	!!!!!!", j, string(b))
			fmt.Println("key:", diff.key)
			fmt.Println("field:", diff.fieldName)
			fmt.Println("old:", diff.oldValue)
			fmt.Println("new:", diff.newValue)
		}
	}

}
