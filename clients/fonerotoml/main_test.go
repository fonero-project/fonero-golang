package fonerotoml

import "log"

// ExampleGetTOML gets the fonero.toml file for coins.asia
func ExampleClient_GetFoneroToml() {
	_, err := DefaultClient.GetFoneroToml("coins.asia")
	if err != nil {
		log.Fatal(err)
	}
}
