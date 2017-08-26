package migrate

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/weebagency/go-api/utils"
)

func Down() {
	Run(getDownList)
}

func Up() {
	Run(getUpList)
}

func Run(filter func([]os.FileInfo) []string) {
	dir, err := filepath.Abs("./migrate/")
	utils.LogFatal(err)

	// Get migration files
	files, err := ioutil.ReadDir(dir)
	utils.LogFatal(err)
	migs := filter(files)

	// Execute migrations
	db := utils.DBConnect()
	defer db.Close()
	if len(migs) > 0 {
		log.Println("--- Running Migration ----")
		for _, file := range migs {
			// Read file
			schema, err := ioutil.ReadFile(dir + "/" + file)
			utils.LogFatal(err)
			// Run Queries
			log.Println("---", file)
			db.MustExec(string(schema))
		}
	} else {
		log.Println("No migration(s) to run")
	}
	log.Println("--- Migrations - Done ---")
}

func getUpList(files []os.FileInfo) []string {
	return getList(".up.sql", files)
}

func getDownList(files []os.FileInfo) []string {
	return getList(".down.sql", files)
}

func getList(pat string, files []os.FileInfo) []string {
	var list []string
	for _, file := range files {
		if strings.Contains(file.Name(), pat) {
			list = append(list, file.Name())
		}
	}
	return list
}
