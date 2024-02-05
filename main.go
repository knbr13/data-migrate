package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

func main() {
	defer db.Close()

	var tx *sql.Tx
	if useTx {
		var err error
		tx, err = db.Begin()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: starting transaction: %s\n", err.Error())
			os.Exit(1)
		}
		defer tx.Rollback()
	}

	isdir, err := isDir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	}

	var files []string

	if isdir {
		files, err = filesInDir(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
			os.Exit(1)
		}
	} else {
		files = []string{path}
	}

	for _, file := range files {
		data, err := openAndReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
			os.Exit(1)
		}
		m := orderedmap.New[string, any]()
		err = json.Unmarshal(data, &m)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
			os.Exit(1)
		}

		for pair := m.Oldest(); pair != nil; pair = pair.Next() {
			switch d := pair.Value.(type) {
			case []any:
				for _, obj := range d {
					switch v := obj.(type) {
					case map[string]any:
						query, values := buildSQLInsertQueryFromMap(v, pair.Key)
						if useTx {
							err = txInsert(tx, query, values...)
						} else {
							err = insert(query, values...)
						}
						if err != nil {
							fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
							if useTx {
								tx.Rollback()
								os.Exit(1)
							}
						}
					default:
						fmt.Fprintf(os.Stderr, "error: %s\n", "unsupported type")
						if useTx {
							tx.Rollback()
							os.Exit(1)
						}
					}
				}
			case map[string]any:
				query, values := buildSQLInsertQueryFromMap(d, pair.Key)
				if useTx {
					err = txInsert(tx, query, values...)
				} else {
					err = insert(query, values...)
				}
				if err != nil {
					fmt.Println("error:", err)
					if useTx {
						tx.Rollback()
						os.Exit(1)
					}
				}
			default:
				fmt.Fprintf(os.Stderr, "unsupported type: %T, values for table names should be object or array of objects.\n", d)
				if useTx {
					tx.Rollback()
					os.Exit(1)
				}
			}
		}
	}
	if useTx {
		err = tx.Commit()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
			os.Exit(1)
		}
	}
	fmt.Println("insertion completed!")
}
