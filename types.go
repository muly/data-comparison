package main

type (
	valueDelta struct { //d
		oldValue string `json:"OldValue"`
		newValue string `json:"NewValue"`
	}
	columnValueDelta map[string]valueDelta // map of column to its delta //cd
	columnOtherDelta map[string]string     //map of column to other delta. example: column missing, new column

	columnDelta struct {
		columnValueDelta
		columnOtherDelta
	}

	rowDelta map[string]columnDelta // map of row key to its delta //rd

	fileDelta struct { //table delta //fd
		oldFileName string `json:"OldFileName"`
		newFileName string `json:"NewFileName"`
		rowDelta    `json:"RowDelta"`
	}

	filesDelta []fileDelta //ds
)

type (
	keys  []string
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
