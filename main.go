package main

import (
	database "go-server/db"
	"go-server/router"
)

func main() {
	database.GetDb()
	router.ListenAndServe(router.GetHandlers())

}
