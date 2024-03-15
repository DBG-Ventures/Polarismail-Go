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

func (d Domains) CheckAvailable(newDomain string) bool {
	formData := url.Values{
		"action": {"isDomainAvailable"},
		"domain": {newDomain},
	}

	var respValue itypes.ErrorResponse
	err := d.c.requestWrapper(formData, &respValue)

	return err == nil
}

func (d Domains) Add(newDomain string) (types.Domain, error) {
	formData := url.Values{
		"action": {"addDomain"},
		"domain": {newDomain},
	}

	var respValue itypes.ErrorResponse
	err := d.c.requestWrapper(formData, &respValue)
	if err != nil {
		return nil, err
	}

	return Domain{
		name: newDomain,
		c:    d.c,
	}, nil
}

func (d Domains) Get(domain string) types.Domain {
	return Domain{
		name: domain,
		c:    d.c,
	}
}

type Domain struct {
	name string
	c    *Client
}

func (d Domain) Info() (types.DomainInfo, error) {
	formData := url.Values{
		"action": {"getDomainInfo"},
		"domain": {d.name},
	}

	var respValue itypes.DomainInfoResponse
	err := d.c.requestWrapper(formData, &respValue)
	if err != nil {
		return types.DomainInfo{}, fmt.Errorf("unable to list all domains: %+v", err)
	}

	return respValue.ReturnData.AsType(), nil
}
