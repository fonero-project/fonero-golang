package helpers

import (
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/fonero-project/fonero-golang/address"
	"github.com/fonero-project/fonero-golang/amount"
	"github.com/fonero-project/fonero-golang/strkey"
)

func init() {
	govalidator.CustomTypeTagMap.Set("fonero_accountid", govalidator.CustomTypeValidator(isFoneroAccountID))
	govalidator.CustomTypeTagMap.Set("fonero_seed", govalidator.CustomTypeValidator(isFoneroSeed))
	govalidator.CustomTypeTagMap.Set("fonero_asset_code", govalidator.CustomTypeValidator(isFoneroAssetCode))
	govalidator.CustomTypeTagMap.Set("fonero_address", govalidator.CustomTypeValidator(isFoneroAddress))
	govalidator.CustomTypeTagMap.Set("fonero_amount", govalidator.CustomTypeValidator(isFoneroAmount))
	govalidator.CustomTypeTagMap.Set("fonero_destination", govalidator.CustomTypeValidator(isFoneroDestination))

}

func Validate(request Request, params ...interface{}) error {
	valid, err := govalidator.ValidateStruct(request)

	if !valid {
		fields := govalidator.ErrorsByField(err)
		for field, errorValue := range fields {
			switch {
			case errorValue == "non zero value required":
				return NewMissingParameter(field)
			case strings.HasSuffix(errorValue, "does not validate as fonero_accountid"):
				return NewInvalidParameterError(field, "Account ID must start with `G` and contain 56 alphanum characters.")
			case strings.HasSuffix(errorValue, "does not validate as fonero_seed"):
				return NewInvalidParameterError(field, "Account secret must start with `S` and contain 56 alphanum characters.")
			case strings.HasSuffix(errorValue, "does not validate as fonero_asset_code"):
				return NewInvalidParameterError(field, "Asset code must be 1-12 alphanumeric characters.")
			case strings.HasSuffix(errorValue, "does not validate as fonero_address"):
				return NewInvalidParameterError(field, "Fonero address must be of form user*domain.com")
			case strings.HasSuffix(errorValue, "does not validate as fonero_destination"):
				return NewInvalidParameterError(field, "Fonero destination must be of form user*domain.com or start with `G` and contain 56 alphanum characters.")
			case strings.HasSuffix(errorValue, "does not validate as fonero_amount"):
				return NewInvalidParameterError(field, "Amount must be positive and have up to 7 decimal places.")
			default:
				return NewInvalidParameterError(field, errorValue)
			}
		}
	}

	return request.Validate(params...)
}

// These are copied from support/config. Should we move them to /strkey maybe?
func isFoneroAccountID(i interface{}, context interface{}) bool {
	enc, ok := i.(string)

	if !ok {
		return false
	}

	_, err := strkey.Decode(strkey.VersionByteAccountID, enc)

	if err == nil {
		return true
	}

	return false
}

func isFoneroSeed(i interface{}, context interface{}) bool {
	enc, ok := i.(string)

	if !ok {
		return false
	}

	_, err := strkey.Decode(strkey.VersionByteSeed, enc)

	if err == nil {
		return true
	}

	return false
}

func isFoneroAssetCode(i interface{}, context interface{}) bool {
	code, ok := i.(string)

	if !ok {
		return false
	}

	if !govalidator.IsByteLength(code, 1, 12) {
		return false
	}

	if !govalidator.IsAlphanumeric(code) {
		return false
	}

	return true
}

func isFoneroAddress(i interface{}, context interface{}) bool {
	addr, ok := i.(string)

	if !ok {
		return false
	}

	_, _, err := address.Split(addr)
	if err != nil {
		return false
	}

	return true
}

func isFoneroAmount(i interface{}, context interface{}) bool {
	am, ok := i.(string)

	if !ok {
		return false
	}

	_, err := amount.Parse(am)
	if err != nil {
		return false
	}

	return true
}

// isFoneroDestination checks if `i` is either account public key or Fonero address.
func isFoneroDestination(i interface{}, context interface{}) bool {
	dest, ok := i.(string)

	if !ok {
		return false
	}

	_, err1 := strkey.Decode(strkey.VersionByteAccountID, dest)
	_, _, err2 := address.Split(dest)

	if err1 != nil && err2 != nil {
		return false
	}

	return true
}
