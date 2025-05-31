package database

import (
	commonDB "github.com/KungadDzingad/shoplifter-common/database"
)

var db commonDB.Dbinstance

func Connection() *commonDB.Dbinstance {
	if db.Db == nil {
		ConnectDb()
	}
	return &db
}

func ConnectDb() {
	db = commonDB.ConnectDb()
}
