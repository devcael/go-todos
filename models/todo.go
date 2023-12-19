package models

type Todo struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func (todo *Todo) Check() {
	todo.Completed = true
}

func (todo *Todo) Uncheck() {
	todo.Completed = false
}
