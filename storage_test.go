package emojidb

import (
	"os"
	"testing"
)

func TestPersistence(t *testing.T) {
	dbPath := "test_persist.db"
	defer os.Remove(dbPath)

	// Step 1: Create DB and insert data
	db, err := Open(dbPath, "secret", false)
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}

	fields := []Field{{Name: "id", Type: FieldTypeInt}}
	db.DefineSchema("items", fields)

	table := db.tables["items"]
	table.HotHeap.MaxRows = 2 // Small limit to trigger sealing

	db.Insert("items", Row{"id": 1})
	db.Insert("items", Row{"id": 2}) // Should trigger seal and persist

	db.Close()

	// Step 2: Reopen and load
	db2, err := Open(dbPath, "secret", false)
	if err != nil {
		t.Fatalf("failed to reopen db: %v", err)
	}
	defer db2.Close()

	// We need to define the schema again in the MVP because we don't persist schemas yet
	// (or rather, we don't load them automatically yet)
	db2.DefineSchema("items", fields)

	err = db2.Load()
	if err != nil {
		t.Fatalf("failed to load db: %v", err)
	}

	table2, ok := db2.tables["items"]
	if !ok {
		t.Fatal("table items not found after load")
	}

	if len(table2.SealedClumps) != 1 {
		t.Errorf("expected 1 sealed clump, got %d", len(table2.SealedClumps))
	}

	row := table2.SealedClumps[0].Rows[0]
	if int(row["id"].(float64)) != 1 { // JSON unmarshals ints as floats
		t.Errorf("expected id 1, got %v", row["id"])
	}
}
