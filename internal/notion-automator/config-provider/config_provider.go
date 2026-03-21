package configprovider

type ConfigMain struct {
	NotionToken string              `json:"notion_token"`
	SetProperties []ConfigSetProperty `json:"set_properties"`
}

type ConfigSetProperty struct {
	PageId       string `json:"page_id"`
	PropertyType string `json:"property_type"`
	NewValue     string `json:"new_value"`
}

type ConfigProvider interface {
	GetConfig(filepath string) ConfigMain
}
