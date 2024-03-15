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

func (a Admin) GetActionHistory() ([]types.AdminActionHistory, error) {
	formData := url.Values{
		"action": {"getAccountStats"},
	}

	var respValue itypes.AdminActionHistoryResponse
	err := a.c.requestWrapper(formData, &respValue)
	if err != nil {
		return []types.AdminActionHistory{}, err
	}

	var history []types.AdminActionHistory

	for _, action := range respValue.ReturnData {
		history = append(history, action.AsType())
	}

	return history, nil
}

func (a Admin) UpdatePassword(newPassword string) error {
	formData := url.Values{
		"action":    {"changePassword"},
		"password1": {a.c.creds.Password},
		"password2": {newPassword},
	}

	var respValue itypes.AdminUpdateResponse
	err := a.c.requestWrapper(formData, &respValue)
	if err != nil {
		return err
	}

	a.c.creds.Password = newPassword

	return nil
}

func (a Admin) UpdateEmail(newEmail string) error {
	formData := url.Values{
		"action":     {"changeEmail"},
		"emailadmin": {newEmail},
	}

	var respValue itypes.AdminUpdateResponse
	err := a.c.requestWrapper(formData, &respValue)
	if err != nil {
		return err
	}

	return nil
}
