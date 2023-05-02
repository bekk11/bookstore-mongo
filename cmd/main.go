package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	routes "mongo_go/controllers"
	"mongo_go/pkg/common/db"
	"path/filepath"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	router := gin.Default()

	router.GET("images/:imagename", func(ctx *gin.Context) {
		fileName := ctx.Param("imagename")
		filePath := filepath.Join("./images/", fileName)
		ctx.File(filePath)
	})

	dbHandler := db.Init(dbUrl)

	routes.CategoryRoutes(router, dbHandler)
	routes.BookRoutes(router, dbHandler)
	routes.OrderRoutes(router, dbHandler)

	router.Run("localhost" + port)
}
