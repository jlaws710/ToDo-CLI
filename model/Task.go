package model

type Task struct {
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func New(description string) Task {
	return Task{Description: description}
}

func (t *Task) MarkDone() {
	t.Completed = true
}
