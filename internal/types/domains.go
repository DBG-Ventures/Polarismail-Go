package types

import "github.com/dbgventures/polarismail-go/types"

type DomainListResponse struct {
	ReturnCode int              `json:"returncode"`
	ReturnData []DomainListData `json:"returndata"`
}

type DomainListData struct {
	ID            string `json:"id"`
	Domain        string `json:"domain"`
	IsAuditDomain string `json:"isauditdomain"`
	DomainType    string `json:"domaintype"`
	Active        string `json:"active"`
}

func (dl *DomainListData) AsType() types.DomainsList {
	domainType := "basic"
	if dl.DomainType == "2" {
		domainType = "enhanced"
	}

	return types.DomainsList{
		ID:            dl.ID,
		Domain:        dl.Domain,
		IsAuditDomain: dl.IsAuditDomain == "1",
		DomainType:    domainType,
		Active:        dl.Active == "1",
	}
}

type DomainInfoResponse struct {
	ReturnData DomainInfoData `json:"returndata"`
	ReturnCode int            `json:"returncode"`
}

type DomainInfoData struct {
	Dtzone              string `json:"dtzone"`
	ExpiryNotifications int    `json:"expiry_notifications"`
	Catchall            string `json:"catchall"`
	DomainType          string `json:"domaintype"`
	Webmail             string `json:"webmail"`
	Vusers              int    `json:"vusers"`
	Aliases             int    `json:"aliases"`
	Lists               int    `json:"lists"`
	Active              string `json:"active"`
	ExchangeEnabled     string `json:"exchange_enabled"`
	AuditEnabled        string `json:"audit_enabled"`
	Localdelivery       string `json:"localdelivery"`
	TotalQuota          int    `json:"totalQuota"`
	Homecountry         string `json:"homecountry"`
}

func (di *DomainInfoData) AsType() types.DomainInfo {
	domainType := "basic"
	if di.DomainType == "2" {
		domainType = "enhanced"
	}

	return types.DomainInfo{
		Timezone:            di.Dtzone,
		ExpiryNotifications: di.ExpiryNotifications,
		CatchAll:            di.Catchall,
		DomainType:          domainType,
		VUsers:              di.Vusers,
		Aliases:             di.Aliases,
		Lists:               di.Lists,
		Enabled:             di.Active == "1",
		ExchangeEnabled:     di.ExchangeEnabled == "1",
		AuditEnabled:        di.AuditEnabled == "1",
		LocalDelivery:       di.Localdelivery == "1",
		TotalQuota:          0,
		HomeCountry:         di.Homecountry,
	}
}

type DomainDKIMStatusResponse struct {
	Returncode int                  `json:"returncode"`
	Returndata DomainDKIMStatusData `json:"returndata"`
}

type DomainDKIMStatusData struct {
	DkimEnabled string `json:"dkim_enabled"`
	DkimHost    string `json:"dkim_host"`
	DkimKey     string `json:"dkim_key"`
}

func (dd DomainDKIMStatusData) AsType() types.DomainDKIMStatus {
	return types.DomainDKIMStatus{
		DkimEnabled: dd.DkimEnabled == "1",
		DkimHost:    dd.DkimHost,
		DkimKey:     dd.DkimKey,
	}
}

type DomainDKIMEnableResponse struct {
	Returncode int    `json:"returncode"`
	Returndata string `json:"returndata"`
	Dkimhost   string `json:"dkimhost"`
	Dkimkey    string `json:"dkimkey"`
}

func (dd DomainDKIMEnableResponse) AsType() types.DomainDKIMStatus {
	return types.DomainDKIMStatus{
		DkimEnabled: true,
		DkimHost:    dd.Dkimhost,
		DkimKey:     dd.Dkimkey,
	}
}

type DomainAliasesResponse struct {
	Returncode int                 `json:"returncode"`
	Returndata []DomainAliasesData `json:"returndata"`
}

type DomainAliasesData struct {
	AliasDomainID int    `json:"alias_domain_id"`
	AliasDomain   string `json:"alias_domain"`
}

func (da DomainAliasesData) AsType() types.DomainAlias {
	return types.DomainAlias(da)
}
