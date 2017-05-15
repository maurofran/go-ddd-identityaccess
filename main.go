package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/facebookgo/inject"
	"github.com/gorilla/mux"
	"github.com/maurofran/go-ddd-identityaccess/application"
	"github.com/maurofran/go-ddd-identityaccess/domain/model"
	"github.com/maurofran/go-ddd-identityaccess/infrastructure/persistence"
	"github.com/maurofran/go-ddd-identityaccess/resource"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2"
	"net/http"
	"os"
	"time"
)

const (
	envDbUrl  = "MONGODB_URL"
	envDbName = "MONGODB_NAME"
)

const (
	defaultDbUrl  = "mongodb://localhost:27017"
	defaultDbName = "identityaccess"
)

var log = logrus.New()

func init() {
	log.Formatter = &logrus.JSONFormatter{}
	log.Out = os.Stdout
	log.Level = logrus.DebugLevel
}

type noOpPublisher struct {}

func (p *noOpPublisher) Publish(events model.DomainEvents) {}

func main() {
	g := inject.Graph{Logger: log}

	dbUrl := os.Getenv(envDbUrl)
	if dbUrl == "" {
		dbUrl = defaultDbUrl
	}

	log.WithField("dbUrl", dbUrl).Info("Connecting to mongo DB")

	session, err := mgo.Dial(dbUrl)
	if err != nil {
		log.WithField("err", err).Fatal("An error occurred while connecting to datanase")
		os.Exit(1)
	}
	defer session.Close()

	dbName := os.Getenv(envDbName)
	if dbName == "" {
		dbName = defaultDbName
	}

	log.WithField("dbName", dbName).Info("Retrieving database")

	db := session.DB(dbName)

	router := mux.NewRouter()

	log.Info("Setting up infrastructure layer.")

	g.Provide(
		&inject.Object{Value: db},
		&inject.Object{Value: validator.New()},
		&inject.Object{Value: router},
		&inject.Object{Value: &noOpPublisher{}},
	)

	log.Info("Setting up persistence layer.")

	g.Provide(&inject.Object{Value: new(persistence.TenantStore)})

	log.Info("Setting up domain layer repositories.")

	g.Provide(&inject.Object{Value: new(model.TenantRepository)})

	log.Info("Setting up domain layer services.")

	g.Provide(&inject.Object{Value: new(model.TenantProvisioningService)})

	log.Info("Setting up application layer services.")

	g.Provide(&inject.Object{Value: new(application.IdentityService)})

	log.Info("Setting up resource layers.")

	apiV1 := new(resource.ApiV1Resources)

	g.Provide(
		&inject.Object{Value: new(resource.TenantResource)},
		&inject.Object{Value: apiV1},
	)

	if err := g.Populate(); err != nil {
		log.WithField("err", err).Fatal("An error occurred while populating object graph")
		os.Exit(1)
	}

	apiV1.Init()

	srv := &http.Server{
		Handler:      router,
		Addr:         ":3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.WithField("err", err).Fatal("An error occurred while starting server")
		os.Exit(1)
	}
}
