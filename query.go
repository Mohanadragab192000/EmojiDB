package emojidb

import (
	"errors"
)

type Query struct {
	db        *Database
	tableName string
	filters   []FilterFunc
	columns   []string
}

type FilterFunc func(Row) bool

func (db *Database) Query(tableName string) *Query {
	return &Query{
		db:        db,
		tableName: tableName,
	}
}

func (q *Query) Filter(f FilterFunc) *Query {
	q.filters = append(q.filters, f)
	return q
}

func (q *Query) Select(columns ...string) *Query {
	q.columns = columns
	return q
}

func (q *Query) Execute() ([]Row, error) {
	q.db.mu.RLock()
	table, ok := q.db.tables[q.tableName]
	q.db.mu.RUnlock()

	if !ok {
		return nil, errors.New("table not found: " + q.tableName)
	}

	var results []Row

	// 1. Scan Hot Heap
	table.mu.RLock()
	for _, row := range table.HotHeap.Rows {
		if q.matches(row) {
			results = append(results, q.project(row))
		}
	}

	// 2. Scan Sealed Clumps (In-memory)
	for _, clump := range table.SealedClumps {
		for _, row := range clump.Rows {
			if q.matches(row) {
				results = append(results, q.project(row))
			}
		}
	}
	table.mu.RUnlock()

	return results, nil
}

func (q *Query) matches(row Row) bool {
	for _, filter := range q.filters {
		if !filter(row) {
			return false
		}
	}
	return true
}

func (q *Query) project(row Row) Row {
	if len(q.columns) == 0 {
		return row
	}

	projected := make(Row)
	for _, col := range q.columns {
		if val, ok := row[col]; ok {
			projected[col] = val
		}
	}
	return projected
}
