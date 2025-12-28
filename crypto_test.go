package emojidb

import (
	"bytes"
	"testing"
)

func TestCrypto(t *testing.T) {
	db := &Database{key: "test-secret"}
	data := []byte("hello world")

	encrypted, err := db.encrypt(data)
	if err != nil {
		t.Fatalf("encryption failed: %v", err)
	}

	if bytes.Equal(encrypted, data) {
		t.Fatal("ciphertext matches plaintext")
	}

	decrypted, err := db.decrypt(encrypted)
	if err != nil {
		t.Fatalf("decryption failed: %v", err)
	}

	if !bytes.Equal(decrypted, data) {
		t.Fatalf("expected %s, got %s", string(data), string(decrypted))
	}
}

func TestEmojiEncoding(t *testing.T) {
	db := &Database{}
	data := []byte{0, 1, 2, 255}

	if len(emojiAlphabet) < 256 {
		t.Logf("Warning: emoji alphabet size is %d, expected at least 256", len(emojiAlphabet))
	}

	encoded := db.encodeToEmojis(data)
	if len(encoded) <= len(data) {
		t.Errorf("expected encoded length to be greater than data length")
	}

	decoded, err := db.decodeFromEmojis(encoded)
	if err != nil {
		t.Fatalf("decoding failed: %v", err)
	}

	if !bytes.Equal(decoded, data) {
		t.Errorf("expected %v, got %v", data, decoded)
	}
}
