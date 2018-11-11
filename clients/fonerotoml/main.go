package fonerotoml

import "net/http"

// FoneroTomlMaxSize is the maximum size of fonero.toml file
const FoneroTomlMaxSize = 5 * 1024

// WellKnownPath represents the url path at which the fonero.toml file should
// exist to conform to the federation protocol.
const WellKnownPath = "/.well-known/fonero.toml"

// DefaultClient is a default client using the default parameters
var DefaultClient = &Client{HTTP: http.DefaultClient}

// Client represents a client that is capable of resolving a Fonero.toml file
// using the internet.
type Client struct {
	// HTTP is the http client used when resolving a Fonero.toml file
	HTTP HTTP

	// UseHTTP forces the client to resolve against servers using plain HTTP.
	// Useful for debugging.
	UseHTTP bool
}

type ClientInterface interface {
	GetFoneroToml(domain string) (*Response, error)
	GetFoneroTomlByAddress(addy string) (*Response, error)
}

// HTTP represents the http client that a stellertoml resolver uses to make http
// requests.
type HTTP interface {
	Get(url string) (*http.Response, error)
}

// Response represents the results of successfully resolving a fonero.toml file
type Response struct {
	AuthServer       string `toml:"AUTH_SERVER"`
	FederationServer string `toml:"FEDERATION_SERVER"`
	EncryptionKey    string `toml:"ENCRYPTION_KEY"`
	SigningKey       string `toml:"SIGNING_KEY"`
}

// GetFoneroToml returns fonero.toml file for a given domain
func GetFoneroToml(domain string) (*Response, error) {
	return DefaultClient.GetFoneroToml(domain)
}

// GetFoneroTomlByAddress returns fonero.toml file of a domain fetched from a
// given address
func GetFoneroTomlByAddress(addy string) (*Response, error) {
	return DefaultClient.GetFoneroTomlByAddress(addy)
}
