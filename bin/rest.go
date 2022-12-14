package main

import "rf"

func main() {
	app := rf.New()
	app.Post("/hello", nil)

	app.Run(":8084")
}
