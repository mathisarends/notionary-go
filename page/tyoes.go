package page

type Page struct {
	ID             string                 `json:"id"`
	Object         string                 `json:"object"`
	CreatedTime    string                 `json:"created_time"`
	LastEditedTime string                 `json:"last_edited_time"`
	Archived       bool                   `json:"archived"`
	Properties     map[string]interface{} `json:"properties"`
}