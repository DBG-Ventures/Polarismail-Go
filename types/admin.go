package types

type Admin interface {
	GetStats() (AdminStats, error)
	GetBrandInfo()
	GetActionHistory()
	UpdatePassword()
	UpdateEmail()
}

type AdminStats struct {
	CurrentBasicEmails    int    `json:"current_basic_emails,omitempty"`
	CurrentEnhancedEmails string `json:"current_enhanced_emails,omitempty"`
	MaxBasicEmails        string `json:"max_basic_emails,omitempty"`
	MaxEnhancedEmails     string `json:"max_enhanced_emails,omitempty"`
	CurrentDomains        string `json:"current_domains,omitempty"`
	MaxDomains            string `json:"max_domains,omitempty"`
	CurrentQuota          string `json:"current_quota,omitempty"`
	MaxQuota              string `json:"max_quota,omitempty"`
	AutogrowQuota         string `json:"autogrow_quota,omitempty"`
}
