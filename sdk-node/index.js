const { spawn } = require('child_process');
const path = require('path');
const readline = require('readline');
const fs = require('fs');
const https = require('https');
const os = require('os');

class BinaryManager {
    constructor() {
        this.platform = os.platform();
        this.arch = os.arch();
        this.binDir = path.join(__dirname, 'bin');
        this.engineName = `emojidb-${this.platform}-${this.arch}${this.platform === 'win32' ? '.exe' : ''}`;
        this.enginePath = path.join(this.binDir, this.engineName);
    }

    async ensureBinary() {
        if (fs.existsSync(this.enginePath)) return this.enginePath;

        console.log(`🚀 EmojiDB: Engine not found for ${this.platform}-${this.arch}. Downloading from GitHub...`);

        if (!fs.existsSync(this.binDir)) {
            fs.mkdirSync(this.binDir, { recursive: true });
        }

        const url = `https://github.com/ikwerre-dev/EmojiDB/releases/latest/download/${this.engineName}`;

        return new Promise((resolve, reject) => {
            const file = fs.createWriteStream(this.enginePath);
            https.get(url, (response) => {
                if (response.statusCode === 302) {
                    // Handle redirect (e.g. to objects.githubusercontent.com)
                    https.get(response.headers.location, (res) => res.pipe(file));
                } else if (response.statusCode !== 200) {
                    reject(new Error(`Failed to download engine: HTTP ${response.statusCode}. Please ensure a release exists at github.com/ikwerre-dev/EmojiDB`));
                    return;
                } else {
                    response.pipe(file);
                }

                file.on('finish', () => {
                    file.close();
                    if (this.platform !== 'win32') {
                        fs.chmodSync(this.enginePath, 0o755);
                    }
                    console.log('✅ EmojiDB: Engine downloaded and ready.');
                    resolve(this.enginePath);
                });
            }).on('error', (err) => {
                fs.unlink(this.enginePath, () => { });
                reject(err);
            });
        });
    }
}

class EmojiDB {
    constructor(options = {}) {
        this.manager = new BinaryManager();
        this.enginePath = options.enginePath || null; // Override for local dev
        this.process = null;
        this.rl = null;
        this.pending = new Map();
    }

    async connect() {
        if (!this.enginePath) {
            this.enginePath = await this.manager.ensureBinary();
        }

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
            if (!this.process || this.process.killed) {
                return reject(new Error("Database not connected. Call db.connect() first."));
            }
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

    async update(table, match, updateData) {
        return this.send('update', { table, match, update: updateData });
    }

    async delete(table, match) {
        return this.send('delete', { table, match });
    }

    async secure() {
        return this.send('secure');
    }

    async rekey(newKey, masterKey) {
        return this.send('rekey', { new_key: newKey, master_key: masterKey });
    }

    async close() {
        if (this.process && !this.process.killed) {
            await this.send('close');
            this.process.kill();
        }
    }
}

module.exports = EmojiDB;
