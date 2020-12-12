package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

//replace eventconf code with viper

func SetupRouter() *gin.Engine {
	c, err := LoadConf()
	if err != nil {
		log.Fatal("unable to load conf")
	}
	oconf := NewOauthConfig(c)
	r := gin.Default()
	p := NewPath()
	o := NewOauth(oconf)

	r.GET("/", p.Proxy)
	r.GET("login", o.Login)
	r.GET("callback", o.Callback)
	return r
}

func main() {
	r := SetupRouter()
	r.Run(":8989")

}
