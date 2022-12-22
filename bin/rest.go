package main

import (
	"github.com/orlandoxu/ginp"
	"log"
)

func main() {
	app := ginp.New()
	app.Post("/hello", func(c *ginp.Context) {
		log.Println(c.Query("a"))
		c.Next()
	}, func(c *ginp.Context) {
	}, func(c *ginp.Context) {
		log.Println("333 start!")
		log.Println("333 end!")
		c.Json(0, "", 121)
	})

	app.Run(":8084")
}
