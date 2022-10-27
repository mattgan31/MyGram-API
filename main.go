package main

import (
	"final-project-fga/database"
	"final-project-fga/router"
)

const PORT = ":3000"

func main() {
	database.StartDB()

	router.StartServer().Run(PORT)

}
