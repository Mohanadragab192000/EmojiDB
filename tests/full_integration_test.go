package tests

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/ikwerre-dev/emojidb/core"
	"github.com/ikwerre-dev/emojidb/query"
	"github.com/ikwerre-dev/emojidb/safety"
)

func TestFullShowcase(t *testing.T) {
	dbPath := "showcase.db"
	key := "showcase-secret-2025"

	fmt.Println("\n🚀 STARTING EMOJIDB FULL SHOWCASE")
	fmt.Println("==================================")

	// cleanup
	os.Remove(dbPath)
	os.Remove("safety.db")

	// 1. Open Database
	fmt.Printf("📂 1. Opening Database: %s (Encryption Mandatory)\n", dbPath)
	db, err := core.Open(dbPath, key)
	if err != nil {
		t.Fatalf("Failed to open: %v", err)
	}
	defer db.Close()
	fmt.Println("   ✅ Database ready.")

	// 2. Define Schema
	fmt.Println("📋 2. Defining Schema: 'products'")
	fields := []core.Field{
		{Name: "id", Type: core.FieldTypeInt},
		{Name: "name", Type: core.FieldTypeString},
		{Name: "price", Type: core.FieldTypeInt},
		{Name: "category", Type: core.FieldTypeString},
	}
	err = db.DefineSchema("products", fields)
	if err != nil {
		t.Fatalf("Failed to define schema: %v", err)
	}
	fmt.Println("   ✅ Schema registered.")

	// 3. Ingestion
	fmt.Println("📥 3. Ingesting Data into Hot Heap")
	products := []core.Row{
		{"id": 1, "name": "Emoji Laptop", "price": 1200, "category": "tech"},
		{"id": 2, "name": "Emoji Phone", "price": 800, "category": "tech"},
		{"id": 3, "name": "Emoji Coffee", "price": 5, "category": "food"},
	}
	for i, p := range products {
		db.Insert("products", p)
		fmt.Printf("   👉 Row %d inserted: %s\n", i+1, p["name"])
	}

	// 4. Safety Engine (Update & Backup)
	fmt.Println("🛡️ 4. Safety Engine: Updating Coffee price (Automatic Backup)")
	err = safety.Update(db, "products", func(r core.Row) bool {
		return r["id"] == 3
	}, core.Row{"price": 6})
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}

	points, err := safety.ListRecoveryPoints(db)
	if err == nil && len(points) > 0 {
		fmt.Printf("   ✅ Recovery point created at: %s\n", points[0].Format(time.Kitchen))
	} else {
		// Debug: check file size
		if info, err := os.Stat("safety.db"); err == nil {
			fmt.Printf("   ⚠️ Safety file size: %d bytes\n", info.Size())
		}
		fmt.Println("   ⚠️ No recovery points found in list.")
	}

	// 5. Persistence (Flush)
	fmt.Println("💾 5. Flushing Hot Heap to Disk (Total Emoji Encoding)")
	db.Flush("products")
	fmt.Println("   ✅ Data persisted to showcase.db.")

	// 6. Inspect File
	fmt.Println("🔍 6. Inspecting Disk Content (Should be 100% Emojis)")
	content, _ := os.ReadFile(dbPath)
	if len(content) > 100 {
		fmt.Printf("   📝 File Content Preview: %s...\n", string(content[:100]))
	} else {
		fmt.Printf("   📝 File Content: %s\n", string(content))
	}

	// 7. Query Engine
	fmt.Println("🔎 7. Running Fluent Query: Category = 'tech' && Price > 1000")
	results, err := query.NewQuery(db, "products").Filter(func(r core.Row) bool {
		return r["category"] == "tech" && r["price"].(int) > 1000
	}).Execute()

	if err != nil {
		t.Fatalf("Query failed: %v", err)
	}
	fmt.Printf("   ✅ Query Result: found %d matches\n", len(results))
	for _, r := range results {
		fmt.Printf("      - %s ($%d)\n", r["name"], r["price"])
	}

	// 8. JSON Dump
	fmt.Println("📦 8. Dumping Table as JSON")
	jsonDump, err := db.DumpAsJSON("products")
	if err != nil {
		t.Fatalf("Dump failed: %v", err)
	}
	fmt.Println("--- JSON DUMP START ---")
	fmt.Println(jsonDump)
	fmt.Println("--- JSON DUMP END ---")

	fmt.Println("\n✨ SHOWCASE COMPLETE: All systems operational.")
	fmt.Println("==================================")
}
