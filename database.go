package main

import (
	"database/sql"
)

var db *sql.DB

func insert(query string, args ...any) error {
	_, err := db.Exec(query, args...)
	return err
}

func txInsert(tx *sql.Tx, query string, args ...any) error {
	_, err := tx.Exec(query, args...)
	return err
}
