package tabular

import (
	"sort"
	"testing"

	"github.com/jpfielding/gotest/testutils"
)

func TestColumns(t *testing.T) {
	rows := []Row{
		Row([]string{"one", "active", "234234"}),
		Row([]string{"two", "sold", "543234"}),
		Row([]string{"three", "pending", "9834543"}),
	}

	cols := Columns([]string{"id", "status", "price"})

	index := cols.Index()
	map1 := rows[0].Map(cols, index)
	testutils.Equals(t, "one", map1["id"])
	testutils.Equals(t, "active", map1["status"])
	testutils.Equals(t, "234234", map1["price"])

	map3 := rows[2].Map(cols, index)
	testutils.Equals(t, "three", map3["id"])
	testutils.Equals(t, "pending", map3["status"])
	testutils.Equals(t, "9834543", map3["price"])

}

func TestRowSort(t *testing.T) {
	rows := []Row{
		Row([]string{"one", "active", "9234234"}),
		Row([]string{"two", "sold", "7543234"}),
		Row([]string{"three", "pending", "334543"}),
	}

	cols := Columns([]string{"id", "status", "price"})
	index := cols.Index()

	sortBy := SortRows{
		Compare: CompareAsInts,
		Row:     rows,
		SortBy:  "price",
		Indexer: index,
	}
	sort.Sort(sortBy)
	testutils.Equals(t, "three", rows[0].Map(cols, index)["id"])
	testutils.Equals(t, "two", rows[1].Map(cols, index)["id"])
	testutils.Equals(t, "one", rows[2].Map(cols, index)["id"])

	sortBy.Compare = nil
	sortBy.SortBy = "id"
	sort.Sort(sortBy)
	testutils.Equals(t, "one", rows[0].Map(cols, index)["id"])
	testutils.Equals(t, "three", rows[1].Map(cols, index)["id"])
	testutils.Equals(t, "two", rows[2].Map(cols, index)["id"])
}
