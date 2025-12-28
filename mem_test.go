package emojidb

import (
	"os"
	"testing"
)

func TestInsert(t *testing.T) {
	dbPath := "test_insert.db"
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
	db.DefineSchema("users", fields)

	// Test successful insert
	err = db.Insert("users", Row{"id": 1, "name": "alice"})
	if err != nil {
		t.Fatalf("failed to insert: %v", err)
	}

	table := db.tables["users"]
	if len(table.HotHeap.Rows) != 1 {
		t.Errorf("expected 1 row in hot heap, got %d", len(table.HotHeap.Rows))
	}

	// Test missing field
	err = db.Insert("users", Row{"id": 2})
	if err == nil {
		t.Error("expected error for missing field")
	}
}

func TestClumpSealing(t *testing.T) {
	dbPath := "test_seal.db"
	defer os.Remove(dbPath)

	db, err := Open(dbPath, "secret", true)
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	fields := []Field{{Name: "id", Type: FieldTypeInt}}
	db.DefineSchema("items", fields)

	table := db.tables["items"]
	table.HotHeap.MaxRows = 3 // Small limit for testing

	db.Insert("items", Row{"id": 1})
	db.Insert("items", Row{"id": 2})
	db.Insert("items", Row{"id": 3})

	if len(table.SealedClumps) != 1 {
		t.Errorf("expected 1 sealed clump, got %d", len(table.SealedClumps))
	}
	if len(table.HotHeap.Rows) != 0 {
		t.Errorf("expected 0 rows in hot heap after sealing, got %d", len(table.HotHeap.Rows))
	}

	db.Insert("items", Row{"id": 4})
	if len(table.HotHeap.Rows) != 1 {
		t.Errorf("expected 1 row in hot heap after new insert, got %d", len(table.HotHeap.Rows))
	}
}
