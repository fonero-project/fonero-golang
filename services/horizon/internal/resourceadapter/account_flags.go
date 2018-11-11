package resourceadapter

import (
	"github.com/fonero-project/fonero-golang/services/horizon/internal/db2/core"
	. "github.com/fonero-project/fonero-golang/protocols/horizon"
)

func PopulateAccountFlags(dest *AccountFlags, row core.Account) {
	dest.AuthRequired = row.IsAuthRequired()
	dest.AuthRevocable = row.IsAuthRevocable()
	dest.AuthImmutable = row.IsAuthImmutable()
}
