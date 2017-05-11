package resource

import (
	"github.com/Sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"os"
)

var log = logrus.New()

func init() {
	log.Formatter = &logrus.JSONFormatter{}
	log.Out = os.Stdout
	log.Level = logrus.InfoLevel
}

// Error represents an handler error. It provides methods for an HTTP status code and embeds the built-in error
// interface.
type Error interface {
	error
	Status() int
}

// StatusError represents an error with an associated HTTP status code.
type StatusError struct {
	Code int
	Err  error
}

// Error will allow StatusError to satisfy the error interface.
func (se StatusError) Error() string {
	return se.Err.Error()
}

// Status will allow StatusError to satisfy the resource.Error interface.
func (se StatusError) Status() int {
	return se.Code
}

// Handler is a type alias for handler functions.
type Handler func(w http.ResponseWriter, r *http.Request) error

// ServeHTTP will satisfy the http.Handler interface
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h(w, r)
	if err != nil {
		switch e := err.(type) {
		case Error:
			// We can retrieve the status here and write out specific HTTP status code.
			log.WithField("status", e.Status()).WithField("err", e.Error()).
				Error("An error occurred while processing request")
			http.Error(w, e.Error(), e.Status())
		case validator.ValidationErrors:
			// TODO Check if provide information about failed validations
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		default:
			log.WithField("err", e.Error()).Error("Unexpected error occurred")
			// Any error types we don't specifically look out for default to serving a HTTP 500
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
