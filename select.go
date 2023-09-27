package deta

type SelectQuery struct {
	builder Builder

	table    string
	cols     []string
	orderBy  []string
	groupBy  []string
	limit    int64
	offset   int64
	distinct bool
}

func NewSelect(builder Builder, db *DB) *SelectQuery {
	return &SelectQuery{
		builder: builder,

		table:    "",
		cols:     []string{},
		distinct: false,
		orderBy:  []string{},
		groupBy:  []string{},
		limit:    -1,
	}
}

func (s *SelectQuery) Build() {

}

func (s *SelectQuery) Unique(t interface{}) {

}

func (s *SelectQuery) All(t interface{}) {

}

func (s *SelectQuery) Select(cols ...string) *SelectQuery {
	s.cols = cols
	return s
}

func (s *SelectQuery) From(table string) *SelectQuery {
	s.table = table
	return s
}

func (s *SelectQuery) Distinct(distinct bool) *SelectQuery {
	s.distinct = distinct
	return s
}

func (s *SelectQuery) Limit(limit int64) *SelectQuery {
	s.limit = limit
	return s
}

func (s *SelectQuery) Offset(offset int64) *SelectQuery {
	s.offset = offset
	return s
}

func (s *SelectQuery) OrderBy(orderBy ...string) *SelectQuery {
	s.orderBy = orderBy
	return s
}

func (s *SelectQuery) GroupBy(groupBy ...string) *SelectQuery {
	s.groupBy = groupBy
	return s
}
