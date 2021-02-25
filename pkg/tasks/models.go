package tasks

// Task model struct
type Task struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

type responseResult struct {
	Task  Task   `json:"task,omitempty"`
	Error string `json:"error,omitempty"`
}
