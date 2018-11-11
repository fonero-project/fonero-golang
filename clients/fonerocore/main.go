// Package fonerocore is a client library for communicating with an
// instance of fonero-core using through the server's HTTP port.
package fonerocore

import "net/http"

// SetCursorDone is the success message returned by fonero-core when a cursor
// update succeeds.
const SetCursorDone = "Done"

// HTTP represents the http client that a fonerocore client uses to make http
// requests.
type HTTP interface {
	Do(req *http.Request) (*http.Response, error)
}

// confirm interface conformity
var _ HTTP = http.DefaultClient
