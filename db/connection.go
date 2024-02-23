package db

import (
	"echo-demo-project/config"
	"echo-demo-project/db/seeders"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(cfg *config.Config) *gorm.DB {
	dataSourceName := fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		cfg.DB.Driver,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name)

	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	userSeeder := seeders.NewUserSeeder(db)
	userSeeder.SetUsers()

	return db
}
