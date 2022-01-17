package main

import (
	"fmt"
	"hc21f/pkg/database"
)

func main() {
	database.Init()
	db := database.Get()

	rs, _ := db.Query(
		"SELECT * FROM friends WHERE uid = 'aa' LIMIT 1",
	)

	for rs.Next() {
		fmt.Println("aaa")
		break
	}
}
