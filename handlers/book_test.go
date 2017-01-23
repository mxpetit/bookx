package handlers

import (
	"github.com/franela/goblin"
	"reflect"
	"testing"
)

func TestBook(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Book handler", func() {
		g.It("should convert string parameters to int, or to default int value on error", func() {
			defaultValue := -1
			testCases := []string{"1", "-8", "a", "4294967295"}
			results := make([]int, len(testCases))
			wanted := []int{1, -8, defaultValue, 4294967295}

			for i, v := range testCases {
				results[i] = convertParameter(v, defaultValue)
			}

			g.Assert(reflect.DeepEqual(results, wanted)).IsTrue()
		})
	})
}
