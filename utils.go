package main

import (
	"fmt"
	"os"
	"path/filepath"
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

func buildSQLInsertQueryFromMap(m map[string]any, tableName string) (query string, values []any) {
	columns := make([]string, 0, len(m))
	values = make([]any, 0, len(m))
	for k, v := range m {
		columns = append(columns, k)
		values = append(values, v)
	}
	return buildSqlInsertQuery(tableName, columns), values
}

func filesInDir(dir string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".db.json") {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
