package deta

type Builder interface {
	RawQuery(query string)
	Select(table string)

	Insert(table string)
	Upsert(table string)
	Delete(table string)
	Update(table string)

	CreateTable()
	RenameTable()
	DropTable()
	TruncateTable()

	AddColumn()
	RenameColumn()
	DropColumn()
	AlterColumn()
}

type BaseBuilder struct {
	db       *DB
	executor Executor
}

func NewBaseBuilder(db *DB, executor Executor) *BaseBuilder {
	return &BaseBuilder{db, executor}
}
