package emojidb

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"io"
)

const (
	MagicNumber = "EMOJI" // 5 bytes
	Version     = 1
)

type FileHeader struct {
	Magic   [5]byte
	Version uint32
}

func (db *Database) writeHeader() error {
	var header FileHeader
	copy(header.Magic[:], MagicNumber)
	header.Version = Version

	err := binary.Write(db.file, binary.LittleEndian, header)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) persistClump(tableName string, clump *SealedClump) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	_, err := db.file.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}

	data, err := json.Marshal(clump)
	if err != nil {
		return err
	}

	// if encryption is enabled, encrypt the data
	var finalData []byte
	isEncrypted := db.config.Encrypt && db.key != ""

	if isEncrypted {
		encrypted, err := db.encrypt(data)
		if err != nil {
			return err
		}
		emojiPayload := db.encodeToEmojis(encrypted)
		finalData = []byte(emojiPayload)
	} else {
		finalData = data
	}

	// write table name
	tbNameBytes := []byte(tableName)
	if err := binary.Write(db.file, binary.LittleEndian, uint32(len(tbNameBytes))); err != nil {
		return err
	}
	if _, err := db.file.Write(tbNameBytes); err != nil {
		return err
	}

	// write encrypted flag
	var encFlag uint8
	if isEncrypted {
		encFlag = 1
	}
	if err := binary.Write(db.file, binary.LittleEndian, encFlag); err != nil {
		return err
	}

	// payload size
	if err := binary.Write(db.file, binary.LittleEndian, uint32(len(finalData))); err != nil {
		return err
	}
	// write payload
	if _, err := db.file.Write(finalData); err != nil {
		return err
	}

	return db.file.Sync()
}

func (db *Database) Load() error {
	_, err := db.file.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	var header FileHeader
	err = binary.Read(db.file, binary.LittleEndian, &header)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return db.writeHeader()
		}
		return err
	}

	if string(header.Magic[:]) != MagicNumber {
		return errors.New("invalid database file format")
	}

	for {
		var nameLen uint32
		err := binary.Read(db.file, binary.LittleEndian, &nameLen)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		nameBytes := make([]byte, nameLen)
		if _, err := io.ReadFull(db.file, nameBytes); err != nil {
			return err
		}
		tableName := string(nameBytes)

		var encFlag uint8
		if err := binary.Read(db.file, binary.LittleEndian, &encFlag); err != nil {
			return err
		}

		var dataLen uint32
		if err := binary.Read(db.file, binary.LittleEndian, &dataLen); err != nil {
			return err
		}

		data := make([]byte, dataLen)
		if _, err := io.ReadFull(db.file, data); err != nil {
			return err
		}

		var finalData []byte
		if encFlag == 1 {
			if db.key == "" {
				return errors.New("database is encrypted but no key provided")
			}
			encrypted, err := db.decodeFromEmojis(string(data))
			if err != nil {
				return err
			}
			decrypted, err := db.decrypt(encrypted)
			if err != nil {
				return err
			}
			finalData = decrypted
		} else {
			finalData = data
		}

		var clump SealedClump
		if err := json.Unmarshal(finalData, &clump); err != nil {
			return err
		}

		db.mu.Lock()
		table, ok := db.tables[tableName]
		if ok {
			table.SealedClumps = append(table.SealedClumps, &clump)
		}
		db.mu.Unlock()
	}

	return nil
}
