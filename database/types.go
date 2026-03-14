package database

type Database struct {
	ID             string `json:"id"`
	Object         string `json:"object"`
	CreatedTime    string `json:"created_time"`
	LastEditedTime string `json:"last_edited_time"`
	Title          []struct {
		PlainText string `json:"plain_text"`
	} `json:"title"`
	Properties map[string]interface{} `json:"properties"`
}