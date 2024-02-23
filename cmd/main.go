package main

import (
	application "echo-demo-project"
	"echo-demo-project/config"
	"echo-demo-project/docs"
	"fmt"
)

//	@title			Echo Demo App
//	@version		1.0
//	@description	This is a demo version of Echo app.

//	@contact.name	Blinov Evgeniy
//	@contact.url	https://evgeniyblinov.ru/
//	@contact.email	evgeniy_blinov@mail.ru

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

// @BasePath	/
func main() {
	cfg := config.NewConfig()

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)

	application.Start(cfg)
}
