package initialize

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

func setupDatabase(session *gocql.Session, keyspace string) {
	err := session.Query(fmt.Sprintf(`CREATE KEYSPACE IF NOT EXISTS %v WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}`, keyspace)).Exec()
	if err != nil {
		log.Fatalf("failed to create keyspace: %v", err)
	}

	err = session.Query(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %v.store (key text PRIMARY KEY, value text)`, keyspace)).Exec()
	if err != nil {
		log.Fatalf("failed to create table: %v", err)
	}
}

func InitCassandraDB(keyspace string) *gocql.Session {
	cluster := gocql.NewCluster("localhost:9042")
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalln("unable to connect to cassandra db:", err)
	}

	setupDatabase(session, keyspace)
	return session
}
