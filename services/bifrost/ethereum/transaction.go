package ethereum

import (
	"math/big"

	"github.com/fonero-project/fonero-golang/services/bifrost/common"
)

func (t Transaction) ValueToFonero() string {
	valueEth := new(big.Rat)
	valueEth.Quo(new(big.Rat).SetInt(t.ValueWei), weiInEth)
	return valueEth.FloatString(common.FoneroAmountPrecision)
}
