package deta

import "database/sql"

type (
	BuilderFunc func(*DB, Executor) Builder
	DB          struct {
		Builder
		sqlDB  *sql.DB
		driver string
	}
)

var BuilderMapFunc = map[string]BuilderFunc{
	"sqlite": NewSqliteBuilder,
}

func newSource(sqlDB *sql.DB, driver string) *DB {
	db := &DB{
		sqlDB:  sqlDB,
		driver: driver,
	}

	db.Builder = db.newBuilder(db.sqlDB)

	return db
}

func Start(driver string, url string) (*DB, error) {
	sql, err := sql.Open(driver, url)
	if err != nil {
		return nil, err
	}

	return newSource(sql, driver), nil
}

func (db *DB) newBuilder(exec Executor) Builder {
	builderFunc, ok := BuilderMapFunc[db.driver]

	if !ok {
		builderFunc = nil
	}

	return builderFunc(db, exec)
}
