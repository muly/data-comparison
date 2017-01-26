package main

type (
	columnDelta struct {
		fieldName string `json:"FieldName"`
		oldValue  string `json:"OldValue"`
		newValue  string `json:"NewValue"`
	}
	rowDelta map[string]columnDelta

	fileDelta struct {
		oldFileName string
		newFileName string
		rowDelta
	}

	filesDelta []fileDelta
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
