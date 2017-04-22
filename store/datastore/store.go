package datastore

import (
	"github.com/gocql/gocql"
	"github.com/mxpetit/bookx/store"
	"log"
	"os"
)

const (
	KEYSPACE            = "bookx"
	TEST_KEYSPACE       = "bookx_test"
	DEFAULT_DATABASE_IP = "127.0.0.1"
)

// datastore represents a database session.
type datastore struct {
	*gocql.Session
}

// New returns a Store given database's IP adress.
func New() store.Store {
	databaseIp := os.Getenv("BOOKX_DATABASE_IP")
	session, err := open(databaseIp, KEYSPACE)

	if err != nil {
		log.Fatal(err)
	}

	return From(session)
}

// From returns a Store given a database session.
func From(session *gocql.Session) store.Store {
	return &datastore{session}
}

// open returns a database session given IP and keyspace.
func open(ip, keyspace string) (*gocql.Session, error) {
	if keyspace == "" {
		keyspace = KEYSPACE
	}

	if ip == "" {
		ip = DEFAULT_DATABASE_IP
	}

	cluster := gocql.NewCluster(ip)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.LocalOne

	return cluster.CreateSession()
}

// openTest opens a new database connection to perform tests.
func openTest() (*gocql.Session, error) {
	var (
		keyspace = TEST_KEYSPACE
		ip       = DEFAULT_DATABASE_IP
	)

	return open(ip, keyspace)
}
