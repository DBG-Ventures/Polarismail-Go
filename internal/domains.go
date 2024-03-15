package internal

import (
	"fmt"
	"net/url"

	itypes "github.com/dbgventures/polarismail-go/internal/types"
	"github.com/dbgventures/polarismail-go/types"
)

type Domains struct {
	c *Client
}

func (d Domains) List() ([]types.DomainsList, error) {
	formData := url.Values{
		"action": {"getAllDomains"},
	}

	var respValue itypes.DomainListResponse
	err := d.c.requestWrapper(formData, &respValue)
	if err != nil {
		return []types.DomainsList{}, fmt.Errorf("unable to list all domains: %+v", err)
	}

	var list []types.DomainsList
	for _, domain := range respValue.ReturnData {
		list = append(list, domain.AsType())
	}

	return list, nil
}
