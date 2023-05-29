package main

import (
	"io/ioutil"
	"log"
	"github.com/gin-gonic/gin"
)

func main(){
	art, err := ioutil.ReadFile("art.txt")
	if err != nil{
		log.Fatal(err)
	}
	r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.Data(200, 
        "application/json; charset=utf-8",
        art)
    })
    r.Run()
}
