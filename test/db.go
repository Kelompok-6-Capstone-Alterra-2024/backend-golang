package test

import (
	"capstone/configs"
	cfg "capstone/repositories/mysql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

func setupTestDB() *gorm.DB {
	configs.LoadEnv()
	config := configs.InitConfigMySQL()
	dbportint, _ := strconv.Atoi(config.DBPort)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPass,
		config.DBHost,
		dbportint,
		config.DBName,
	)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	cfg.InitMigrate(DB)
	return DB
}
