package datastore

import (
	"github.com/franela/goblin"
	"github.com/gocql/gocql"
	"log"
	"testing"
)

func TestShelve(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping store/datastore tests.")
	}

	g := goblin.Goblin(t)
	session, err := openTest()
	defer session.Close()

	g.Assert(err == nil).IsTrue()
	g.Assert(session == nil).IsFalse()

	store := From(session)

	g.Describe("package datastore > shelve", func() {
		g.Describe("function CreateShelve", func() {
			g.BeforeEach(func() {
				session.Query("TRUNCATE shelve").Exec()
			})

			g.It("should create an empty shelve", func() {
				uuid, err := store.CreateShelve("shelve1")
				log.Println("err", err)
				g.Assert(err == nil).IsTrue()

				shelve, err := store.GetShelve(uuid)
				g.Assert(err == nil).IsTrue()

				parsedUUID, err := gocql.ParseUUID(uuid)
				g.Assert(err == nil).IsTrue()

				g.Assert(shelve.Name == "shelve1").IsTrue()
				g.Assert(shelve.Id == parsedUUID).IsTrue()
			})
		})
	})
}
