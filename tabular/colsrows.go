package tabular

import "strconv"

// Columns is a simple typedef
type Columns []string

// Indexer typedef for the Index func
type Indexer func(colName string, row Row) string

// Index of the Columns
func (cols Columns) Index() Indexer {
	index := make(map[string]int)
	for i, c := range cols {
		index[c] = i
	}
	return func(col string, row Row) string {
		pos := index[col]
		return row[pos]
	}
}

// Row typedef
type Row []string

// Map parsed the rows into maps
func (r Row) Map(cols Columns, indexer Indexer) map[string]string {
	tmp := make(map[string]string)
	for _, col := range cols {
		val := indexer(col, r)
		tmp[col] = val
	}
	return tmp
}

// Compare provides an override for the basic comparison
type Compare func(this, that string) bool

// CompareAsInts ...
func CompareAsInts(this, that string) bool {
	thisInt, err := strconv.ParseInt(this, 10, 64)
	if err != nil {
		return false
	}
	thatInt, err := strconv.ParseInt(that, 10, 64)
	if err != nil {
		return false
	}
	return thisInt < thatInt
}

// SortRows sorts rows by the given field
type SortRows struct {
	Row     []Row
	SortBy  string
	Indexer Indexer
	Compare Compare
}

func (a SortRows) Len() int      { return len(a.Row) }
func (a SortRows) Swap(i, j int) { a.Row[i], a.Row[j] = a.Row[j], a.Row[i] }
func (a SortRows) Less(i, j int) bool {
	if a.Compare == nil {
		return a.Indexer(a.SortBy, a.Row[i]) < a.Indexer(a.SortBy, a.Row[j])
	}
	return a.Compare(a.Indexer(a.SortBy, a.Row[i]), a.Indexer(a.SortBy, a.Row[j]))
}
