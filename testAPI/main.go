package main

import (
	"io/ioutil"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/pprof"
)

func main(){

	r := gin.Default()
	pprof.Register(r) // 性能
  r.GET("/", func(c *gin.Context) {
  	art, err := ioutil.ReadFile("art.txt")
  	if err != nil{
	  	log.Fatal(err)
	  }
    c.Data(200, 
        "application/json; charset=utf-8",
        art)
  })
  r.Run(":8081")
}
