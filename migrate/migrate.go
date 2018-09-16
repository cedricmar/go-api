package migrate

import (
	"log"

	"github.com/cedricmar/go-api/utils"
	"github.com/cedricmar/go-api/utils/cli"
)

// Migrations are the schema updates
// @TODO - Avoid keeping a registry
var Migrations = map[int]func(){
	1: InitSchema,
}

// Seeds are the data updates
// @TODO - Avoid keeping a registry
var Seeds = map[int]func(){
	1: cli.Seed,
}

func Down() {
	run(false)
}

func Up() {
	run(true)
}

func run(u bool) {
	// Execute migrations
	db := utils.DBConnect()
	defer db.Close()

	log.Println(Migrations)

	if len(Migrations) > 0 {
		log.Println("--- Running Migration ----")
		for i := 0; i < len(Migrations); i++ {
			version := i + 1
			log.Println("--- Running version", version, "---")
			// Run migration
			Migrations[version]()
			// Seed
			log.Println(Seeds)
			Seeds[version]()
		}
	} else {
		log.Println("No migration(s) to run")
	}
	log.Println("--- Migrations - Done ---")
}
