package dto

type FrontendConfig struct {
	BaseSiteURL   string `json:"base_site_url"`
	BaseAPIURL    string `json:"base_api_url"`
	BaseH5URL     string `json:"base_h5_url"`
	BaseUploadURL string `json:"base_upload_url"`
}
