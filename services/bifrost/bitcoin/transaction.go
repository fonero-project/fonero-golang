package bitcoin

import (
	"math/big"

	"github.com/fonero-project/fonero-golang/services/bifrost/common"
)

func (t Transaction) ValueToFonero() string {
	valueSat := new(big.Int).SetInt64(t.ValueSat)
	valueBtc := new(big.Rat).Quo(new(big.Rat).SetInt(valueSat), satInBtc)
	return valueBtc.FloatString(common.FoneroAmountPrecision)
}
