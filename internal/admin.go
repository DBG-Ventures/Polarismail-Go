package internal

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	itypes "github.com/dbgventures/polarismail-go/internal/types"
	"github.com/dbgventures/polarismail-go/types"
)

type Admin struct {
	c *Client
}

func (a Admin) GetStats() (types.AdminStats, error) {
	apiKey, err := a.c.authenticate()
	if err != nil {
		return types.AdminStats{}, err
	}

	formData := url.Values{
		"action": {"getAdminStats"},
		"token":  {apiKey},
	}

	payload := strings.NewReader(formData.Encode())

	var respValue itypes.AdminStatsResponse
	resp, err := a.c.newRequest(payload, &respValue)
	if err != nil {
		return types.AdminStats{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return types.AdminStats{}, fmt.Errorf("unable to get admin stats: %+v", resp)
	}

	return respValue.ReturnData.AsType(), nil
}

func (a Admin) GetBrandInfo() {

}
func (a Admin) GetActionHistory() {

}

func (a Admin) UpdatePassword() {

}

func (a Admin) UpdateEmail() {

}
