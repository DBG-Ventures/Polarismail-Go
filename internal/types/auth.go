package types

type ApiKeyResponse struct {
	ReturnCode int    `json:"returncode,omitempty"`
	ReturnData string `json:"returndata,omitempty"`
}
