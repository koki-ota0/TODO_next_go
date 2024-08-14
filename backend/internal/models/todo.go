package models

type Todo struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Status   string `json:"status"`
	Deadline string `json:"deadline"` // 追加
}
