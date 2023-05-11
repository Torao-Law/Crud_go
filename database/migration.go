package database

import (
	"Product/models"
	"Product/pkg/mysql"

	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.Product{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
