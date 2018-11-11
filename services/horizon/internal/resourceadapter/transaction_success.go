package resourceadapter

import (
	"context"

	"github.com/fonero-project/fonero-golang/services/horizon/internal/httpx"
	"github.com/fonero-project/fonero-golang/services/horizon/internal/txsub"
	. "github.com/fonero-project/fonero-golang/protocols/horizon"
	"github.com/fonero-project/fonero-golang/support/render/hal"
)

// Populate fills out the details
func PopulateTransactionSuccess(ctx context.Context, dest *TransactionSuccess, result txsub.Result) {
	dest.Hash = result.Hash
	dest.Ledger = result.LedgerSequence
	dest.Env = result.EnvelopeXDR
	dest.Result = result.ResultXDR
	dest.Meta = result.ResultMetaXDR

	lb := hal.LinkBuilder{httpx.BaseURL(ctx)}
	dest.Links.Transaction = lb.Link("/transactions", result.Hash)
	return
}
