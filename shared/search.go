package shared

type SearchRequest struct {
	Query       string        `json:"query"`
	Filter      *SearchFilter `json:"filter,omitempty"`
	StartCursor string        `json:"start_cursor,omitempty"`
}

type SearchFilter struct {
	Value    string `json:"value"`
	Property string `json:"property"`
}

type SearchResponse struct {
	Results    []SearchResult `json:"results"`
	HasMore    bool           `json:"has_more"`
	NextCursor string         `json:"next_cursor"`
}

type RichText struct {
	PlainText string `json:"plain_text"`
}

type SearchResult struct {
	ID         string                   `json:"id"`
	Object     string                   `json:"object"`
	Title      []RichText               `json:"title"`
	Properties map[string]PropertyValue `json:"properties"`
}

type PropertyValue struct {
	Title []RichText `json:"title"`
}

func (r SearchResult) PlainTitle() string {
	for _, t := range r.Title {
		if t.PlainText != "" {
			return t.PlainText
		}
	}
	for _, prop := range r.Properties {
		for _, t := range prop.Title {
			if t.PlainText != "" {
				return t.PlainText
			}
		}
	}
	return ""
}