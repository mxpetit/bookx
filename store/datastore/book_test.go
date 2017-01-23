package datastore

import (
	"github.com/franela/goblin"
	"github.com/gocql/gocql"
	"github.com/mxpetit/bookx/model"
	"testing"
)

func TestBook(t *testing.T) {
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

		g.It("shoud get a list", func() {
			_, err1 := store.CreateBook("Foo", 250)
			_, err2 := store.CreateBook("Bar", 250)
			results, err3 := store.GetAllBooks(gocql.UUID{}, 2)

			g.Assert(err1 == nil).IsTrue()
			g.Assert(err2 == nil).IsTrue()
			g.Assert(err3 == nil).IsTrue()
			g.Assert(results.Length == 2)

			book1 := results.Results[0].(model.Book)
			book2 := results.Results[1].(model.Book)

			g.Assert(book1.Title != "").IsTrue()
			g.Assert(book1.NumberOfPages != 0).IsTrue()
			g.Assert(book2.Title != "").IsTrue()
			g.Assert(book2.NumberOfPages != 0).IsTrue()
		})

		g.It("should get one element from the list", func() {
			_, err1 := store.CreateBook("Foo", 250)
			_, err2 := store.CreateBook("Bar", 250)
			_, err3 := store.CreateBook("Baz", 250)
			_, err4 := store.CreateBook("Xyzzy", 250)
			results, err5 := store.GetAllBooks(gocql.UUID{}, 1)

			g.Assert(err1 == nil).IsTrue()
			g.Assert(err2 == nil).IsTrue()
			g.Assert(err3 == nil).IsTrue()
			g.Assert(err4 == nil).IsTrue()
			g.Assert(err5 == nil).IsTrue()
			g.Assert(results.Length == 1).IsTrue()

			book := results.Results[0].(model.Book)

			g.Assert(book.NumberOfPages != 0).IsTrue()
			g.Assert(book.Title != "").IsTrue()
		})

		g.It("shouldn't get a list because end is lower than 1", func() {
			results, err := store.GetAllBooks(gocql.UUID{}, 0)

			g.Assert(err != nil).IsTrue()
			g.Assert(results.Length == 0).IsTrue()
		})

		g.It("should get by ID", func() {
			id, err1 := store.CreateBook("Foo", 250)
			books, err2 := store.GetBook(id)

			g.Assert(err1 == nil).IsTrue()
			g.Assert(err2 == nil).IsTrue()
			g.Assert(books.Title == "Foo").IsTrue()
			g.Assert(books.NumberOfPages == 250).IsTrue()
		})
	})
}
