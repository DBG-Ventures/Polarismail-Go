package types

type ErrorCheck struct {
	ReturnCode int `json:"returncode"`
}

type ErrorResponse struct {
	ReturnCode int    `json:"returncode"`
	ReturnData string `json:"returndata"`
}
