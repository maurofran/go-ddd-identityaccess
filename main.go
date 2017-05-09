package main

import (
	"fmt"
	"github.com/facebookgo/inject"
	"github.com/maurofran/go-ddd-identityaccess/application"
	"github.com/maurofran/go-ddd-identityaccess/domain/model"
	"github.com/maurofran/go-ddd-identityaccess/infrastructure/persistence"
	"github.com/maurofran/go-ddd-identityaccess/resource"
	"gopkg.in/mgo.v2"
	"os"
)

const (
	envDbUrl  = "MONGODB_URL"
	envDbName = "MONGODB_NAME"
)

const (
	defaultDbUrl  = "mongodb://localhost:27017"
	defaultDbName = "identityaccess"
)

func main() {
	var g inject.Graph
	var session *mgo.Session

	dbUrl := os.Getenv(envDbUrl)
	if dbUrl == "" {
		dbUrl = defaultDbUrl
	}
	session, err := mgo.Dial(dbUrl)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer session.Close()

	dbName := os.Getenv(envDbName)
	if dbName == "" {
		dbName = defaultDbName
	}

	g.Provide(
		// Setup database connection
		&inject.Object{Value: session.DB(dbName)},
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
