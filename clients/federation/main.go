package federation

import (
	"net/http"
	"net/url"

	"github.com/fonero-project/fonero-golang/clients/horizon"
	"github.com/fonero-project/fonero-golang/clients/fonerotoml"
	proto "github.com/fonero-project/fonero-golang/protocols/federation"
)

// FederationResponseMaxSize is the maximum size of response from a federation server
const FederationResponseMaxSize = 100 * 1024

// DefaultTestNetClient is a default federation client for testnet
var DefaultTestNetClient = &Client{
	HTTP:        http.DefaultClient,
	Horizon:     horizon.DefaultTestNetClient,
	FoneroTOML: fonerotoml.DefaultClient,
}

// DefaultPublicNetClient is a default federation client for pubnet
var DefaultPublicNetClient = &Client{
	HTTP:        http.DefaultClient,
	Horizon:     horizon.DefaultPublicNetClient,
	FoneroTOML: fonerotoml.DefaultClient,
}

// Client represents a client that is capable of resolving a federation request
// using the internet.
type Client struct {
	FoneroTOML FoneroTOML
	HTTP        HTTP
	Horizon     Horizon
	AllowHTTP   bool
}

type ClientInterface interface {
	LookupByAddress(addy string) (*proto.NameResponse, error)
	LookupByAccountID(aid string) (*proto.IDResponse, error)
	ForwardRequest(domain string, fields url.Values) (*proto.NameResponse, error)
}

// Horizon represents a horizon client that can be consulted for data when
// needed as part of the federation protocol
type Horizon interface {
	HomeDomainForAccount(aid string) (string, error)
}

// HTTP represents the http client that a federation client uses to make http
// requests.
type HTTP interface {
	Get(url string) (*http.Response, error)
}

// FoneroTOML represents a client that can resolve a given domain name to
// fonero.toml file.  The response is used to find the federation server that a
// query should be made against.
type FoneroTOML interface {
	GetFoneroToml(domain string) (*fonerotoml.Response, error)
}

// confirm interface conformity
var _ FoneroTOML = fonerotoml.DefaultClient
var _ HTTP = http.DefaultClient
var _ ClientInterface = &Client{}
