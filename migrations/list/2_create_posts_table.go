package list

import (
	postgres "github.com/EvgeniyBlinov/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreatePostTable struct{}

func (m *CreatePostTable) GetName() string {
	return "CreatePostTable"
}

func (m *CreatePostTable) Up(con *sqlx.DB) {
	table := postgres.NewTable("posts", con)
	table.Column("uuid").Type("uuid").NotNull().Default("uuid_generate_v7()")
	table.PrimaryKey("uuid")
	table.String("title", 500).Nullable()
	table.String("content", 1000).Nullable()
	table.Column("deleted_at").Type("timestamp").Nullable()
	table.Column("user_uuid").Type("uuid").NotNull()
	table.ForeignKey("user_uuid").
		Reference("users").
		On("uuid").
		OnDelete("cascade").
		OnUpdate("cascade")
	table.WithTimestamps()

	table.MustExec()
}

func (m *CreatePostTable) Down(con *sqlx.DB) {
	postgres.DropTable("posts", con).MustExec()
}
