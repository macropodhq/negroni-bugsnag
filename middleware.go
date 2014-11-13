// Package negronibugsnag provides a negroni middleware that reports exceptions
// to bugsnag.
package negronibugsnag

import (
	"net/http"

	"github.com/bugsnag/bugsnag-go"
)

type NegroniBugsnag struct {
	notifier *bugsnag.Notifier
}

// New returns a new instance of NegroniBugsnag, given a *bugsnag.Notifier
func New(notifier *bugsnag.Notifier) *NegroniBugsnag {
	return &NegroniBugsnag{
		notifier: notifier,
	}
}

func (nbs *NegroniBugsnag) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer nbs.notifier.AutoNotify()
	next(w, r)
}
