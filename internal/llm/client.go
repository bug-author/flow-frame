package llm

type OllamaClient struct {
	Model string
}

type Task struct {
	Title     string
	Content   string
	StartDate string
	EndDate   string
}
