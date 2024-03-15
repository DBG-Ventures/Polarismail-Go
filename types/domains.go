package types

type Domains interface {
	List() ([]DomainsList, error)
	CheckAvailable(newDomain string) bool
	Add(newDomain string) (Domain, error)
	Get(domain string) Domain
	// Delete()
}

type Domain interface {
	Info() (DomainInfo, error)
}

type DomainInfo struct {
	Timezone            string `json:"timezone"`
	ExpiryNotifications int    `json:"expiry_notifications"`
	CatchAll            string `json:"catch_all"`
	DomainType          string `json:"domain_type"`
	VUsers              int    `json:"v_users"`
	Aliases             int    `json:"aliases"`
	Lists               int    `json:"lists"`
	Active              bool   `json:"active"`
	ExchangeEnabled     bool   `json:"exchange_enabled"`
	AuditEnabled        bool   `json:"audit_enabled"`
	LocalDelivery       bool   `json:"local_delivery"`
	TotalQuota          int    `json:"total_Quota"`
	HomeCountry         string `json:"home_country"`
}

type DomainsList struct {
	ID            string `json:"id,omitempty"`
	Domain        string `json:"domain,omitempty"`
	IsAuditDomain bool   `json:"is_audit_domain,omitempty"`
	DomainType    string `json:"domain_type,omitempty"`
	Active        bool   `json:"active,omitempty"`
}
