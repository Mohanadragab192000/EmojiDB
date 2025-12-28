package core

import (
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
