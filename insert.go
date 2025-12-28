package emojidb

import (
	"errors"
)

func (db *Database) Insert(tableName string, record Row) error {
	db.mu.RLock()
	table, ok := db.tables[tableName]
	db.mu.RUnlock()

	if !ok {
		return errors.New("table not found: " + tableName)
	}

	return table.Insert(record)
}
