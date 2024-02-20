package main

import (
	"fmt"
	application "fund-aplly-back"
	"fund-aplly-back/config"
	"fund-aplly-back/docs"
)

//	@title			Fund Apply Backend
//	@version		1.0
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

// @BasePath	/
func main() {
	cfg := config.NewConfig()

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)

	application.Start(cfg)
}
