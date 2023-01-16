package main

import (
	"log"

	App "github.com/NonsoAmadi10/bitcoin-wallet/app"
)

func main() {

	err := App.StartApp().Start("localhost:8082")
	if err != nil {
		log.Fatal(err)
	}
}
