package fonerotoml

import (
	"net/http"
	"strings"
	"testing"

	"github.com/fonero-project/fonero-golang/support/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientURL(t *testing.T) {
	//HACK:  we're testing an internal method rather than setting up a http client
	//mock.

	c := &Client{UseHTTP: false}
	assert.Equal(t, "https://fonero.org/.well-known/fonero.toml", c.url("fonero.org"))

	c = &Client{UseHTTP: true}
	assert.Equal(t, "http://fonero.org/.well-known/fonero.toml", c.url("fonero.org"))
}

func TestClient(t *testing.T) {
	h := httptest.NewClient()
	c := &Client{HTTP: h}

	// happy path
	h.
		On("GET", "https://fonero.org/.well-known/fonero.toml").
		ReturnString(http.StatusOK,
			`FEDERATION_SERVER="https://localhost/federation"`,
		)
	stoml, err := c.GetFoneroToml("fonero.org")
	require.NoError(t, err)
	assert.Equal(t, "https://localhost/federation", stoml.FederationServer)

	// fonero.toml exceeds limit
	h.
		On("GET", "https://toobig.org/.well-known/fonero.toml").
		ReturnString(http.StatusOK,
			`FEDERATION_SERVER="https://localhost/federation`+strings.Repeat("0", FoneroTomlMaxSize)+`"`,
		)
	stoml, err = c.GetFoneroToml("toobig.org")
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "fonero.toml response exceeds")
	}

	// not found
	h.
		On("GET", "https://missing.org/.well-known/fonero.toml").
		ReturnNotFound()
	stoml, err = c.GetFoneroToml("missing.org")
	assert.EqualError(t, err, "http request failed with non-200 status code")

	// invalid toml
	h.
		On("GET", "https://json.org/.well-known/fonero.toml").
		ReturnJSON(http.StatusOK, map[string]string{"hello": "world"})
	stoml, err = c.GetFoneroToml("json.org")

	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "toml decode failed")
	}
}
