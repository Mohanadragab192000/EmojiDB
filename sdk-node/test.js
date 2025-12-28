const EmojiDB = require('./index');
const path = require('path');

async function runTest() {
    // Standalone mode: auto-downloads engine from GitHub if not found
    const db = new EmojiDB();

    try {
        console.log('--- EMOJIDB NODE.JS TEST ---');

        await db.connect();
        console.log('1. Engine Connected');

        await db.open('node_showcase.db', 'node-secret-2025');
        console.log('2. Database Opened');

        await db.defineSchema('users', [
            { Name: 'id', Type: 0, Unique: true },
            { Name: 'username', Type: 1, Unique: false }
        ]);
        console.log('3. Schema Defined');

        await db.insert('users', { id: 101, username: 'emoji_king' });
        await db.insert('users', { id: 102, username: 'node_master' });
        console.log('4. Data Inserted');

        const results = await db.query('users', { id: 101 });
        console.log('5. Query Results:', results);

        if (results.length > 0 && results[0].username === 'emoji_king') {
            console.log('✅ TEST PASSED: Node.js successfully interacted with Go core!');
        } else {
            console.log('❌ TEST FAILED: Data mismatch');
        }

        await db.close();
        console.log('6. Connection Closed');

    } catch (err) {
        console.error('❌ TEST ERROR:', err.message);
    }
}

runTest();
