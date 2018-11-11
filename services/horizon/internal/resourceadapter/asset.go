package resourceadapter

import (
	"context"

	"github.com/fonero-project/fonero-golang/xdr"
	. "github.com/fonero-project/fonero-golang/protocols/horizon"

)

func PopulateAsset(ctx context.Context, dest *Asset, asset xdr.Asset) error {
	return asset.Extract(&dest.Type, &dest.Code, &dest.Issuer)
}
