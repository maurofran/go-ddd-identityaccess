package main

import (
	"fmt"
	"github.com/facebookgo/inject"
	"github.com/maurofran/go-ddd-identityaccess/application"
	"github.com/maurofran/go-ddd-identityaccess/domain/model"
	"github.com/maurofran/go-ddd-identityaccess/infrastructure/persistence"
	"github.com/maurofran/go-ddd-identityaccess/resource"
	"os"
)

func main() {
	var g inject.Graph

	g.Provide(
		// Setup domain
		&inject.Object{Value: model.TenantRepository{}},
		&inject.Object{Value: model.TenantProvisioningService{}},
		// Setup persistence
		&inject.Object{Value: persistence.TenantStore{}},
		// Setup application
		&inject.Object{Value: application.IdentityApplicationService{}},
		// Setup resources
		&inject.Object{Value: resource.TenantResource{}},
	)

	if err := g.Populate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
