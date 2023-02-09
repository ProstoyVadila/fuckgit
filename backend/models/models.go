package models

type Content struct {
	Name string `json:"name"`
	Body string `json:"body"`
}

type Solution struct {
	URL string `json:"url"`
	ID  int    `json:"id"`
}

type Question struct {
	Left     *Question `json:"left"`
	Right    *Question `json:"right"`
	Content  *Content  `json:"content"`
	Solution *Solution `json:"solution"`
	ID       int       `json:"id"`
}

type Questions struct {
	Root   *Question `json:"root"`
	LastID int       `json:"last_id"`
}

func NewSolution(url string, id int) *Solution {
	return &Solution{
		URL: url,
		ID:  id,
	}
}

func NewContent(name, body string) *Content {
	return &Content{
		Name: name,
		Body: body,
	}
}

func NewQuestion(content *Content, id int, solutions ...*Solution) *Question {
	var solution *Solution
	if len(solutions) > 0 {
		solution = solutions[0]
	}
	return &Question{
		Content:  content,
		ID:       id,
		Solution: solution,
	}
}

func NewQuestions() *Questions {
	return &Questions{
		LastID: 0,
	}
}

func (qs *Questions) Insert(content *Content, isRight bool, solutions ...*Solution) {
	if qs.Root == nil {
		qs.Root = &Question{
			Content: content,
			ID:      qs.LastID,
		}
		qs.LastID++
		return
	}

	qs.Root.insert(content, qs.LastID, isRight, solutions...)
	qs.LastID++
}

func (q *Question) insert(content *Content, id int, isRight bool, solutions ...*Solution) {
	if isRight {
		q.Right = NewQuestion(content, id, solutions...)
	} else {
		q.Left = NewQuestion(content, id, solutions...)
	}
}
