package search

type QueryBuilder struct {
	config WorkspaceQueryConfig
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{config: DefaultQueryConfig()}
}

func (b *QueryBuilder) WithQuery(query string) *QueryBuilder {
	b.config.Query = &query
	return b
}

func (b *QueryBuilder) WithPagesOnly() *QueryBuilder {
	t := QueryObjectPage
	b.config.ObjectType = &t
	return b
}

func (b *QueryBuilder) WithDataSourcesOnly() *QueryBuilder {
	t := QueryObjectDataSource
	b.config.ObjectType = &t
	return b
}

func (b *QueryBuilder) WithSortAscending() *QueryBuilder {
	b.config.SortDirection = SortAscending
	return b
}

func (b *QueryBuilder) WithSortDescending() *QueryBuilder {
	b.config.SortDirection = SortDescending
	return b
}

func (b *QueryBuilder) WithSortByCreated() *QueryBuilder {
	b.config.SortTimestamp = SortByCreated
	return b
}

func (b *QueryBuilder) WithSortByLastEdited() *QueryBuilder {
	b.config.SortTimestamp = SortByLastEdited
	return b
}

func (b *QueryBuilder) WithPageSize(size int) *QueryBuilder {
	if size > 100 {
		size = 100
	}
	b.config.PageSize = size
	return b
}

func (b *QueryBuilder) WithTotalLimit(limit int) *QueryBuilder {
	b.config.TotalResultsLimit = &limit
	return b
}

func (b *QueryBuilder) WithStartCursor(cursor string) *QueryBuilder {
	b.config.StartCursor = &cursor
	return b
}

func (b *QueryBuilder) Build() WorkspaceQueryConfig {
	return b.config
}