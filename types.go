package main

type (
	delta struct {
		key       interface{}
		fieldName string `json:"FieldName"`
		oldValue  string `json:"OldValue"`
		newValue  string `json:"NewValue"`
	}

	deltas struct {
		oldFileName string
		newFileName string
		delta       []delta
	}

	deltasAll []deltas
)

type (
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
)
