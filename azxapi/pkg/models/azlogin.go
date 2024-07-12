// pkg/models/azlogin.go
package models

type AzLoginOutput struct {
	CloudName          string `json:"cloudName"`
	HomeTenantId       string `json:"homeTenantId"`
	Id                 string `json:"id"`
	IsDefault          bool   `json:"isDefault"`
	ManagedByTenants   []interface{} `json:"managedByTenants"`
	Name               string `json:"name"`
	State              string `json:"state"`
	TenantDefaultDomain string `json:"tenantDefaultDomain"`
	TenantDisplayName  string `json:"tenantDisplayName"`
	TenantId           string `json:"tenantId"`
	User               struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"user"`
}
