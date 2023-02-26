package neo4j_db

import (
	"github.com/ProstoyVadila/fuckgit/config"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rs/zerolog/log"
)

type Neo4jDB struct {
	driver neo4j.DriverWithContext
}

func New(config config.Config) *Neo4jDB {
	driver, err := neo4j.NewDriverWithContext(config.Neo4jURL, neo4j.BasicAuth(config.Neo4jUser, config.Neo4jPassword, ""))
	if err != nil {
		log.Fatal().Err(err).Msg("Neo4j driver error")
	}
	return &Neo4jDB{driver: driver}
}
