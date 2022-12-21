package main

import (
	"log"
	"rf"
)

func main() {
	app := rf.New()
	app.Post("/hello", func(c *rf.Context) {
		log.Println("111 start!")
		c.Next()
		log.Println("111 end!")
	}, func(c *rf.Context) {
		log.Println("222 start!")
		c.Next()
		log.Println("222 end!")
	}, func(c *rf.Context) {
		log.Println("333 start!")
		log.Println("333 end!")
	})

	app.Run(":8084")
}
