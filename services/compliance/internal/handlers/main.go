package handlers

import (
	"strconv"
	"time"

	"github.com/fonero-project/fonero-golang/clients/federation"
	"github.com/fonero-project/fonero-golang/clients/fonerotoml"
	"github.com/fonero-project/fonero-golang/services/compliance/internal/config"
	"github.com/fonero-project/fonero-golang/services/compliance/internal/crypto"
	"github.com/fonero-project/fonero-golang/services/compliance/internal/db"
	"github.com/fonero-project/fonero-golang/support/http"
)

// RequestHandler implements compliance server request handlers
type RequestHandler struct {
	Config                  *config.Config                 `inject:""`
	Client                  http.SimpleHTTPClientInterface `inject:""`
	Database                db.Database                    `inject:""`
	SignatureSignerVerifier crypto.SignerVerifierInterface `inject:""`
	FoneroTomlResolver     fonerotoml.ClientInterface    `inject:""`
	FederationResolver      federation.ClientInterface     `inject:""`
	NonceGenerator          NonceGeneratorInterface        `inject:""`
}

type NonceGeneratorInterface interface {
	Generate() string
}

type NonceGenerator struct{}

func (n *NonceGenerator) Generate() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

type TestNonceGenerator struct{}

func (n *TestNonceGenerator) Generate() string {
	return "nonce"
}
