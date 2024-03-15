package internal

import (
	"fmt"
	"net/url"

	itypes "github.com/dbgventures/polarismail-go/internal/types"
	"github.com/dbgventures/polarismail-go/types"
)

type Admin struct {
	c *Client
}

func (a Admin) GetStats() (types.AdminStats, error) {
	formData := url.Values{
		"action": {"getAdminStats"},
	}

	var respValue itypes.AdminStatsResponse
	err := a.c.requestWrapper(formData, &respValue)
	if err != nil {
		return types.AdminStats{}, fmt.Errorf("unable to get admin stats: %+v", err)
	}

	return respValue.ReturnData.AsType(), nil
}

func (a Admin) GetBrandInfo() (types.AdminBrandInfo, error) {
	formData := url.Values{
		"action": {"getAdminBrandInfo"},
	}

	var respValue itypes.AdminBrandInfoResponse
	err := a.c.requestWrapper(formData, &respValue)
	if err != nil {
		return types.AdminBrandInfo{}, fmt.Errorf("unable to get admin brand info: %+v", err)
	}

	return respValue.ReturnData.AsType(), nil
}

func (a Admin) GetActionHistory() {

}

func (a Admin) UpdatePassword() {

}

func (a Admin) UpdateEmail() {

}
