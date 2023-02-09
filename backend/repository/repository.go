package repository

import (
	"github.com/ProstoyVadila/fuckgit/models"
)

const (
	Left  = false
	Right = true
)

type Store interface {
	Questions() (*models.Questions, error)
}

type Mock struct{}

func New() Store {
	return &Mock{}
}

func (m *Mock) Questions() (*models.Questions, error) {
	return getQuestions(), nil
}

func getQuestions() *models.Questions {
	questions := models.NewQuestions()
	questions.Insert(
		models.NewContent("start", "Надо откатить изменения?"),
		Left,
	)
	questions.Insert(
		models.NewContent("start_right", "Изменения уже закоммичены?"),
		Right,
	)
	questions.Insert(
		models.NewContent("local", "Только локально?"),
		Right,
	)
	return questions
}
