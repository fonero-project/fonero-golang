package resourceadapter

import (
	"context"

	"github.com/fonero-project/fonero-golang/services/horizon/internal/db2/history"
	. "github.com/fonero-project/fonero-golang/protocols/horizon"
)

func PopulateHistoryAccount(ctx context.Context, dest *HistoryAccount, row history.Account) {
	dest.ID = row.Address
	dest.AccountID = row.Address
}
