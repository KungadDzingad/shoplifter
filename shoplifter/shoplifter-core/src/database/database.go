package database

import (
	commonDB "github.com/KungadDzingad/shoplifter-common/database"
)

var DB commonDB.Dbinstance

func ConnectDb() {
	DB = commonDB.ConnectDb()
}
