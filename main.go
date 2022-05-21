package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yintamaa/998244353Plus1/entity/config"
)

func main() {
	config.GetConfigInstance()
	r := gin.Default()
	initRouter(r)

	r.Run("10.0.20.8:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
