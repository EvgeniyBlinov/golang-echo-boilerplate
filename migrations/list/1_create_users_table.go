package list

import (
	postgres "github.com/EvgeniyBlinov/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateUserTable struct{}

func (m *CreateUserTable) GetName() string {
	return "CreateUserTable"
}

func (m *CreateUserTable) Up(con *sqlx.DB) {
	table := postgres.NewTable("users", con)
	table.Column("uuid").Type("uuid").Default("uuid_generate_v7()")
	table.PrimaryKey("uuid")
	table.String("name", 500).Nullable()
	table.String("password", 1000).Nullable()
	table.Column("deleted_at").Type("timestamp").Nullable()
	table.WithTimestamps()

	table.MustExec()
}

func (m *CreateUserTable) Down(con *sqlx.DB) {
	postgres.DropTable("users", con).MustExec()
}
