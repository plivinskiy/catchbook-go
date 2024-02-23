package main

import (
	"catchbook/internal/migration"
	"fmt"
)

func main() {
	fmt.Println("Database migration")

	migrations := migration.Migrations()
	for _, m := range migrations {
		m.Migrate()
	}
}
