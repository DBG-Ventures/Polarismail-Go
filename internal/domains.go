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

	var respValue itypes.StatusResponse
	err := d.c.requestWrapper(formData, &respValue)

	return err == nil
}

func (d Domains) Add(newDomain string) (types.Domain, error) {
	formData := url.Values{
		"action": {"addDomain"},
		"domain": {newDomain},
	}

	var respValue itypes.StatusResponse
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
		return types.DomainInfo{}, fmt.Errorf("unable to get domain info: %+v", err)
	}

	return respValue.ReturnData.AsType(), nil
}

func (d Domain) Edit() types.DomainEdit {
	return DomainEdit(d)
}

func (d Domain) DKIM() types.DomainDKIM {
	return DomainDKIM(d)
}

func (d Domain) Footer() types.DomainFooter {
	return DomainFooter(d)
}

func (d Domain) DomainAliases() types.DomainAliases {
	return DomainAliases(d)
}

func (d Domain) Delete() error {
	formData := url.Values{
		"action": {"removeDomain"},
		"domain": {d.name},
	}

	var respValue itypes.StatusResponse
	err := d.c.requestWrapper(formData, &respValue)
	if err != nil {
		return fmt.Errorf("unable to set domain footer %+v", err)
	}

	return nil
}

type DomainEdit struct {
	name string
	c    *Client
}

func (d DomainEdit) DomainEnabled(enabled bool) error {
	activeStatus := "0"
	if enabled {
		activeStatus = "1"
	}

	formData := url.Values{
		"action":           {"updateDomain"},
		"domain":           {d.name},
		"editdomainactive": {activeStatus},
	}

	var respValue itypes.StatusResponse
	err := d.c.requestWrapper(formData, &respValue)
	if err != nil {
		return err
	}

	return nil
}

func (d DomainEdit) CatchAll(status string) error {
	formData := url.Values{
		"action":       {"updateDomain"},
		"domain":       {d.name},
		"editcatchall": {status},
	}

	var respValue itypes.StatusResponse
	err := d.c.requestWrapper(formData, &respValue)
	if err != nil {
		return err
	}

	return nil
}

func (d DomainEdit) Timezone(timezone string) error {
	formData := url.Values{
		"action":     {"updateDomain"},
		"domain":     {d.name},
		"editdtzone": {timezone},
	}

	var respValue itypes.StatusResponse
	err := d.c.requestWrapper(formData, &respValue)
	if err != nil {
		return err
	}

	return nil
}

func (d DomainEdit) ExchangeEnabled(enabled bool) error {
	activeStatus := "0"
	if enabled {
		activeStatus = "1"
	}

	formData := url.Values{
		"action":              {"updateDomain"},
		"domain":              {d.name},
		"editexchangeenabled": {activeStatus},
	}

	var respValue itypes.StatusResponse
	err := d.c.requestWrapper(formData, &respValue)
	if err != nil {
		return err
	}

	return nil
}

type DomainDKIM struct {
	name string
	c    *Client
}

func (d DomainDKIM) GetStatus() (types.DomainDKIMStatus, error) {
	formData := url.Values{
		"action": {"getDKIM"},
		"domain": {d.name},
	}

	var respValue itypes.DomainDKIMStatusResponse
	err := d.c.requestWrapper(formData, &respValue)
	if err != nil {
		return types.DomainDKIMStatus{}, fmt.Errorf("unable to get dkim status: %+v", err)
	}

	return respValue.Returndata.AsType(), nil
}

func (d DomainDKIM) Enable() (types.DomainDKIMStatus, error) {
	formData := url.Values{
		"action": {"enableDKIM"},
		"domain": {d.name},
	}

	var respValue itypes.DomainDKIMEnableResponse
	err := d.c.requestWrapper(formData, &respValue)
	if err != nil {
		return types.DomainDKIMStatus{}, fmt.Errorf("unable to enable dkim: %+v", err)
	}

	return respValue.AsType(), nil
}

func (d DomainDKIM) Disable() error {
	formData := url.Values{
		"action": {"disableDKIM"},
		"domain": {d.name},
	}

	var respValue itypes.StatusResponse
	err := d.c.requestWrapper(formData, &respValue)
	if err != nil {
		return fmt.Errorf("unable to disable dkim: %+v", err)
	}

	return nil
}

type DomainFooter struct {
	name string
	c    *Client
}

func (d DomainFooter) Get() (string, error) {
	formData := url.Values{
		"action": {"getFooterDomain"},
		"domain": {d.name},
	}

	var respValue itypes.StatusResponse
	err := d.c.requestWrapper(formData, &respValue)
	if err != nil {
		return "", fmt.Errorf("unable to get domain footer %+v", err)
	}

	return respValue.ReturnData, nil
}

func (d DomainFooter) Set(footer string) error {
	formData := url.Values{
		"action":    {"setFooterDomain"},
		"domain":    {d.name},
		"footertxt": {footer},
	}

	var respValue itypes.StatusResponse
	err := d.c.requestWrapper(formData, &respValue)
	if err != nil {
		return fmt.Errorf("unable to set domain footer %+v", err)
	}

	return nil
}

type DomainAliases struct {
	name string
	c    *Client
}

func (d DomainAliases) List() ([]types.DomainAlias, error) {
	formData := url.Values{
		"action": {"getAliasDomains"},
		"domain": {d.name},
	}

	var respValue itypes.DomainAliasesResponse
	err := d.c.requestWrapper(formData, &respValue)
	if err != nil {
		return []types.DomainAlias{}, fmt.Errorf("unable to list all domain aliases: %+v", err)
	}

	var list []types.DomainAlias
	for _, alias := range respValue.Returndata {
		list = append(list, alias.AsType())
	}

	return list, nil
}

func (d DomainAliases) Add(alias string) error {
	formData := url.Values{
		"action":  {"addAliasDomain"},
		"domain":  {d.name},
		"adomain": {alias},
	}

	var respValue itypes.StatusResponse
	err := d.c.requestWrapper(formData, &respValue)
	if err != nil {
		return fmt.Errorf("unable to set domain footer %+v", err)
	}

	return nil
}

func (d DomainAliases) Delete(alias string) error {
	formData := url.Values{
		"action":  {"removeAliasDomain"},
		"domain":  {d.name},
		"adomain": {alias},
	}

	var respValue itypes.StatusResponse
	err := d.c.requestWrapper(formData, &respValue)
	if err != nil {
		return fmt.Errorf("unable to set domain footer %+v", err)
	}

	return nil
}
