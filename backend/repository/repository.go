package repository

import (
	"errors"

	"github.com/ProstoyVadila/fuckgit/models"
)

type Store interface {
	Questions() (*models.Questions, error)
}

type Mock struct{}

func New() Store {
	return &Mock{}
}

func (m *Mock) Questions() (*models.Questions, error) {
	return nil, errors.New("questions not found")
}
