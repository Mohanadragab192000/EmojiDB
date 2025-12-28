package emojidb

import (
	"os"
	"testing"
)

func TestQuery(t *testing.T) {
	dbPath := "test_query.db"
	defer os.Remove(dbPath)

	db, err := Open(dbPath, "secret", false)
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	fields := []Field{
		{Name: "id", Type: FieldTypeInt},
		{Name: "name", Type: FieldTypeString},
		{Name: "age", Type: FieldTypeInt},
	}
	db.DefineSchema("users", fields)

	db.Insert("users", Row{"id": 1, "name": "alice", "age": 30})
	db.Insert("users", Row{"id": 2, "name": "bob", "age": 25})
	db.Insert("users", Row{"id": 3, "name": "charlie", "age": 35})

	// Test 1: Simple Filter
	results, err := db.Query("users").
		Filter(func(r Row) bool {
			return int(r["age"].(int)) > 28
		}).
		Execute()

	if err != nil {
		t.Fatalf("query failed: %v", err)
	}

	if len(results) != 2 {
		t.Errorf("expected 2 results, got %d", len(results))
	}

	// Test 2: Select Projection
	results, _ = db.Query("users").
		Filter(func(r Row) bool { return r["name"] == "bob" }).
		Select("name").
		Execute()

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if len(results[0]) != 1 {
		t.Errorf("expected projected row to have 1 column, got %d", len(results[0]))
	}

	if _, ok := results[0]["age"]; ok {
		t.Error("age column should have been filtered out by Select")
	}
}
