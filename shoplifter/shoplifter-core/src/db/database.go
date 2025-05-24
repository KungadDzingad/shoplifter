package database

import commonDB "github.com/KungadDzingad/shoplifter-common/database"

var Db commonDB.Dbinstance

func ConnectDb() {
	Db = commonDB.ConnectDb()
}
