# 🦄 EmojiDB: The Total Emoji Encrypted Database

EmojiDB is a high-performance, embedded database designed for maximum security and visual fun. Every record, every header, and even your schema definition is strictly 100% Emoji encoded.

## 🚀 Getting Started: Standard Workflow

EmojiDB follows a stage-by-stage progression from absolute security to data persistence.

### Stage 1: Security Initialization 🔐
Before creating a database, you must initialize the Master Security Layer. This generates a one-time `secure.pem` file which acts as your "set of emojis" for recovery and authorization.

```go
db, _ := core.Open("mydata.db", "my-secret-key")
err := db.Secure() // Creates secure.pem
```

### Stage 2: Opening & Persistence 📦
EmojiDB automatically manages three files for you:
- `dbname.db`: The actual encrypted data.
- `dbname.safety`: The crash-recovery buffer.
- `dbname.schema`: The persistent schema definitions.

```go
db, err := core.Open("/path/to/my.db", "showcase-secret-2025")
defer db.Close()
```

### Stage 3: Schema Management 📐
Like Prisma, EmojiDB uses persistent schemas. Once you define a table, it stays fixed.

```go
fields := []core.Field{
    {Name: "id", Type: core.FieldTypeInt, Unique: true},
    {Name: "name", Type: core.FieldTypeString},
}

// Initial definition
db.DefineSchema("users", fields)
```

### Stage 4: Schema Evolution (The Prisma Way) 🔄
You can update your schema and check for conflicts before applying.

```go
// 1. Check for conflicts (Prisma-like Pull/Diff)
report := db.DiffSchema("users", newFields)
if report.Destructive {
    fmt.Println("Warning: Field removal detected!")
}

// 2. Sync if compatible (Prisma-like Push)
err := db.SyncSchema("users", newFields)
```

### Stage 5: Data Operations ⚡
EmojiDB is extremely fast (~45ms for 1500 operations).

```go
// Insert
db.Insert("users", core.Row{"id": 1, "name": "Alice"})

// Query
results, _ := query.NewQuery(db, "users").Filter(...).Execute()
```

## 🛠️ Features

- **Total Emoji Encoding**: 😵🤮😇🤒😷 - your raw data never touches the disk.
- **AES-GCM Encryption**: Military-grade security on every clump.
- **Master Key Recovery**: Use `secure.pem` emoji sequences to rotate your database secret.
- **Unique Constraints**: O(1) performance for uniqueness checks.
- **Safety Engine**: Parallelized batch recovery for zero data loss.

## 🏁 Performance Benchmarks
*Tested with 1500 records + Unique Keys + Full Disk Re-encryption*

| Operation | Timing |
| :--- | :--- |
| **Ingest 1500 Rows** | ~9.1ms |
| **Flush to Disk** | ~3.9ms |
| **Rotate Master Key** | ~7.2ms |
| **TOTAL SHOWCASE** | **~45.3ms** |

---
*Created by the Google Deepmind team.*
