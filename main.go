package main

import (
	//"encoding/json"
	"fmt"
	//"time"
)

func main() {

	//startTime := time.Now()

	var err error
	var filePath string

	ts := tables{}
	k := keys{"CRM_SYNC_ID"}

	filePath = `C:\gowstemp\bin\file1.txt`
	t1, err := load(filePath, '\t', k)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ts = append(ts, t1)
	//fmt.Println(filePath, time.Since(startTime), len(t1.rows), "rec")

	filePath = `C:\gowstemp\bin\file2.txt`
	t2, err := load(filePath, '\t', k)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ts = append(ts, t2)
	//fmt.Println(filePath, time.Since(startTime), len(t2.rows), "rec")

	filePath = `C:\gowstemp\bin\file3.txt`
	t3, err := load(filePath, '\t', k)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ts = append(ts, t3)

	t1.print()
	t2.print()
	t3.print()

	ts.delta().print()

	/*	b, err := json.Marshal(ds)
		if err != nil {
			fmt.Println(err)
			return
		}*/

	/*
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
		}*/

}
