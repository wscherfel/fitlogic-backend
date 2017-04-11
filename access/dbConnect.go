package access

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const DbName = "fitlogic.db"

func ConnectToDb() (*gorm.DB, error) {
	return gorm.Open("sqlite3", DbName)
}
