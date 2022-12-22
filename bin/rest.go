package main

import (
	"ginPro"
	"log"
)

func main() {
	app := ginPro.New()
	app.Post("/hello", func(c *ginPro.Context) {
		c.Next()
	}, func(c *ginPro.Context) {
	}, func(c *ginPro.Context) {
		log.Println("333 start!")
		log.Println("333 end!")
		c.Json(0, "", 121)
	})

	app.Run(":8084")
}
