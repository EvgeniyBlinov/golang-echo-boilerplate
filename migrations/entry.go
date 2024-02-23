package main

import (
	"echo-demo-project/migrations/list"

	gmEngine "github.com/EvgeniyBlinov/go-migrations/engine"
	gmStore "github.com/EvgeniyBlinov/go-migrations/engine/store"
)

func main() {
	e := gmEngine.NewEngine()
	e.Run(getMigrationsList())
	e.GetConnector().Close()
}

func getMigrationsList() []gmStore.Migratable {
	return []gmStore.Migratable{
		&list.CreateUserTable{},
		&list.CreatePostTable{},
		&list.UpdateUserTable{},
	}
}
