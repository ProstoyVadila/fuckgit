package neo4j_db

import (
	"context"

	"github.com/ProstoyVadila/fuckgit/models"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rs/zerolog/log"
)

func (db *Neo4jDB) Questions(ctx context.Context) (*models.Questions, error) {
	session := db.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	log.Info().Msg("getting question from db")
	query := "match (q:Question {name: 'test'}) return q.name as name, q.body as body"
	res, err := session.Run(ctx, query, nil)
	if err != nil {
		log.Err(err).Msg("error getting question from db")
		return nil, err
	}
	record, err := res.Single(ctx)
	if err != nil {
		return nil, err
	}

	name, ok := record.Get("name")
	if !ok {
		log.Err(err).Msg("there is no name in question")
		return nil, res.Err()
	}
	body, ok := record.Get("body")
	if !ok {
		log.Err(err).Msg("there is no body in question")
		return nil, res.Err()
	}
	questions := models.NewQuestions()
	questions.Root = &models.Question{
		Name: name.(string),
		Body: body.(string),
	}
	return questions, nil
}
