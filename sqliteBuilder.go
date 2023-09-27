package deta

type SqliteBuilder struct {
	*BaseBuilder
}

func NewSqliteBuilder(db *DB, executor Executor) Builder {
	return &SqliteBuilder{
		NewBaseBuilder(db, executor),
	}
}

func (sqlite *SqliteBuilder) Select(table string, cols ...string) *SelectQuery {
	return NewSelect(sqlite, sqlite.db).Select(cols...).From(table)
}
