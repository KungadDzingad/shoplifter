package database

import commonDB "github.com/KungadDzingad/shoplifter-common/database"

var CONNECTION commonDB.Dbinstance

func ConnectDb() {
	CONNECTION = commonDB.ConnectDb()
}
