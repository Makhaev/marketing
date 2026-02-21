package main

import (
	"fmt"

	"github.com/Makhaev/marketing/internal/db"
)

func main() {
	database := db.InitPostgres()
	fmt.Println("DB connected:", database != nil)
}
