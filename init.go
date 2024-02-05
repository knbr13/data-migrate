package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
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

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating database connection: %s\n", err.Error())
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating database connection: %s\n", err.Error())
		os.Exit(1)
	}
}
