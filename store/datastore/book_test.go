package datastore

import (
	"github.com/franela/goblin"
	"github.com/gocql/gocql"
	"testing"
)

func TestBook(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipped")
	}

	g := goblin.Goblin(t)
	session, err := openTest()

	g.Assert(err == nil).IsTrue()
	g.Assert(session == nil).IsFalse()

	defer session.Close()
	store := From(session)

	g.Describe("Book datastore", func() {
		g.BeforeEach(func() {
			session.Query("TRUNCATE book").Exec()
		})

		g.Describe("GetAllBooks", func() {
			g.It("should get a list of size 2", func() {
				_, err1 := store.CreateBook("Foo", 250)
				_, err2 := store.CreateBook("Bar", 250)
				results, err3 := store.GetAllBooks("", "10")

				g.Assert(err1 == nil).IsTrue()
				g.Assert(err2 == nil).IsTrue()
				g.Assert(err3 == nil).IsTrue()
				g.Assert(len(results) == 2)

				for i := 0; i < len(results); i++ {
					g.Assert(results[i].Id != gocql.UUID{}).IsTrue()
					g.Assert(results[i].NumberOfPages != 0).IsTrue()
					g.Assert(results[i].Title != "").IsTrue()
				}
			})

			g.It("should return an error (no_books_available)", func() {
				results, err := store.GetAllBooks("", "10")

				g.Assert(err == ErrNoBooksAvailable).IsTrue()
				g.Assert(len(results) == 0)
			})
		})

		g.Describe("GetBook", func() {
			g.It("should get one book", func() {
				id, err1 := store.CreateBook("Foo", 250)
				parsedId, err2 := gocql.ParseUUID(id)
				book, err3 := store.GetBook(id)

				g.Assert(err1 == nil).IsTrue()
				g.Assert(err2 == nil).IsTrue()
				g.Assert(err3 == nil).IsTrue()
				g.Assert(book.Title == "Foo").IsTrue()
				g.Assert(book.NumberOfPages == 250).IsTrue()
				g.Assert(book.Id == parsedId).IsTrue()
			})

			g.It("should return an error (ErrBookDoesNotExists)", func() {
				// Inexistant UUID, refer to https://tools.ietf.org/html/rfc4122#page-4
				_, err2 := store.GetBook("f81d4fae-7dec-11d0-a765-00a0c91e6bf6")

				g.Assert(err2 == ErrBookDoesNotExists).IsTrue()
			})
		})
	})
}
