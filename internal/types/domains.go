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

	active := false
	if dl.Active == "1" {
		active = true
	}

	return types.DomainsList{
		ID:            dl.ID,
		Domain:        dl.Domain,
		IsAuditDomain: dl.IsAuditDomain,
		DomainType:    domainType,
		Active:        active,
	}
}
