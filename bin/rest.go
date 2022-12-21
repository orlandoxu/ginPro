package main

import (
	"log"
	"rf"
)

func main() {
	app := rf.New()
	app.Post("/hello", func(c *rf.Context) {
		c.Next()
	}, func(c *rf.Context) {
	}, func(c *rf.Context) {
		log.Println("333 start!")
		log.Println("333 end!")
		c.Json(0, "", 121)
	})

	app.Run(":8084")
}
