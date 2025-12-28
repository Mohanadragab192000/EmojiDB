package emojidb

import (
	"os"
	"sync"
)

type Config struct {
	MemoryLimitMB   int
	ClumpSizeMB     int
	FlushIntervalMS int
	Encrypt         bool
}

type Database struct {
	mu      sync.RWMutex
	path    string
	key     string
	file    *os.File
	config  *Config
	schemas map[string]*Schema
	tables  map[string]*Table
}

type Table struct {
	mu           sync.RWMutex
	db           *Database
	Name         string
	Schema       *Schema
	HotHeap      *HotHeap
	SealedClumps []*SealedClump
}

func Open(path, key string, encrypt bool) (*Database, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	db := &Database{
		path:    path,
		key:     key,
		file:    file,
		config:  &Config{Encrypt: encrypt},
		schemas: make(map[string]*Schema),
		tables:  make(map[string]*Table),
	}

 	if err := db.Load(); err != nil {
		file.Close()
		return nil, err
	}
	return db, nil
}

func (db *Database) Close() error {
	db.mu.Lock()
	defer db.mu.Unlock()
	if db.file != nil {
		return db.file.Close()
	}
	return nil
}
