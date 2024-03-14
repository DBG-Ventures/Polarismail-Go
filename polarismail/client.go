package polarismail

import (
	"net/http"

	"github.com/dbgventures/polarismail-go/internal"
	"github.com/dbgventures/polarismail-go/types"
)

func NewClient(creds *types.Credentials) (types.Client, error) {
	return internal.NewClient(creds, nil)
}

func NewClientWithCustomHttpClient(creds *types.Credentials, httpClient *http.Client) (types.Client, error) {
	return internal.NewClient(creds, httpClient)
}
