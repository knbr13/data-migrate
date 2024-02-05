package main

import (
	"fmt"
	"os"
	"strings"
)

func isDir(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}

func buildSqlInsertQuery(tableName string, columns []string) string {
	s := strings.Join(columns, ", ")
	qs := strings.TrimSuffix(strings.Repeat("?, ", len(columns)), ", ")
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", tableName, s, qs)
}
