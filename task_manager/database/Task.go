package database

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

var TaskList []Task

func init() {
	t1 := Task{ID: 1, Title: "Learn Go", Description: "Study basics of Go", Status: "To Do"}
	t2 := Task{ID: 2, Title: "Build API", Description: "Make CRUD API in Go", Status: "In Progress"}
	t3 := Task{ID: 3, Title: "Test App", Description: "Test CRUD endpoints", Status: "Completed"}
	TaskList = append(TaskList, t1, t2, t3)
}
