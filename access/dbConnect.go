package access

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// ConnectToDb will return connected DB, this is the only place
// programmer would have to rewrite if customer wants DB with some other
// technology (postgres or mysql)
func ConnectToDb() (*gorm.DB, error) {
	return gorm.Open("sqlite3", "fitlogic.db")
}
