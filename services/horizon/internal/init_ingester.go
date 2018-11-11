package horizon

import (
	"log"

	"github.com/fonero-project/fonero-golang/services/horizon/internal/ingest"
)

func initIngester(app *App) {
	if !app.config.Ingest {
		return
	}

	if app.networkPassphrase == "" {
		log.Fatal("Cannot start ingestion without network passphrase.  Please confirm connectivity with fonero-core.")
	}

	app.ingester = ingest.New(
		app.networkPassphrase,
		app.config.FoneroCoreURL,
		app.CoreSession(nil),
		app.HorizonSession(nil),
		ingest.Config{
			EnableAssetStats: app.config.EnableAssetStats,
		},
	)

	app.ingester.SkipCursorUpdate = app.config.SkipCursorUpdate
	app.ingester.HistoryRetentionCount = app.config.HistoryRetentionCount
}

func init() {
	appInit.Add("ingester", initIngester, "app-context", "log", "horizon-db", "core-db", "foneroCoreInfo")
}
