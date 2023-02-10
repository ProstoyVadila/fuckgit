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
	return mockQuestions(), nil
}

func mockQuestions() *models.Questions {
	questions := models.NewQuestions()
	questions.Root = &models.Question{
		Name: "start",
		Body: "Нужно откатить изменения?",
		Left: &models.Question{
			Name: "unknown",
			Body: "Тогда не знаю, чем помочь, друг",
			Solution: &models.Solution{
				URL: "/",
			},
			Right: &models.Question{
				Name: "locally",
				Body: "Нужно откатить изменения локально?",
				Left: &models.Question{
					Name: "remotely",
					Body: "Нужно откатить изменения в удаленном репозитории???",
					Left: &models.Question{
						Name: "unknown",
						Body: "Тогда не знаю, чем помочь, друг",
						Solution: &models.Solution{
							URL: "/",
						},
					},
					Right: &models.Question{
						Name: "remotely_true",
						Solution: &models.Solution{
							URL:         "/solution",
							Command:     "git push --force blabla",
							Description: "Откатить изменения в удаленном репозитории",
						},
					},
				},
				Right: &models.Question{
					Name: "locally_true",
					Solution: &models.Solution{
						URL:         "/solution",
						Command:     "git reset --hard HEAD~1",
						Description: "Откатить изменения локально, cтерев их из истории",
					},
				},
			},
		},
	}
	return questions
}
