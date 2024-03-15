package types

type Domains interface {
	List() ([]DomainsList, error)
	// CheckAvailable()
	// Add()
	// Get()
	// Delete()
}

type DomainsList struct {
	ID            string `json:"id,omitempty"`
	Domain        string `json:"domain,omitempty"`
	IsAuditDomain string `json:"is_audit_domain,omitempty"`
	DomainType    string `json:"domain_type,omitempty"`
	Active        bool   `json:"active,omitempty"`
}
