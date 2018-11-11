package handlers

import (
	"github.com/fonero-project/fonero-golang/clients/federation"
	"github.com/fonero-project/fonero-golang/clients/horizon"
	"github.com/fonero-project/fonero-golang/clients/fonerotoml"
	"github.com/fonero-project/fonero-golang/services/bridge/internal/config"
	"github.com/fonero-project/fonero-golang/services/bridge/internal/db"
	"github.com/fonero-project/fonero-golang/services/bridge/internal/listener"
	"github.com/fonero-project/fonero-golang/services/bridge/internal/submitter"
	"github.com/fonero-project/fonero-golang/support/http"
)

// RequestHandler implements bridge server request handlers
type RequestHandler struct {
	Config               *config.Config                          `inject:""`
	Client               http.SimpleHTTPClientInterface          `inject:""`
	Horizon              horizon.ClientInterface                 `inject:""`
	Database             db.Database                             `inject:""`
	FoneroTomlResolver  fonerotoml.ClientInterface             `inject:""`
	FederationResolver   federation.ClientInterface              `inject:""`
	TransactionSubmitter submitter.TransactionSubmitterInterface `inject:""`
	PaymentListener      *listener.PaymentListener               `inject:""`
}

func (rh *RequestHandler) isAssetAllowed(code string, issuer string) bool {
	for _, asset := range rh.Config.Assets {
		if asset.Code == code && asset.Issuer == issuer {
			return true
		}
	}
	return false
}
