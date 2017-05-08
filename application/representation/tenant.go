package representation

type Tenant struct {
	TenantID    string `json:"tenantId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}
