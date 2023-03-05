package repository

import "log"

func migration() {
	err := DB.Set("gorm: table_options", "charset=utf8mb4").
		AutoMigrate(
			&Task{},
		)
	if err != nil {
		log.Fatalln("migration failed, err:", err)
	}
}
