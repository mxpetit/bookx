package renderer

import (
	"github.com/franela/goblin"
	"reflect"
	"testing"
)

func TestTranslations(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping translations...")
	}

	g := goblin.Goblin(t)

	g.Describe("Renderer > translations", func() {
		g.Describe("translateArray", func() {
			g.It("should translate an array", func() {
				strings := []string{
					"book_created",
					"book_doesnt_exists",
				}

				wanted := []string{
					"Le livre a été crée.",
					"Le livre n'existe pas.",
				}

				translateFunction := getTranslateFunction("fr-FR")

				result := translateArray(strings, translateFunction)
				g.Assert(reflect.DeepEqual(result, wanted))
			})

			g.It("should let the array as it was since the language provided doesn't exists", func() {
				strings := []string{
					"book_created",
					"book_doesnt_exists",
				}

				translateFunction := getTranslateFunction("language_that_doesnt_exists")

				result := translateArray(strings, translateFunction)
				g.Assert(reflect.DeepEqual(result, strings)).IsTrue()
			})
		})

		g.Describe("translateResponse", func() {
			g.It("should translate the whole response", func() {
				response := map[string]interface{}{
					"toto": "book_created",
					"tata": map[string]interface{}{
						"tutu": "book_created",
						"tata": []string{
							"book_created",
							"book_created",
						},
					},
				}

				wanted := map[string]interface{}{
					"toto": "Le livre a été crée.",
					"tata": map[string]interface{}{
						"tutu": "Le livre a été crée.",
						"tata": []string{
							"Le livre a été crée.",
							"Le livre a été crée.",
						},
					},
				}

				translateFunction := getTranslateFunction("fr-FR")

				result := translateResponse(response, translateFunction)
				g.Assert(reflect.DeepEqual(result, wanted))
			})
		})
	})
}
