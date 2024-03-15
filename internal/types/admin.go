package types

import "github.com/dbgventures/polarismail-go/types"

type AdminStatsResponse struct {
	ReturnData AdminStatsData `json:"returndata"`
	ReturnCode int            `json:"returncode"`
}

type AdminStatsData struct {
	CurrentBasicEmails    int    `json:"curemails"`
	CurrentEnhancedEmails string `json:"cureemails"`
	Curocemails           string `json:"curocemails"` // TODO: Update name to better reflect: not listed in documentation
	Curexemails           string `json:"curexemails"` // TODO: Update name to better reflect: not listed in documentation
	MaxBasicEmails        string `json:"maxemails"`
	MaxEnhancedEmails     string `json:"maxeemails"`
	Maxocemails           string `json:"maxocemails"` // TODO: Update name to better reflect: not listed in documentation
	Maxexemails           string `json:"maxexemails"` // TODO: Update name to better reflect: not listed in documentation
	CurrentDomains        string `json:"curdomains"`
	MaxDomains            string `json:"maxdomains"`
	CurrentQuota          string `json:"curquota"`
	MaxQuota              string `json:"maxquota"`
	AutogrowQuota         string `json:"autogrow_quota"`
}

func (a AdminStatsData) AsType() types.AdminStats {
	return types.AdminStats{
		CurrentBasicEmails:    a.CurrentBasicEmails,
		CurrentEnhancedEmails: a.CurrentEnhancedEmails,
		MaxBasicEmails:        a.MaxBasicEmails,
		MaxEnhancedEmails:     a.MaxEnhancedEmails,
		CurrentDomains:        a.CurrentDomains,
		MaxDomains:            a.MaxDomains,
		CurrentQuota:          a.CurrentQuota,
		MaxQuota:              a.MaxQuota,
		AutogrowQuota:         a.AutogrowQuota,
	}
}

type AdminBrandInfoResponse struct {
	ReturnData AdminBrandInfoData `json:"returndata"`
	ReturnCode int                `json:"returncode"`
}

type AdminBrandInfoData struct {
	BasicLogo     string `json:"basic_logo"`
	BasicLogoHref string `json:"basic_logo_href"`
	SupportEmail  string `json:"supportemail"`
	BrandName     string `json:"brandname"`
	BrandLink     string `json:"brandlink"`
	BrandColor    string `json:"brandcolor"`
}

func (a AdminBrandInfoData) AsType() types.AdminBrandInfo {
	return types.AdminBrandInfo{
		BasicLogo:     a.BasicLogo,
		BasicLogoHref: a.BasicLogoHref,
		SupportEmail:  a.SupportEmail,
		BrandName:     a.BrandName,
		BrandLink:     a.BrandLink,
		BrandColor:    a.BrandColor,
	}
}
