package database

import "fmt"

var (
	userName = "root"
	password = "root"
	host = "localhost"
	port = "3306"
	dbName = "demo"
	mySqlConnStr = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", userName, password, host, port, dbName)
)