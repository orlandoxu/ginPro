package main

import "rf"

func main() {
	app := rf.New()
	app.Post("/hello", nil)

	app.Run("127.0.0.1:8084")
}
