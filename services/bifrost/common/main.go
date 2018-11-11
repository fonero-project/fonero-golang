package common

import (
	"github.com/fonero-project/fonero-golang/support/log"
)

const FoneroAmountPrecision = 7

func CreateLogger(serviceName string) *log.Entry {
	return log.DefaultLogger.WithField("service", serviceName)
}
