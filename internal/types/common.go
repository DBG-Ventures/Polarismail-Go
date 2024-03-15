package types

type ErrorCheck struct {
	ReturnCode int `json:"returncode"`
}

type StatusResponse struct {
	ReturnCode int    `json:"returncode"`
	ReturnData string `json:"returndata"`
}
