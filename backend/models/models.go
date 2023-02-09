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

func NewContent(name, body string) *Content {
	return &Content{
		Name: name,
		Body: body,
	}
}

func NewQuestion(content *Content, id int) *Question {
	return &Question{
		Content: content,
		ID:      id,
	}
}

func NewQuestions() *Questions {
	return &Questions{
		LastID: 0,
	}
}

func (qs *Questions) Insert(content *Content, isRight bool) {
	if qs.Root == nil {
		qs.Root = &Question{
			Content: content,
			ID:      0,
		}
		qs.LastID++
		return
	}

	qs.Root.insert(content, qs.LastID, isRight)
	qs.LastID++
}

func (q *Question) insert(content *Content, id int, isRight bool, solutions ...*Solution) {
	// TODO: add Solutions
	if isRight {
		q.Right = NewQuestion(content, id)
	} else {
		q.Left = NewQuestion(content, id)
	}
}
