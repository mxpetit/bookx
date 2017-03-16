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

// New returns a Store given environnement variable.
func New() store.Store {
	session, err := open(os.Getenv("BOOKX_IP"), os.Getenv("BOOKX_KEYSPACE"))

	if err != nil {
		log.Fatal(err)
	}

	return From(session)
}

// From returns a Store given a session.
func From(session *gocql.Session) store.Store {
	return &datastore{session}
}

// open returns a session given IP and keyspace.
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
