package scylla

import (
	"github.com/gocql/gocql"
	"time"
)

type DBConnection struct {
	consistency gocql.Consistency
	keyspace    string
	hosts       []string
}

func (conn *DBConnection) createCluster() *gocql.ClusterConfig {
	retryPolicy := &gocql.ExponentialBackoffRetryPolicy{
		Min:        time.Second,
		Max:        10 * time.Second,
		NumRetries: 5,
	}

	cluster := gocql.NewCluster(conn.hosts...)
	cluster.Consistency = conn.consistency
	cluster.Keyspace = conn.keyspace
	cluster.Timeout = 5 * time.Second
	cluster.RetryPolicy = retryPolicy
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())

	return cluster
}

func (conn *DBConnection) createSession(cluster *gocql.ClusterConfig) (*gocql.Session, error) {
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	return session, nil
}

func NewScyllaDBConnection(consistency gocql.Consistency, keyspace string, hosts ...string) *DBConnection {
	return &DBConnection{
		consistency,
		keyspace,
		hosts,
	}
}

func (conn *DBConnection) GetConnection() (*gocql.Session, error) {
	cluster := conn.createCluster()
	return conn.createSession(cluster)
}
