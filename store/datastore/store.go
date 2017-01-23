package datastore

import (
	"errors"
	"github.com/gocql/gocql"
	"github.com/mxpetit/bookx/store"
	"log"
	"os"
)

// datastore represents a session.
type datastore struct {
	*gocql.Session
}

// New create a connection to the database from environnement variable and
// returns a Store.
func New() store.Store {
	session, err := open(os.Getenv("BOOKX_IP"), os.Getenv("BOOKX_KEYSPACE"))

	if err != nil {
		log.Fatal(err)
	}

	return From(session)
}

// From returns a Store based on the session provided.
func From(session *gocql.Session) store.Store {
	return &datastore{session}
}

// open opens a new database connection with the specified parameters and
// returns a session.
func open(ip, keyspace string) (*gocql.Session, error) {
	if keyspace == "" || ip == "" {
		return nil,
			errors.New("Keyspace or IP missing. Unable to open database connection")
	}

	cluster := gocql.NewCluster(ip)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.LocalOne

	return cluster.CreateSession()
}

// openTest opens a new database connection in order to perform tests.
func openTest() (*gocql.Session, error) {
	var (
		keyspace = "bookx_test"
		ip       = "127.0.0.1"
	)

	return open(ip, keyspace)
}
