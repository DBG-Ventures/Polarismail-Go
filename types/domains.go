package types

type Domains interface {
	List() ([]DomainsList, error)
	CheckAvailable(newDomain string) bool
	Add(newDomain string) (Domain, error)
	Get(domain string) Domain
}

type Domain interface {
	Info() (DomainInfo, error)
	Edit() DomainEdit
	DKIM() DomainDKIM
}

type DomainEdit interface {
	DomainEnabled(enabled bool) error
	CatchAll(status string) error
	Timezone(timezone string) error
	ExchangeEnabled(enabled bool) error
}

type DomainDKIM interface {
	GetStatus() (DomainDKIMStatus, error)
	Enable() (DomainDKIMStatus, error)
	Disable() error
}

type DomainInfo struct {
	Timezone            string `json:"timezone"`
	ExpiryNotifications int    `json:"expiry_notifications"`
	CatchAll            string `json:"catch_all"`
	DomainType          string `json:"domain_type"`
	VUsers              int    `json:"v_users"`
	Aliases             int    `json:"aliases"`
	Lists               int    `json:"lists"`
	Enabled             bool   `json:"enabled"`
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

type DomainDKIMStatus struct {
	DkimEnabled bool   `json:"dkim_enabled"`
	DkimHost    string `json:"dkim_host"`
	DkimKey     string `json:"dkim_key"`
}
