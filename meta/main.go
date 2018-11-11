// Package meta provides helpers for processing the metadata that is produced by
// fonero-core while processing transactions.
package meta

import "github.com/fonero-project/fonero-golang/xdr"

// Bundle represents all of the metadata emitted from the application of a single
// fonero transaction; Both fee meta and result meta is included.
type Bundle struct {
	FeeMeta         xdr.LedgerEntryChanges
	TransactionMeta xdr.TransactionMeta
}
