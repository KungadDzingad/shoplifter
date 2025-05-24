package database

import "github.com/KungadDzingad/shoplifter-common/database"

var Db database.Dbinstance

func ConnectDb() {
	Db = database.ConnectDb()
}
