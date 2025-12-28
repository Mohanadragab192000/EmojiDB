package storage

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"errors"
	"io"
	"os"
	"sync"

	"github.com/ikwerre-dev/emojidb/crypto"
)

const MagicRaw = "EMOJI"

func WriteHeader(file *os.File) error {
	mEncoded := crypto.EncodeToEmojis([]byte(MagicRaw))

	vBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(vBytes, 1)
	vEncoded := crypto.EncodeToEmojis(vBytes)

	_, err := file.WriteString(mEncoded + vEncoded)
	return err
}

func PersistClump(file *os.File, mu *sync.RWMutex, tableName string, clump interface{}, key string, encryptFn func([]byte, string) ([]byte, error), encodeFn func([]byte) string) error {
	mu.Lock()
	defer mu.Unlock()

	if err := InternalPersistClump(file, tableName, clump, key, encryptFn, encodeFn); err != nil {
		return err
	}
	return file.Sync()
}

func InternalPersistClump(file *os.File, tableName string, clump interface{}, key string, encryptFn func([]byte, string) ([]byte, error), encodeFn func([]byte) string) error {
	_, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}

	data, err := json.Marshal(clump)
	if err != nil {
		return err
	}

	encrypted, err := encryptFn(data, key)
	if err != nil {
		return err
	}
	payloadEncoded := encodeFn(encrypted)

	tbNameBytes := []byte(tableName)
	tbNameEncoded := encodeFn(tbNameBytes)

	tbLenBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(tbLenBytes, uint32(len(tbNameBytes)))
	tbLenEncoded := encodeFn(tbLenBytes)

	pLenBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(pLenBytes, uint32(len(encrypted)))
	pLenEncoded := encodeFn(pLenBytes)

	_, err = file.WriteString(tbLenEncoded + tbNameEncoded + pLenEncoded + payloadEncoded)
	return err
}

func Load(file *os.File, mu *sync.RWMutex, key string, decryptFn func([]byte, string) ([]byte, error), handleClump func(string, []byte) error) error {
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	br := bufio.NewReader(file)

	magic, err := readEmojis(br, 5)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return WriteHeader(file)
		}
		return err
	}

	if string(magic) != MagicRaw {
		return errors.New("invalid database magic")
	}

	vBytes, err := readEmojis(br, 4)
	if err != nil {
		return err
	}
	_ = binary.LittleEndian.Uint32(vBytes)

	for {
		tlBytes, err := readEmojis(br, 4)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		tbLen := binary.LittleEndian.Uint32(tlBytes)

		nameBytes, err := readEmojis(br, int(tbLen))
		if err != nil {
			return err
		}
		tableName := string(nameBytes)

		plBytes, err := readEmojis(br, 4)
		if err != nil {
			return err
		}
		pLen := binary.LittleEndian.Uint32(plBytes)

		payloadBytes, err := readEmojis(br, int(pLen))
		if err != nil {
			return err
		}

		decrypted, err := decryptFn(payloadBytes, key)
		if err != nil {
			return err
		}

		if err := handleClump(tableName, decrypted); err != nil {
			return err
		}
	}
	return nil
}

func readEmojis(r *bufio.Reader, count int) ([]byte, error) {
	var res []byte
	for i := 0; i < count; i++ {
		b, err := crypto.DecodeOne(r)
		if err != nil {
			return nil, err
		}
		res = append(res, b)
	}
	return res, nil
}
