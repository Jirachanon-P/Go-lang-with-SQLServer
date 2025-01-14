package main

import (
	"first-api/Config"
	"first-api/Models"
	"first-api/Routes"
	"fmt"
	"log"

	_ "first-api/docs"

	_ "github.com/denisenkom/go-mssqldb"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var err error

// @title Teat Golang Swagger API
// @version 1.0
// @description Test Golog Swagger
// @termsOfService http://somewhere.com/

// @contact.name API Support
// @contact.url http://somewhere.com/support
// @contact.email support@somewhere.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @schemes http https

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @BasePath /api/v1
func main() {
	dsn := Config.DbURL(Config.BuildDBConfig())
	fmt.Print(Config.DbURL(Config.BuildDBConfig()))
	Config.DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Status:", err)
	}
	err = Config.DB.AutoMigrate(&Models.Note{})
	if err != nil {
		log.Fatalf("Error during migration: %v", err)
	}
	r := Routes.SetupRouter()
	//running
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
