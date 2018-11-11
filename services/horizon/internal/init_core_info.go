package horizon

func initFoneroCoreInfo(app *App) {
	app.UpdateFoneroCoreInfo()
	return
}

func init() {
	appInit.Add("foneroCoreInfo", initFoneroCoreInfo, "app-context", "log")
}
