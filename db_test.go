package emojidb

import (
	"os"
	"testing"
)

func TestOpen(t *testing.T) {
	dbPath := "test_open.db"
	defer os.Remove(dbPath)

	db, err := Open(dbPath, "secret", true)
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	if db.path != dbPath {
		t.Errorf("expected path %s, got %s", dbPath, db.path)
	}
	if !db.config.Encrypt {
		t.Errorf("expected encryption enabled")
	}
}

func TestDefineSchema(t *testing.T) {
	dbPath := "test_schema.db"
	defer os.Remove(dbPath)

	db, err := Open(dbPath, "secret", true)
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	fields := []Field{
		{Name: "id", Type: FieldTypeInt},
		{Name: "name", Type: FieldTypeString},
	}

	err = db.DefineSchema("users", fields)
	if err != nil {
		t.Fatalf("failed to define schema: %v", err)
	}

	schema, ok := db.schemas["users"]
	if !ok {
		t.Fatal("schema not found")
	}

	if len(schema.Fields) != 2 {
		t.Errorf("expected 2 fields, got %d", len(schema.Fields))
	}

	err = db.DefineSchema("users", fields)
	if err == nil {
		t.Error("expected error when redefining schema")
	}
}
