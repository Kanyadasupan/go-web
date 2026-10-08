package main

import (
	"gorm.io/gorm"
)

//ประกาศ struct ที่มีโครงสร้างเดียวกับตารางในdatabase
type Product struct {
	gorm.Model //จะสร้าง field id,created_at,updated_at,deleted_at ให้โดยอัตโนมัติ
	Name string 
	Price float32
	Quantity uint
}

