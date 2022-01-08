package main

import (
	"LineManWongNaiInternShip/miniproject/config"
	"github.com/gin-gonic/gin"
	ginLogRus "github.com/toorop/gin-logrus"
	"log"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		} else {
			c.Next()
		}
	}
}
func init() { log.SetFlags(log.Lshortfile | log.LstdFlags) }

func main() {
	appConfig := config.Get()
	logs := setupLog()
	router := gin.New()
	router.Use(ginLogRus.Logger(logs), gin.Recovery())
	router.Use(CORSMiddleware())
	_ = newApp(appConfig).RegisterRoute(router)
	_ = router.Run()
}
