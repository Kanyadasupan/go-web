package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"
	port     = "3306"
	user     = "root"
	password = ""
	dbname   = "stock_db"
)

func main() {
	//สร้างตัวแปรเก็บ database connection
	//dsn := "root:@tcp(localhost:3306)/stock_db?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)

	//สร้าง Log เพื่อเอาไว้ trace
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			Colorful:                  false,         // Disable color
		},
	)

	//สร้าง เชื่อมต่อฐานข้อมูล
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	//ตรวจสอบว่า เชื่อมต่อ database ได้มั้ย
	if err != nil {
		panic("Failed to connect to database")
	}

	//แสดงข้อความว่าเชื่อมต่อฐานข้อมูลสำเร็จ
	fmt.Println(db)
	fmt.Println("Database connection success")

	// ทำ Migration ตัว Moddel กับฐานข้อมูล
	db.AutoMigrate(&Product{})
	fmt.Println("Migration success")
}
