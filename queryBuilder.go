package deta

import "strings"

type QueryBuilder interface {
	BuildSelect(cols []string, distinct bool) string
	BuildFrom(table string) string
	BuildGroupBy(cols []string) string
	BuildOrderByAndLimit(string, []string, int64, int64) string
}

type BaseQueryBuilder struct {
	db *DB
}

func NewQueryBuilder(db *DB) *BaseQueryBuilder {
	return &BaseQueryBuilder{db}
}

func (qb *BaseQueryBuilder) BuildSelect(cols []string, distinct bool, table string) string {
	query := "SELECT "

	if distinct {
		query += "DISTINCT "
	}

	for _, col := range cols {
		column := []string{table, ".", col}
		query += strings.Join(column, ", ")
	}

	return query
}

func (qb *BaseQueryBuilder) BuildFrom(table string) string {
	query := []string{"FROM", table, ""}
	return strings.Join(query, " ")
}

func (qb *BaseQueryBuilder) BuildGroupBy(cols []string) string {
	return ""
}

func (qb *BaseQueryBuilder) BuildOrderByAndLimit(string, []string, int64, int64) string {
	return ""
}
