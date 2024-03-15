package types

type Admin interface {
	GetStats() (AdminStats, error)
	GetBrandInfo() (AdminBrandInfo, error)
	GetActionHistory() ([]AdminActionHistory, error)
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

type AdminBrandInfo struct {
	BasicLogo     string `json:"basic_logo"`
	BasicLogoHref string `json:"basic_logo_href"`
	SupportEmail  string `json:"support_email"`
	BrandName     string `json:"brand_name"`
	BrandLink     string `json:"brand_link"`
	BrandColor    string `json:"brand_color"`
}

type AdminActionHistory struct {
	Action   string `json:"action"`
	Username string `json:"username"`
	RemoteIp string `json:"remote_ip"`
	Date     string `json:"date"`
}
