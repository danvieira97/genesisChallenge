package main

import (
	"github.com/danvieira97/genesisChallenge/database"
	"github.com/danvieira97/genesisChallenge/router"
)

func main() {
	database.ConnectDB()

	router.Initialize()
}
