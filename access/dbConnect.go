package access

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/wscherfel/fitlogic-backend"
)

func ConnectToDb() (*gorm.DB, error) {
	return gorm.Open("sqlite3", fitlogic.DbName)
}
