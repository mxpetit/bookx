package datastore

import (
	"github.com/franela/goblin"
	"github.com/gocql/gocql"
	"testing"
)

func TestBook(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping store/datastore tests.")
	}

	g := goblin.Goblin(t)
	session, err := openTest()
	defer session.Close()

	g.Assert(err == nil).IsTrue()
	g.Assert(session == nil).IsFalse()

	store := From(session)

	g.Describe("package datastore > book", func() {
		g.Describe("function GetAllBooks", func() {
			g.BeforeEach(func() {
				session.Query("TRUNCATE book").Exec()
			})

			g.It("should return an error since there is no book in the store (ErrNoBooksAvailable)", func() {
				results, err := store.GetAllBooks("10")

				g.Assert(err == ErrNoBooksAvailable).IsTrue()
				g.Assert(len(results) == 0)
			})

			g.It("should get an array of 2 books", func() {
				_, err1 := store.CreateBook("Foo", 250)
				_, err2 := store.CreateBook("Bar", 250)
				results, err3 := store.GetAllBooks("10")

				g.Assert(err1 == nil).IsTrue()
				g.Assert(err2 == nil).IsTrue()
				g.Assert(err3 == nil).IsTrue()
				g.Assert(len(results) == 2)
			})

			g.It("should get an array of 1 books since no limit was provided (DEFAULT_MIN_LIMIT)", func() {
				_, err1 := store.CreateBook("Foo", 250)
				_, err2 := store.CreateBook("Bar", 250)
				results, err3 := store.GetAllBooks("")

				g.Assert(err1 == nil).IsTrue()
				g.Assert(err2 == nil).IsTrue()
				g.Assert(err3 == nil).IsTrue()
				g.Assert(len(results) == 1)
			})
		})

		g.Describe("function GetBook", func() {
			g.BeforeEach(func() {
				session.Query("TRUNCATE book").Exec()
			})

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

			g.It("should return an error since there is no book in the store (ErrBookDoesNotExists)", func() {
				// Refering to https://tools.ietf.org/html/rfc4122#page-4
				_, err2 := store.GetBook("f81d4fae-7dec-11d0-a765-00a0c91e6bf6")

				g.Assert(err2 == ErrBookDoesNotExists).IsTrue()
			})

			g.It("should return an error since the UUID is invalid (ErrUUIDInvalid)", func() {
				_, err1 := store.CreateBook("Foo", 250)
				_, err2 := store.GetBook("invalid_uuid")

				g.Assert(err1 == nil).IsTrue()
				g.Assert(err2 == ErrUUIDInvalid).IsTrue()
			})
		})

		g.Describe("function scanBooks", func() {
			g.BeforeEach(func() {
				session.Query("TRUNCATE book").Exec()
			})

			g.It("should get an array of 2 books", func() {
				_, err1 := store.CreateBook("Foo", 250)
				_, err2 := store.CreateBook("Bar", 250)
				iter := session.Query(getAllBooks, 2).Iter()
				results, err3 := scanBooks(iter)

				g.Assert(err1 == nil).IsTrue()
				g.Assert(err2 == nil).IsTrue()
				g.Assert(err3 == nil).IsTrue()
				g.Assert(len(results) == 2)
			})

			g.It("should get an error since iter is nil (ErrUnableToRetrievesBooks)", func() {
				results, err := scanBooks(nil)

				g.Assert(err == ErrUnableToRetrievesBooks).IsTrue()
				g.Assert(len(results) == 0)
			})

			g.It("should get an error since there is no book in the store (ErrNoBooksAvailable)", func() {
				iter := session.Query(getAllBooks, 10).Iter()
				results, err := scanBooks(iter)

				g.Assert(err == ErrNoBooksAvailable).IsTrue()
				g.Assert(len(results) == 0)
			})
		})
	})
}
