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
	formData := url.Values{
		"action": {"getAdminStats"},
		"token":  {a.c.apiKey},
	}

	payload := strings.NewReader(formData.Encode())

	req, err := a.c.newRequest(payload)
	if err != nil {
		return types.AdminStats{}, err
	}

	var respValue itypes.AdminStatsResponse
	resp, err := a.c.do(req, &respValue)
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
