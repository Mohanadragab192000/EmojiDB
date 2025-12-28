import { BinaryManager } from './index.js';

console.log('📦 EmojiDB Post-Install: Checking for engine binary...');

const manager = new BinaryManager();

manager.ensureBinary()
    .then((path) => {
        console.log(`✅ EmojiDB Engine installed at: ${path}`);
        process.exit(0);
    })
    .catch((err) => {
        console.error(`❌ EmojiDB Engine Download Failed: ${err.message}`);
        console.error('   You may need to check your internet connection or manually install the binary.');
        // We exit with 0 to not break the entire npm install process, 
        // allowing the user to try again or use the library if they have the binary elsewhere.
        process.exit(0);
    });
