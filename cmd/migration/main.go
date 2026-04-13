package main

import (
	"fmt"

	"github.com/jarusuraj/schoolsystem/config"
)

func main() {
	config.LoadEnv()
	migration()
}
func migration() {
	var choice string
	fmt.Println("Select")
	fmt.Println("1 -> Migrate one step up")
	fmt.Println("2 -> Migrate one step down")
	fmt.Println("3 -> Migrate UP All")
	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)
	switch choice {
	case "1":
		config.OneUPQueryMigration()
	case "2":
		config.OneDownQueryMigration()
	case "3":
		config.AllUPQueryMigration()
	default:
		fmt.Println("No migration executed")
	}
}
