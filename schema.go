package emojidb

import (
	"errors"
)

type FieldType int

const (
	FieldTypeInt FieldType = iota
	FieldTypeString
	FieldTypeFloat
	FieldTypeBool
)

type Field struct {
	Name string
	Type FieldType
}

type Schema struct {
	Version int
	Fields  []Field
}

func (db *Database) DefineSchema(tableName string, fields []Field) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.schemas[tableName]; ok {
		return errors.New("schema already defined for table: " + tableName)
	}

	schema := &Schema{
		Version: 1,
		Fields:  fields,
	}

	db.schemas[tableName] = schema
	db.tables[tableName] = &Table{
		db:      db,
		Name:    tableName,
		Schema:  schema,
		HotHeap: NewHotHeap(1000), // default limit for MVP
	}

	return nil
}
