package main

import (
	"fmt"
)

func (ts tables) delta() (ds filesDelta) {

	for i, _ := range ts {
		if i == 0 {
			continue
		}

		t1 := ts[i-1]
		t2 := ts[i]

		fdelta := fileDelta{}
		fdelta.oldFileName = t1.sourceFileName
		fdelta.newFileName = t2.sourceFileName
		fdelta.rowDelta = rowDelta{}

		for key, _ := range t1.rows { // for each key in t1, compare with t2
			//fmt.Println("key:", key)

			if _, ok := t2.rows[key]; !ok { //// see if the key exists in t2
				//TODO: report key missing in t2
				continue
			}

			//// if key exists,
			cvDelta := columnValueDelta{}
			coDelta := columnOtherDelta{}
			for col, _ := range t1.rows[key] { // for each column
				//fmt.Println("col:", col, t1.rows[key][col], t2.rows[key][col])

				if _, ok := t2.rows[key][col]; !ok { // see if the column exists in t2
					coDelta[col] = "column missing"
					//TODO: report column missing in t2
					continue
				}
				// if column exists, then compare
				if t1.rows[key][col] != t2.rows[key][col] {
					cvDelta[col] = valueDelta{
						oldValue: t1.rows[key][col],
						newValue: t2.rows[key][col],
					}
				}
			}

			if len(cvDelta) > 0 {
				columnDelta := fdelta.rowDelta[key]
				columnDelta.columnValueDelta = cvDelta
				fdelta.rowDelta[key] = columnDelta
			}
			if len(coDelta) > 0 {
				columnDelta := fdelta.rowDelta[key]
				columnDelta.columnValueDelta = coDelta
				fdelta.rowDelta[key] = columnDelta
			}

		}
		ds = append(ds, fdelta)

		for key, _ := range t2.rows { // for each key in t2, compare with t1

			if _, ok := t1.rows[key]; !ok { //// see if the key exists in t1
				//TODO: report key missing in t1
				continue
			}

			//// if key exists,
			coDelta := columnOtherDelta{}
			for col, _ := range t2.rows[key] { // for each column
				if _, ok := t1.rows[key][col]; !ok { // see if the column exists in t1
					//TODO: report column missing in t1
					coDelta[col] = "new column"
					continue
				}
			}
			if len(coDelta) > 0 {
				columnDelta := fdelta.rowDelta[key]
				columnDelta.columnValueDelta = coDelta
				fdelta.rowDelta[key] = columnDelta
			}

		}
	}

	return
}

func (ds filesDelta) print() {
	if len(ds) == 0 {
		fmt.Println("no files to compare")
		return
	}
	tab := ""
	for i, fd := range ds { //for each pair of files
		tab = "-"
		fmt.Println()
		fmt.Println(tab, i, "comparing file", fd.oldFileName, "with", fd.newFileName, ":", len(fd.rowDelta), "diff")

		if len(fd.rowDelta) == 0 {
			tab := "--"
			fmt.Println(tab, "no differences")
			continue
		}

		for key, cd := range fd.rowDelta {
			if len(cd.columnValueDelta) == 0 && len(cd.columnValueDelta) == 0 {
				continue
			}
			tab := "--"
			fmt.Println(tab, "differences for key", key)

			for col, d := range cd.columnValueDelta {
				tab := "---"
				fmt.Println(tab, "column", col, "changed from", d.oldValue, "to", d.newValue)
			}
		}
	}
}
