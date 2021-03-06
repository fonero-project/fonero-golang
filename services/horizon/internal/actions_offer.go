package horizon

import (
	"github.com/fonero-project/fonero-golang/protocols/horizon"
	"github.com/fonero-project/fonero-golang/services/horizon/internal/db2"
	"github.com/fonero-project/fonero-golang/services/horizon/internal/db2/core"
	"github.com/fonero-project/fonero-golang/services/horizon/internal/db2/history"
	"github.com/fonero-project/fonero-golang/services/horizon/internal/render/sse"
	"github.com/fonero-project/fonero-golang/services/horizon/internal/resourceadapter"
	"github.com/fonero-project/fonero-golang/support/render/hal"
)

// This file contains the actions:

// OffersByAccountAction renders a page of offer resources, for a given
// account.  These offers are present in the ledger as of the latest validated
// ledger.
type OffersByAccountAction struct {
	Action
	Address   string
	PageQuery db2.PageQuery
	Records   []core.Offer
	Ledgers   history.LedgerCache
	Page      hal.Page
}

// JSON is a method for actions.JSON
func (action *OffersByAccountAction) JSON() {
	action.Do(
		action.loadParams,
		action.loadRecords,
		action.loadLedgers,
		action.loadPage,
		func() {
			hal.Render(action.W, action.Page)
		},
	)
}

// SSE is a method for actions.SSE
func (action *OffersByAccountAction) SSE(stream sse.Stream) {
	action.Do(
		action.loadParams,
		action.loadRecords,
		action.loadLedgers,
		func() {
			stream.SetLimit(int(action.PageQuery.Limit))
			for _, record := range action.Records[stream.SentCount():] {
				ledger, found := action.Ledgers.Records[record.Lastmodified]
				ledgerPtr := &ledger
				if !found {
					ledgerPtr = nil
				}
				var res horizon.Offer
				resourceadapter.PopulateOffer(action.R.Context(), &res, record, ledgerPtr)
				stream.Send(sse.Event{ID: res.PagingToken(), Data: res})
			}
		},
	)
}

func (action *OffersByAccountAction) loadParams() {
	action.PageQuery = action.GetPageQuery()
	action.Address = action.GetString("account_id")
}

// loadLedgers populates the ledger cache for this action
func (action *OffersByAccountAction) loadLedgers() {
	for _, offer := range action.Records {
		action.Ledgers.Queue(offer.Lastmodified)
	}
	action.Err = action.Ledgers.Load(action.HistoryQ())
}

func (action *OffersByAccountAction) loadRecords() {
	action.Err = action.CoreQ().OffersByAddress(
		&action.Records,
		action.Address,
		action.PageQuery,
	)
}

func (action *OffersByAccountAction) loadPage() {
	for _, record := range action.Records {
		ledger, found := action.Ledgers.Records[record.Lastmodified]
		ledgerPtr := &ledger
		if !found {
			ledgerPtr = nil
		}

		var res horizon.Offer
		resourceadapter.PopulateOffer(action.R.Context(), &res, record, ledgerPtr)
		action.Page.Add(res)
	}

	action.Page.FullURL = action.FullURL()
	action.Page.Limit = action.PageQuery.Limit
	action.Page.Cursor = action.PageQuery.Cursor
	action.Page.Order = action.PageQuery.Order
	action.Page.PopulateLinks()
}
