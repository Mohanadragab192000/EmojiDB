package emojidb

import (
	"errors"
	"time"
)

type Row map[string]interface{}

type HotHeap struct {
	Rows      []Row
	Size      int
	MaxRows   int
	CreatedAt time.Time
}

type SealedClump struct {
	Rows     []Row
	Metadata ClumpMetadata
	SealedAt time.Time
}

type ClumpMetadata struct {
	RowCount      int
	SchemaVersion int
	CreatedAt     time.Time
}

func NewHotHeap(maxRows int) *HotHeap {
	return &HotHeap{
		Rows:      make([]Row, 0, maxRows),
		MaxRows:   maxRows,
		CreatedAt: time.Now(),
	}
}

func (t *Table) Insert(record Row) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.HotHeap == nil {
		t.HotHeap = NewHotHeap(1000) // leaving this limit here sinc we're just testing
	}

	for _, field := range t.Schema.Fields {
		val, ok := record[field.Name]
		if !ok {
			return errors.New("missing field: " + field.Name)
		}
		// when i have strength, i'd add type valiudation here
		_ = val
	}

	t.HotHeap.Rows = append(t.HotHeap.Rows, record)

	if len(t.HotHeap.Rows) >= t.HotHeap.MaxRows {
		t.sealHotHeap()
	}

	return nil
}

func (t *Table) sealHotHeap() {
	clump := &SealedClump{
		Rows:     t.HotHeap.Rows,
		SealedAt: time.Now(),
		Metadata: ClumpMetadata{
			RowCount:      len(t.HotHeap.Rows),
			SchemaVersion: t.Schema.Version,
			CreatedAt:     t.HotHeap.CreatedAt,
		},
	}

	t.SealedClumps = append(t.SealedClumps, clump)
	if t.db != nil {
		t.db.persistClump(t.Name, clump)
	}
	t.HotHeap = NewHotHeap(t.HotHeap.MaxRows)
}
