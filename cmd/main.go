package main

import "github.com/juanlavallen/godgh/internal/application"

func main() {
	app := application.NewApplication()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
