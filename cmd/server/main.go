package main

import "github.com/olukkas/ushort/internal"

func main() {
	app := internal.NewApp()
	app.Init()

	err := app.Close()
	if err != nil {
		panic(err.Error())
	}
}
