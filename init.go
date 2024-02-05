package main

import (
	"flag"
	"os"
)

var dsn, path string
var useTx bool

func init() {
	flag.StringVar(&dsn, "dsn", "", "data source name. e.g. root:toor@tcp(localhost:3306)/mysqldb")
	flag.StringVar(&path, "path", "", "path to the json file that contains the data to insert into the database")
	flag.BoolVar(&useTx, "tx", true, "use transaction to insert into the database")
	flag.Parse()

	if dsn == "" || path == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}
}
