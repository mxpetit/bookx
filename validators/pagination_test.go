package validators

import (
	"errors"
	"github.com/franela/goblin"
	"net/http"
	"reflect"
	"testing"
)

func TestPaginationValidator(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Validators > pagination", func() {
		g.Describe("validate", func() {
			g.It("shouldn't return an error since the offset is optionnal", func() {
				parameters := map[string]string{}
				validator := New(parameters)
				validator.AddRules(Pagination{})
				response := validator.Validate()

				g.Assert(response.Code == http.StatusOK).IsTrue()
				g.Assert(response.Data == nil).IsTrue()
			})

			g.It("should return an error since the offset is not parsable to an int", func() {
				parameters := map[string]string{
					"offset": "bad_parameter",
				}
				validator := New(parameters)
				validator.AddRules(Pagination{})
				response := validator.Validate()
				messsages, ok := response.Data["messages"].([]string)

				if !ok {
					err := errors.New("Cannot get response's messages. Failing...")
					g.Fail(err)
				}

				g.Assert(response.Code == http.StatusBadRequest).IsTrue()
				g.Assert(response.Data["length"] == 1).IsTrue()
				g.Assert(len(messsages) == 1).IsTrue()
				g.Assert(messsages[0] == "offset_not_number").IsTrue()
			})

			g.It("should return an error since the offset is not accepted", func() {
				parameters := map[string]string{
					"offset": "0",
				}
				validator := New(parameters)
				validator.AddRules(Pagination{})
				response := validator.Validate()
				messsages, ok := response.Data["messages"].([]string)

				if !ok {
					err := errors.New("Cannot get response's messages. Failing...")
					g.Fail(err)
				}

				g.Assert(response.Code == http.StatusBadRequest).IsTrue()
				g.Assert(response.Data["length"] == 1).IsTrue()
				g.Assert(len(messsages) == 1).IsTrue()
				g.Assert(messsages[0] == "offset_invalid").IsTrue()
			})

			g.It("shouldn't return an error since the offset is valid", func() {
				parameters := map[string]string{
					"offset": "10",
				}
				validator := New(parameters)
				validator.AddRules(Pagination{})
				response := validator.Validate()

				g.Assert(response.Code == http.StatusOK).IsTrue()
				g.Assert(response.Data == nil).IsTrue()
			})
		})

		g.Describe("isOffsetAccepted", func() {
			g.It("should return true since the offsets are accepted", func() {
				var results []bool
				wanted := []bool{true, true, true, true}

				for i := 0; i < len(acceptedOffsets); i++ {
					results = append(results, isOffsetAccepted(acceptedOffsets[i]))
				}

				g.Assert(reflect.DeepEqual(results, wanted)).IsTrue()
			})

			g.It("should return false since the offset is not accepted", func() {
				result := isOffsetAccepted(0)

				g.Assert(result).IsFalse()
			})
		})
	})
}
