# EmojiDB

EmojiDB is a fast, memory-efficient embedded database engine that batches data into clumps, performs computation during ingestion, and supports emoji-based encrypted persistence.

## Features

- Memory-first writes
- Append-only persistence
- Schema-aware relational core
- Transparent encryption with emoji encoding
- Function-based fluent query API

## Usage

```go
db, err := emojidb.Open("data.db", "secret-key", true)
if err != nil {
    log.Fatal(err)
}
defer db.Close()
```
