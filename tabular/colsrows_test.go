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

	test := func(row int, id, status, price string) func(*testing.T) {
		return func(tt *testing.T) {
			m := rows[row].Map(cols, index)
			testutils.Equals(tt, id, m["id"])
			testutils.Equals(tt, status, m["status"])
			testutils.Equals(tt, price, m["price"])
		}
	}
	t.Run("cols", test(0, "one", "active", "234234"))
	t.Run("cols", test(1, "two", "sold", "543234"))
	t.Run("cols", test(2, "three", "pending", "9834543"))
}

func TestRowSort(t *testing.T) {
	rows := []Row{
		Row([]string{"one", "active", "9234234"}),
		Row([]string{"two", "sold", "7543234"}),
		Row([]string{"three", "pending", "334543"}),
	}

	cols := Columns([]string{"id", "status", "price"})
	index := cols.Index()

	test := func(row int, id, status, price string) func(*testing.T) {
		return func(tt *testing.T) {
			m := rows[row].Map(cols, index)
			testutils.Equals(tt, id, m["id"])
			testutils.Equals(tt, status, m["status"])
			testutils.Equals(tt, price, m["price"])
		}
	}

	sortBy := SortRows{
		Row:     rows,
		Indexer: index,
		Compare: CompareAsInts,
		SortBy:  "price",
	}

	sortBy.Compare = CompareAsInts
	sortBy.SortBy = "price"
	sort.Sort(sortBy)

	t.Run("by price", test(0, "three", "pending", "334543"))
	t.Run("by price", test(1, "two", "sold", "7543234"))
	t.Run("by price", test(2, "one", "active", "9234234"))

	sortBy.Compare = nil
	sortBy.SortBy = "id"
	sort.Sort(sortBy)

	t.Run("by id", test(1, "three", "pending", "334543"))
	t.Run("by id", test(2, "two", "sold", "7543234"))
	t.Run("by id", test(0, "one", "active", "9234234"))
}
