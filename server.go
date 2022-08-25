package main

import (
	"ECHO-GORM/db"
	"ECHO-GORM/route"
)

func main() {
	db.Init()
	e := route.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
