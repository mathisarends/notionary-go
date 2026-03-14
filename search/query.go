package search

import "strings"

type SortDirection string
type SortTimestamp string
type QueryObjectType string

const (
	SortAscending  SortDirection = "ascending"
	SortDescending SortDirection = "descending"

	SortByLastEdited SortTimestamp = "last_edited_time"
	SortByCreated    SortTimestamp = "created_time"

	QueryObjectPage       QueryObjectType = "page"
	QueryObjectDataSource QueryObjectType = "database"
)

type WorkspaceQueryConfig struct {
	Query               *string
	ObjectType          *QueryObjectType
	SortDirection       SortDirection
	SortTimestamp       SortTimestamp
	PageSize            int
	StartCursor         *string
	TotalResultsLimit   *int
}

func DefaultQueryConfig() WorkspaceQueryConfig {
	return WorkspaceQueryConfig{
		SortDirection: SortDescending,
		SortTimestamp: SortByLastEdited,
		PageSize:      100,
	}
}

func (c WorkspaceQueryConfig) ToAPIParams() map[string]any {
	params := map[string]any{}

	if c.Query != nil && strings.TrimSpace(*c.Query) != "" {
		params["query"] = *c.Query
	}

	if c.ObjectType != nil {
		params["filter"] = map[string]any{
			"property": "object",
			"value":    string(*c.ObjectType),
		}
	}

	params["sort"] = map[string]any{
		"direction": string(c.SortDirection),
		"timestamp": string(c.SortTimestamp),
	}

	pageSize := c.PageSize
	if pageSize < 1 {
		pageSize = 1
	} else if pageSize > 100 {
		pageSize = 100
	}
	params["page_size"] = pageSize

	if c.StartCursor != nil {
		params["start_cursor"] = *c.StartCursor
	}

	return params
}

func (c WorkspaceQueryConfig) WithPagesOnly() WorkspaceQueryConfig {
	t := QueryObjectPage
	c.ObjectType = &t
	return c
}

func (c WorkspaceQueryConfig) WithStartCursor(cursor string) WorkspaceQueryConfig {
	if cursor == "" {
		c.StartCursor = nil
		return c
	}
	c.StartCursor = &cursor
	return c
}