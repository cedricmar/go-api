package migrate

import (
	"fmt"
	"io/ioutil"
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
	if len(migs) > 0 {
		for _, file := range migs {
			// Read file
			dat, err := ioutil.ReadFile(dir + "/" + file)
			utils.LogFatal(err)

			// Get Statements
			stmts := strings.Split(string(dat), ";")
			stmts = stmts[:len(stmts)-1]

			// Run Queries
			fmt.Println("Running migration", file)
			for _, stmt := range stmts {
				_, err = db.Exec(stmt)
				utils.LogFatal(err)
			}
		}
	} else {
		fmt.Println("No migration(s) to run")
	}
	db.Close()
	fmt.Println("migration finished with status")
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
