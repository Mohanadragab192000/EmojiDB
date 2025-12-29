# EmojiDB

**High-performance encrypted database engine with Node.js SDK**

EmojiDB is an embedded database written in Go that encrypts all data, headers, and schemas into emoji sequences. Features automatic binary downloads, schema evolution, and military-grade encryption.

## Quick Start

### Installation
```bash
npm install @ikwerre-dev/emojidb
```

The engine binary is automatically downloaded for your platform (Mac, Linux, Windows).

### Basic Usage
```javascript
import EmojiDB from '@ikwerre-dev/emojidb';

const db = new EmojiDB();
await db.connect();
await db.open('my_app.db', 'super-secret-key');
```

## Schema Management

Define schemas before storing data to enforce structure and data integrity.

### Defining Schemas
```javascript
await db.defineSchema('users', [
    { Name: 'id',       Type: 0, Unique: true  },
    { Name: 'username', Type: 1, Unique: true  }
]);
```

### Field Types
| Type ID | Data Type | Example |
|---------|-----------|---------|
| `0` | Integer | `123` |
| `1` | String | `"robinson"` |
| `2` | Boolean | `true` |
| `3` | Float | `10.5` |
| `4` | Map | `{ "a": 1 }` |

Schemas are persisted as readable JSON files in `emojidb/*.schema.json`.

## Schema Evolution

### Automatic Migration
```javascript
await db.migrate();
```
Syncs all tables from your local `emojidb/*.schema.json` file.

### Explicit Migration
```javascript
await db.migrate('users', [...fields]);
```

### Force Migration (Destructive)
```javascript
await db.migrate('users', true);
```
Validates all rows and drops any that don't match the new schema or contain duplicates.

### Pull Schema
```javascript
await db.pull();
```
Regenerates local schema files from the current database state.

## Data Operations

### Insert
```javascript
await db.insert('users', {
    id: 1,
    username: 'emoji_king',
    active: true
});
```

### Query
```javascript
const users = await db.query('users', { id: 1 });
console.log(users);
// Output: [{ id: 1, username: 'emoji_king', active: true }]
```

### Update
```javascript
await db.update('users', { id: 1 }, { username: 'robinson_honour' });
```

### Delete
```javascript
await db.delete('users', { id: 1 });
```

## Utilities

### Count Records
```javascript
const count = await db.count('users', { active: true });
```

### Drop Table
```javascript
await db.dropTable('logs');
```

### Force Persist to Disk
```javascript
await db.flush('users');
```

## Security

EmojiDB provides military-grade encryption:

- **AES-GCM Encryption**: All data encrypted at rest
- **Emoji Encoding**: Ciphertext encoded as emojis for obfuscation
- **Master Key Rotation**: Built-in support via `db.rekey()`

### Security Files
All database artifacts are stored in the `emojidb/` directory:
- `*.db`: Encrypted data
- `*.safety`: Crash recovery logs
- `secure.pem`: Optional master key file

## Platform Support

Automated builds for:
- **macOS**: ARM64 (M1/M2/M3) and Intel x64
- **Linux**: x64 and ARM64
- **Windows**: x64 and ARM64

---

*Built by Robinson Honour*