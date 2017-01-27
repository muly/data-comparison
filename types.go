package main

type (
	delta struct { //d
		oldValue string `json:"OldValue"`
		newValue string `json:"NewValue"`
	}
	columnDelta map[string]delta // map of column to its delta //cd

	rowDelta map[string]columnDelta // map of row key to its delta //rd

	fileDelta struct { //table delta //fd
		oldFileName string `json:"OldFileName"`
		newFileName string `json:"NewFileName"`
		rowDelta    `json:"RowDelta"`
	}

	filesDelta []fileDelta //ds
)

type (
	row   map[string]string // map of column to its value
	rows  map[string]row    // map of key to row
	table struct {
		sourceFileName string
		rows
	}
	tables []table
)

/*type (
	CertKey struct {
		CertID string
	}

	CertData struct {
		CertType          string
		OriginalGrantDate string
		CertEndDate       string
		CertStatus        string
	}

	CertRecord map[CertKey]CertData

	CertRecords []CertRecord

	CertFile struct {
		ProcessedFile string
		CertRecords
	}
)*/
