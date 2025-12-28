const { spawn } = require('child_process');
const path = require('path');
const readline = require('readline');

class EmojiDB {
    constructor(enginePath = null) {
        this.enginePath = enginePath || path.join(__dirname, '..', 'emojidb-engine');
        this.process = null;
        this.rl = null;
        this.pending = new Map();
    }

    async connect() {
        return new Promise((resolve, reject) => {
            this.process = spawn(this.enginePath);

            this.rl = readline.createInterface({
                input: this.process.stdout,
                terminal: false
            });

            this.rl.on('line', (line) => {
                try {
                    const res = JSON.parse(line);
                    const p = this.pending.get(res.id);
                    if (p) {
                        if (res.error) p.reject(new Error(res.error));
                        else p.resolve(res.data);
                        this.pending.delete(res.id);
                    }
                } catch (e) {
                    console.error('Failed to parse engine response:', e);
                }
            });

            this.process.stderr.on('data', (data) => {
                console.error(`Engine Error: ${data}`);
            });

            this.process.on('error', (err) => {
                reject(new Error(`Failed to start engine: ${err.message}`));
            });

            // Simple delay to ensure process is ready
            setTimeout(resolve, 100);
        });
    }

    async send(method, params = {}) {
        const id = Math.random().toString(36).substring(7);
        return new Promise((resolve, reject) => {
            this.pending.set(id, { resolve, reject });
            const payload = JSON.stringify({ id, method, params });
            this.process.stdin.write(payload + '\n');
        });
    }

    async open(dbPath, key) {
        return this.send('open', { path: dbPath, key });
    }

    async defineSchema(table, fields) {
        return this.send('define_schema', { table, fields });
    }

    async insert(table, row) {
        return this.send('insert', { table, row });
    }

    async query(table, match = {}) {
        return this.send('query', { table, match });
    }

    async secure() {
        return this.send('secure');
    }

    async rekey(newKey, masterKey) {
        return this.send('rekey', { new_key: newKey, master_key: masterKey });
    }

    async close() {
        await this.send('close');
        this.process.kill();
    }
}

module.exports = EmojiDB;
