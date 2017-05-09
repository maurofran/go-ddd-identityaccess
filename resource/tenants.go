package resource

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/maurofran/go-ddd-identityaccess/application"
	"net/http"
)

type TenantResource struct {
	ias *application.IdentityApplicationService
}

// GetTenant will send the tenant with provided tenant id.
func (tr *TenantResource) GetTenant(w http.ResponseWriter, r *http.Request) error {
	tenantId := mux.Vars(r)["tenantId"]
	t, err := tr.ias.Tenant(r.Context(), tenantId)
	if err != nil {
		return err
	}
	if t == nil {
		w.WriteHeader(http.StatusNotFound)
		return nil
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", iasMediaType)
	json.NewEncoder(w).Encode(t)
	return nil
}